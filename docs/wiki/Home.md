**weft** is a small Go library for composing effectful computations as typed arrows, with the category laws checked by tests. It has no dependencies outside the standard library.

```go
type Arrow[A, B any] func(ctx context.Context, a A) (B, error)
```

[[/images/compose-par-sum.svg|alt=String diagrams of Compose (f then g), Par (f and g side by side) and Sum (route to f or g)]]

A pure function, an HTTP call, an LLM request and a whole multi-step pipeline all have this one type. So the same combinators compose all of them: `Compose`, `Par`, `Sum`, `Traverse`, `WithRetry` and the rest.

## Where weft fits in heedle-ai

| Repo | Role |
|---|---|
| **weft** | Composition algebra (this repo) |
| [draft](https://github.com/heedle-ai/draft/wiki) | Language: s-expressions that compile to shuttle plans |
| [warp](https://github.com/heedle-ai/warp/wiki) | Content-addressed object store |
| [shuttle](https://github.com/heedle-ai/shuttle/wiki) | Durable execution on Temporal |
| [bobbin](https://github.com/heedle-ai/bobbin/wiki) | A fresh Podman container per agent step |
| [drawdown](https://github.com/heedle-ai/drawdown/wiki) | Web UI: write, run and browse programs and records |
| [loom](https://github.com/heedle-ai/loom/wiki) | Runs them together: compile, execute, record |

## Pages

- [Getting Started](Getting-Started): install Go, add weft, run the tests
- [Combinators](Combinators): every exported function, grouped by what it does
- [Laws](Laws): which category laws hold, and how each one is tested
- [Invariants](Invariants): every guarantee, the test that guards it, and the mistake it came from
- [Contributing](Contributing): conventions for changes

## Requirements

- Go **1.23** or later
- Linux (amd64 or arm64), or macOS on Apple Silicon
