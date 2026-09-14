$text = @"
	success := tryModel(r.cfg.Ollama.GenerationModel, false)
	var handled bool
	var result executor.RunResult
	var rejectReason string

	if success {
		handled, result, rejectReason = executeAndCheck()
		if handled {
			r.recordGenerated(generated)
			r.saveHistory(input, "nl", generated, result.ExitCode, result.Output, result.DurationMs)
			return
		}
	}

	if (!success || (!handled && rejectReason != "")) && r.cfg.Ollama.FallbackModel != "" && r.cfg.Ollama.FallbackModel != r.cfg.Ollama.GenerationModel {
		success = tryModel(r.cfg.Ollama.FallbackModel, true)
		if success {
			handled, result, rejectReason = executeAndCheck()
			if handled {
				r.recordGenerated(generated)
				r.saveHistory(input, "nl", generated, result.ExitCode, result.Output, result.DurationMs)
				return
			}
		}
	}

	if !success {
"@
Set-Content -Path handleNL_bottom_replacement.txt -Value $text
