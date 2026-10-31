## Ground rules

- **Standard library only.** weft has no dependencies, and a change that adds one won't be merged.
- **Laws first.** A new combinator needs a law test in `laws_test.go` if a law applies, or a behaviour test if not.
- **Every exported function gets a runnable example** in `doc_examples_test.go` with an `// Output:` comment. The examples are the documentation, and `go test` keeps them honest.
- **Update this wiki in the same change.** The pages live in `docs/wiki/` in this repo, and every push to `main` publishes them here. Don't edit pages on GitHub directly; the next sync overwrites them.

- Supported platforms are Linux (amd64, arm64) and macOS on Apple Silicon.

## Before opening a pull request

```bash
gofmt -l .        # prints nothing
go vet ./...
go test -race ./...
```
