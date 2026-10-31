# weft

Categorical composition primitives for Go. Part of [heedle-ai](https://github.com/heedle-ai).

Every step of a computation (a pure function, an LLM call, a tool call, a
parallel fan-out) is the same type:

```go
type Arrow[A, B any] func(ctx context.Context, a A) (B, error)
```

`Arrow` is a Kleisli arrow for the error-and-context effect. The combinators
(`Compose`, `Pipe2`–`Pipe6`, `Par`, `Sum`, `Map`, `Traverse`, `WithRetry`,
`WithTimeout`, `WithTap`, `Fallback`, …) work uniformly on any arrow.

The category laws (identity, associativity, functor, product, coproduct) are
the package's specification and are checked by property-based tests in
`laws_test.go`.

```bash
go test ./...
```

This module depends only on the Go standard library.

## Provenance

Copied from [`vinodhalaharvi/weft`](https://github.com/vinodhalaharvi/weft) at
`edca607`, core package (`weft/`) only, flattened to the module root. The
`llm/` and `mcp/` packages and the `cmd/` examples were left behind on purpose.
