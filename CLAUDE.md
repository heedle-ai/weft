# weft

Categorical composition primitives for Go. Part of heddle:
weft (algebra) → warp (object store) → shuttle (execution) → drawdown (UI).

## Rules for every change

1. **Keep the wiki current.** Any change that affects how a developer
   installs, builds, tests, uses or understands this repo updates the wiki
   in the same piece of work: https://github.com/heedle-ai/weft/wiki
   (git: `https://github.com/heedle-ai/weft.wiki.git`). Code, README and wiki
   must never disagree. If the wiki can't be updated, say so explicitly.
2. **Wiki quality bar.** Every command is copy-pasteable and has been run.
   Every code snippet compiles and is checked by a test; prefer quoting a
   Go `Example` test with an `// Output:` comment. State required versions.
   No placeholders, TODOs or "coming soon" on published pages. Mark design
   that isn't built yet as **Planned**, plainly.
3. **Platforms: Linux (amd64, arm64) and macOS on Apple Silicon only.** Don't
   write Windows instructions, scripts, CI jobs or path handling.
4. **Bare minimum.** Add only what the task needs. No speculative packages,
   dependencies or abstractions.
5. **Tests pass before commit:** `go vet ./...` and `go test -race ./...`.

## Repo-specific

- **Standard library only.** weft has no dependencies; keep it that way.
- **The laws are the spec.** `laws_test.go` holds the category laws. A change
  that breaks one is wrong, not the test. New combinators get law or
  behaviour tests, and a runnable `Example` in `doc_examples_test.go`.
