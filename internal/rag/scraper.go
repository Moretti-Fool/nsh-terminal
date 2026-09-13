package rag

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	scriptRegex      = regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`)
	styleRegex       = regexp.MustCompile(`(?is)<style[^>]*>.*?</style>`)
	tagRegex         = regexp.MustCompile(`(?is)<[^>]+>`)
	spaceRegex       = regexp.MustCompile(`[ \t]+`)
	newlineRegex     = regexp.MustCompile(`\n{3,}`)
)

func ScrapeAndChunk(url string) (Document, []string, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Document{}, nil, err
	}
	req.Header.Set("User-Agent", "nsh-rag-bot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return Document{}, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Document{}, nil, fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Document{}, nil, err
	}
	html := string(body)

	// Very basic title extraction
	title := url
	titleMatch := regexp.MustCompile(`(?is)<title>(.*?)</title>`).FindStringSubmatch(html)
	if len(titleMatch) > 1 {
		title = strings.TrimSpace(titleMatch[1])
	}

	// Strip HTML
	text := scriptRegex.ReplaceAllString(html, " ")
	text = styleRegex.ReplaceAllString(text, " ")
	text = tagRegex.ReplaceAllString(text, " ")
	text = spaceRegex.ReplaceAllString(text, " ")
	
	// Convert entities roughly or just ignore for v1
	text = strings.ReplaceAll(text, "&nbsp;", " ")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&amp;", "&")

	// Normalize newlines
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = newlineRegex.ReplaceAllString(text, "\n\n")

	// Chunking
	paragraphs := strings.Split(text, "\n\n")
	var chunks []string
	var currentChunk strings.Builder

	maxLen := 1500 // rough char limit per chunk (~300 tokens)
	for _, p := range paragraphs {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		
		if currentChunk.Len()+len(p) > maxLen && currentChunk.Len() > 0 {
			chunks = append(chunks, currentChunk.String())
			currentChunk.Reset()
		}
		
		if currentChunk.Len() > 0 {
			currentChunk.WriteString("\n\n")
		}
		currentChunk.WriteString(p)
	}
	if currentChunk.Len() > 0 {
		chunks = append(chunks, currentChunk.String())
	}

	docID := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:16]

	doc := Document{
		ID:    docID,
		Title: title,
		URL:   url,
	}

	return doc, chunks, nil
}
