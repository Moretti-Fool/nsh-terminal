package repl

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/nsh-terminal/nsh/internal/rag"
)

func (r *REPL) handleDoc(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: nsh doc <add|list|ask> [args...]")
		return
	}

	cmd := args[0]
	switch cmd {
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: nsh doc add <url>")
			return
		}
		r.handleAddDoc(args[1])
	case "list":
		r.handleListDocs()
	case "ask":
		if len(args) < 2 {
			fmt.Println("Usage: nsh doc ask <question>")
			return
		}
		r.handleAskDoc(strings.Join(args[1:], " "))
	default:
		fmt.Printf("Unknown doc command: %s\n", cmd)
	}
}

func (r *REPL) handleAddDoc(url string) {
	if r.ragStore == nil {
		fmt.Println("RAG store is not initialized.")
		return
	}

	fmt.Printf("Fetching %s...\n", url)
	doc, chunks, err := rag.ScrapeAndChunk(url)
	if err != nil {
		fmt.Printf("Failed to scrape: %v\n", err)
		return
	}

	fmt.Printf("Embedding %d chunks (this may take a moment)...\n", len(chunks))
	ctx := context.Background()
	// Using the default embedding model nomic-embed-text
	embeddings, err := r.ollama.GenerateEmbeddings(ctx, "nomic-embed-text", chunks)
	if err != nil {
		fmt.Printf("Failed to embed: %v\n", err)
		return
	}

	var ragChunks []rag.Chunk
	for i, emb := range embeddings {
		ragChunks = append(ragChunks, rag.Chunk{
			ID:        fmt.Sprintf("%s-%d", doc.ID, i),
			DocID:     doc.ID,
			Text:      chunks[i],
			Embedding: emb,
		})
	}

	err = r.ragStore.AddDocument(doc, ragChunks)
	if err != nil {
		fmt.Printf("Failed to save doc: %v\n", err)
		return
	}

	fmt.Printf("Successfully added document: %s\n", doc.Title)
}

func (r *REPL) handleListDocs() {
	if r.ragStore == nil {
		fmt.Println("RAG store is not initialized.")
		return
	}

	docs := r.ragStore.ListDocuments()
	if len(docs) == 0 {
		fmt.Println("No documents available. Add one using 'nsh doc add <url>'.")
		return
	}

	fmt.Println("Available Documentation:")
	for _, doc := range docs {
		fmt.Printf("  - %s (%s)\n", doc.Title, doc.URL)
	}
}

func (r *REPL) handleAskDoc(query string) {
	if r.ragStore == nil {
		fmt.Println("RAG store is not initialized.")
		return
	}

	ctx := context.Background()
	emb, err := r.ollama.GenerateEmbeddings(ctx, "nomic-embed-text", []string{query})
	if err != nil || len(emb) == 0 {
		fmt.Printf("Failed to embed query: %v\n", err)
		return
	}

	results := r.ragStore.Search(emb[0], 3, "")
	if len(results) == 0 {
		fmt.Println("No relevant information found in the documentation.")
		return
	}

	var contextBuilder strings.Builder
	for i, res := range results {
		contextBuilder.WriteString(fmt.Sprintf("--- Excerpt %d ---\n%s\n\n", i+1, res.Chunk.Text))
	}

	prompt := fmt.Sprintf("You are an expert AI assistant guiding a user based ONLY on the provided documentation context.\n\nContext:\n%s\nUser Question: %s", contextBuilder.String(), query)

	fmt.Printf("\033[1;36m[nsh doc] Asking local model...\033[0m\n")
	answer, err := r.ollama.Ask(ctx, prompt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "AI error: %v\n", err)
		return
	}

	fmt.Printf("\n%s\n\n", answer)
}
