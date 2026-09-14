$text = @"
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.cfg.Ollama.TimeoutMs)*time.Millisecond)
	defer cancel()

	judgeModel := r.cfg.Ollama.JudgeModel
	if judgeModel == "" {
		judgeModel = "qwen2.5:0.5b"
	}

	// 1. Check local cache layer for exact or very similar past commands
	cachedCmd := r.semanticCacheMatch(ctx, input, r.history.SuccessfulNL(200))
	if cachedCmd != "" {
		fmt.Printf("\n[nsh] (from memory) \033[36m> %s\033[0m\n", cachedCmd)
		if r.confirmIfDestructive(cachedCmd) {
			ok, result := r.runGeneratedCommands(cachedCmd)
			if ok {
				r.recordGenerated(cachedCmd)
				r.saveHistory(input, "nl", cachedCmd, result.ExitCode, result.Output, result.DurationMs)
			}
		}
		return
	}
"@
Set-Content -Path handleNL_top_replacement.txt -Value $text
