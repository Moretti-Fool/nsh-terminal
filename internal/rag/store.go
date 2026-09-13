package rag

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type Chunk struct {
	ID        string            `json:"id"`
	DocID     string            `json:"doc_id"`
	Text      string            `json:"text"`
	Embedding []float32         `json:"embedding"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type Document struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type DB struct {
	Documents map[string]Document `json:"documents"`
	Chunks    []Chunk             `json:"chunks"`
}

type Store struct {
	path string
	db   DB
	mu   sync.RWMutex
}

func NewStore(dataDir string) (*Store, error) {
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		return nil, err
	}
	path := filepath.Join(dataDir, "rag_db.json")
	
	s := &Store{
		path: path,
		db: DB{
			Documents: make(map[string]Document),
			Chunks:    []Chunk{},
		},
	}
	s.load()
	return s, nil
}

func (s *Store) load() {
	b, err := os.ReadFile(s.path)
	if err == nil {
		if err := json.Unmarshal(b, &s.db); err != nil {
			fmt.Printf("[nsh] Warning: failed to parse RAG database (%s): %v\n", s.path, err)
		}
	}
	if s.db.Documents == nil {
		s.db.Documents = make(map[string]Document)
	}
}

func (s *Store) save() error {
	b, err := json.MarshalIndent(s.db, "", "  ")
	if err != nil {
		return err
	}
	tmpPath := s.path + ".tmp"
	if err := os.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}
	return os.Rename(tmpPath, s.path)
}

func (s *Store) AddDocument(doc Document, chunks []Chunk) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.db.Documents[doc.ID] = doc
	
	// Remove old chunks for this doc if it already exists
	var newChunks []Chunk
	for _, c := range s.db.Chunks {
		if c.DocID != doc.ID {
			newChunks = append(newChunks, c)
		}
	}
	newChunks = append(newChunks, chunks...)
	s.db.Chunks = newChunks

	return s.save()
}

func (s *Store) AddChunk(chunk Chunk) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.db.Chunks = append(s.db.Chunks, chunk)
	return s.save()
}

func (s *Store) ListDocuments() []Document {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var docs []Document
	for _, d := range s.db.Documents {
		docs = append(docs, d)
	}
	return docs
}

type SearchResult struct {
	Chunk Chunk
	Score float32
}

func cosineSimilarity(a, b []float32) float32 {
	if len(a) != len(b) {
		return 0
	}
	var dot, normA, normB float32
	for i := range a {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dot / (float32(math.Sqrt(float64(normA))) * float32(math.Sqrt(float64(normB))))
}

func (s *Store) Search(queryEmbedding []float32, topK int, filters map[string]string) []SearchResult {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var results []SearchResult
	for _, chunk := range s.db.Chunks {
		match := true
		if filters != nil && chunk.Metadata != nil {
			for k, v := range filters {
				if chunk.Metadata[k] != v {
					match = false
					break
				}
			}
		}
		if !match {
			continue
		}
		
		score := cosineSimilarity(queryEmbedding, chunk.Embedding)
		results = append(results, SearchResult{Chunk: chunk, Score: score})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if len(results) > topK {
		results = results[:topK]
	}
	return results
}
