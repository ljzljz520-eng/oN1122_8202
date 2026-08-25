# BUG_REPRO

The following failures were observed while validating the initial project state.
Each section records what failed, how to reproduce it, and the complete command output.
They are preserved intentionally; only failing build gates are omitted from the generated Dockerfile.

## Failure 1: Go test (.)

- Observed problem: `Go test (.)` failed in the initial project state.
- Working directory: `.`
- Command: `cd /app && GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -count=1 ./...`
- Exit status: `1`

```text
ok  	example.com/cookproposal/cmd/cookproposal	0.007s
ok  	example.com/cookproposal/internal/api	0.021s
ok  	example.com/cookproposal/internal/audit	0.006s
ok  	example.com/cookproposal/internal/domain	0.009s
--- FAIL: Test1122BusinessRegression (0.01s)
    flow_test.go:75: record r2 permission changed to team
FAIL
FAIL	example.com/cookproposal/internal/flow027	0.036s
ok  	example.com/cookproposal/internal/importer	0.021s
ok  	example.com/cookproposal/internal/query	0.006s
ok  	example.com/cookproposal/internal/store	0.027s
ok  	example.com/cookproposal/internal/workflow	0.005s
FAIL
```

## Architecture reproduction

### linux/amd64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/cookproposal): exit `0`
### linux/arm64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/cookproposal): exit `0`
