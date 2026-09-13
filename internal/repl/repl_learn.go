package repl

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/nsh-terminal/nsh/internal/rag"
)

func (r *REPL) handleLearn(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: nsh learn \"<description>\" \"<command>\"")
		return
	}

	nl := args[0]
	cmd := args[1]

	r.learnSingle(nl, cmd)
}

func (r *REPL) learnSingle(nl, cmd string) {
	if r.ragStore == nil {
		fmt.Println("RAG store is not initialized.")
		return
	}

	ctx := context.Background()
	emb, err := r.ollama.GenerateEmbeddings(ctx, "nomic-embed-text", []string{nl})
	if err != nil || len(emb) == 0 {
		fmt.Printf("Failed to embed: %v\n", err)
		return
	}

	// We'll store it in a special "Learned Commands" document
	docID := "learned_commands"
	
	// Create doc if it doesn't exist
	docs := r.ragStore.ListDocuments()
	found := false
	for _, d := range docs {
		if d.ID == docID {
			found = true
			break
		}
	}
	if !found {
		// Just to ensure doc exists
		r.ragStore.AddDocument(rag.Document{
			ID:    docID,
			Title: "Learned Commands",
			URL:   "local://learned",
		}, nil)
	}

	chunk := rag.Chunk{
		ID:        fmt.Sprintf("learn-%d", len(r.ragStore.ListDocuments())*10000+len(emb[0])), // Rough ID
		DocID:     docID,
		Text:      nl,
		Embedding: emb[0],
		Metadata:  map[string]string{"command": cmd, "type": "few-shot"},
	}

	err = r.ragStore.AddChunk(chunk) // We need to add AddChunk to Store
	if err != nil {
		fmt.Printf("Failed to save learned command: %v\n", err)
		return
	}
	fmt.Printf("[nsh] Learned: %s -> %s\n", nl, cmd)
}

func (r *REPL) handleLearnImport(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: nsh learn-import <dataset.jsonl>")
		return
	}

	path := args[0]
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	var nlList []string
	var cmdList []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var entry map[string]interface{}
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			continue
		}
		
		messages, ok := entry["messages"].([]interface{})
		if !ok {
			continue
		}

		var nl, cmd string
		for _, mInt := range messages {
			m, ok := mInt.(map[string]interface{})
			if !ok {
				continue
			}
			role, _ := m["role"].(string)
			content, _ := m["content"].(string)
			if role == "" {
				continue
			}

			if role == "user" && nl == "" {
				nl = content // Only take the FIRST user message (the actual intent)
			} else if role == "assistant" {
				// Parse command JSON (take the LAST successful one, overwriting any failed attempts)
				var cmdJSON map[string]interface{}
				if json.Unmarshal([]byte(content), &cmdJSON) == nil {
					if cmds, ok := cmdJSON["commands"].([]interface{}); ok && len(cmds) > 0 {
						var cmdLines []string
						for _, cInt := range cmds {
							cmdLines = append(cmdLines, cInt.(string))
						}
						cmd = strings.Join(cmdLines, "; ")
					}
				}
			}
		}

		if nl != "" && cmd != "" {
			nlList = append(nlList, nl)
			cmdList = append(cmdList, cmd)
		}
	}

	if len(nlList) == 0 {
		fmt.Println("No valid training pairs found.")
		return
	}

	fmt.Printf("Found %d pairs. Embedding... This will take a while.\n", len(nlList))

	ctx := context.Background()
	
	// Batch processing embeddings
	batchSize := 20
	var ragChunks []rag.Chunk
	docID := "learned_dataset"

	for i := 0; i < len(nlList); i += batchSize {
		end := i + batchSize
		if end > len(nlList) {
			end = len(nlList)
		}

		fmt.Printf("\rEmbedding batch %d/%d...", (i/batchSize)+1, (len(nlList)+batchSize-1)/batchSize)
		emb, err := r.ollama.GenerateEmbeddings(ctx, "nomic-embed-text", nlList[i:end])
		if err != nil {
			fmt.Printf("\nFailed to embed batch: %v\n", err)
			return
		}

		for j, embedding := range emb {
			ragChunks = append(ragChunks, rag.Chunk{
				ID:        fmt.Sprintf("ds-%d", i+j),
				DocID:     docID,
				Text:      nlList[i+j],
				Embedding: embedding,
				Metadata:  map[string]string{"command": cmdList[i+j], "type": "few-shot"},
			})
		}
	}
	fmt.Println()

	err = r.ragStore.AddDocument(rag.Document{
		ID:    docID,
		Title: "Imported Dataset",
		URL:   path,
	}, ragChunks)

	if err != nil {
		fmt.Printf("Failed to save dataset: %v\n", err)
		return
	}

	fmt.Printf("Successfully learned %d commands from dataset.\n", len(nlList))
}
