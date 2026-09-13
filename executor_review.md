# Executor Review Report

I have reviewed `internal/executor/pshost.go` and `internal/executor/executor.go` and found several critical logic bugs related to timeouts, unclosed pipes, missing error handles, and infinite loops.

## 1. Unclosed Channel and Potential Deadlock (`pshost.go`)
**Location:** `startPSHost()` goroutine (lines 183-195)
**Bug:** The goroutine that reads from the `stdoutPipe` sends an error to `h.lines` upon EOF or an error (like `bufio.ErrTooLong`), but **it never closes the `h.lines` channel**. 
**Impact:** If the channel is not closed, the goroutine exits but leaves the channel open. If `run()` processes the error and returns, subsequent calls to `run()` will wait on `h.lines`. Because no goroutine is writing to it anymore and it wasn't closed, `run()` will block indefinitely until the 2-minute `psCommandTimeout` expires.
**Fix:** Add `defer close(h.lines)` to the top of the goroutine.

## 2. Missing Error Handling & Process Leaks (`pshost.go`)
**Location:** `run()` (lines 250-254)
**Bug:** When `run()` receives a closed channel (`!ok`) or an error (`got.err != nil`) from `h.lines`, it simply returns the error. It completely fails to call `h.killUnlocked()` or set `h.ready = false`.
**Impact:** The host is left in a "ready" state even though the stdout reader goroutine has died. If the PowerShell process is still alive (e.g., if the scanner failed with `ErrTooLong`), any subsequent stdout writes from PowerShell will block forever because the pipe's read end is unread. This freezes the PowerShell process indefinitely (unclosed pipe / missing reader).
**Fix:** Call `_ = h.killUnlocked()` before returning in both the `!ok` and `got.err != nil` branches.

## 3. Timeout Logic Bug: Double Execution (`executor.go`)
**Location:** `RunGenerated()` (lines 135-146)
**Bug:** If `h.run(command, cwd)` returns an error (which includes the "powershell command timed out" error), the executor drops the host, forces a restart via `ensurePSHost()`, and then **retries the exact same command**.
**Impact:** If a user command legitimately hangs or contains an infinite loop, it will hit the 2-minute timeout in `h.run()`. `RunGenerated` will then restart the host and run the identical hanging command a second time, waiting another 2 minutes. This turns a 2-minute timeout into a 4-minute delay and repeats unintended side-effects.
**Fix:** Inspect the error returned by `h.run()`. If it is a timeout error, return the error immediately rather than retrying.

## 4. Missing Timeouts causing Infinite Hangs (`executor.go`)
**Location:** `runPowerShellOnce()`, `runDirect()`, `runViaShell()`, `RunSilent()`
**Bug:** These execution methods use `cmd.Run()` or `cmd.CombinedOutput()` synchronously without any `context.Context` timeouts. (Note that `RunInvestigate` correctly applies a 10-second timeout).
**Impact:** If a command executed via these methods waits for user input, triggers a prompt, or enters an infinite loop, the executor will block forever. Because there is no timeout mechanism on these paths, the caller will completely freeze without recovery.
**Fix:** Use `context.WithTimeout` on `exec.CommandContext` to ensure that runaway processes are forcefully killed and control is returned.
