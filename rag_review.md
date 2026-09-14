# RAG Package Review Report

This report outlines the issues found during a line-by-line review of the `internal/rag` package.

## 1. Logic Bugs
*   **Disk I/O inside Mutex (`store.go`):** In `AddDocument` and `AddChunk`, the `s.save()` function is called while holding the exclusive write lock (`s.mu.Lock()`). `s.save()` marshals the entire database and writes it to disk synchronously. This means all other operations (including reads like `Search`) are completely blocked while writing to disk. This is a severe performance bottleneck.
*   **OOM Risk on Save (`store.go`):** `s.save()` uses `json.Marshal(s.db)` on the entire in-memory database. For a large vector store, this will allocate a massive byte slice in memory, potentially leading to Out-Of-Memory (OOM) crashes.

## 2. Memory / Resource Leaks
*   **Dynamic Regex Compilation (`scraper.go`):** On line 47, `regexp.MustCompile("(?is)<title>(.*?)</title>")` is compiled inside the `ScrapeAndChunk` function. This means a new regex object is compiled on every single function call, which causes unnecessary CPU overhead and memory allocations. It should be extracted to a global variable like `scriptRegex`.

## 3. Missed Error Returns
*   **Ignored JSON Unmarshal Error (`store.go`):** In the `load()` function on line 58, `json.Unmarshal(b, &s.db)` is called but the returned error is completely ignored. If the `rag_db.json` file is corrupted or malformed, the application will silently fail to load the data instead of reporting the issue.

## 4. Hardcoded Variables
*   **HTTP Client Timeout (`scraper.go`):** `Timeout: 15 * time.Second` (Line 22) is hardcoded. It should ideally be configurable via parameters or a context.
*   **User-Agent (`scraper.go`):** `"nsh-rag-bot/1.0"` (Line 27) is hardcoded.
*   **Chunk Size Limit (`scraper.go`):** `maxLen := 1500` (Line 73) is hardcoded, limiting flexibility in chunking strategies.
*   **Database Filename (`store.go`):** `"rag_db.json"` (Line 42) is hardcoded. It should be customizable if a different filename or extension is required.
