Every function below works on any `Arrow[A, B]`, whatever it wraps. The **Example** links point to runnable examples in `doc_examples_test.go`.

## Constructing arrows

| Function | Signature | Notes |
|---|---|---|
| `Id` | `Id[A]() Arrow[A, A]` | Identity: returns its input unchanged |
| `Pure` | `Pure(f func(A) B) Arrow[A, B]` | Lifts a function that can't fail |
| `ArrowFunc` | `ArrowFunc(f func(ctx, A) (B, error)) Arrow[A, B]` | Names a plain function as an Arrow |

## Sequential composition

| Function | Signature | Notes |
|---|---|---|
| `Compose` | `Compose(f Arrow[A, B], g Arrow[B, C]) Arrow[A, C]` | Runs `f`, then `g`. Stops at the first error. Example: `ExampleCompose` |
| `Pipe2` … `Pipe6` | `Pipe3(f, g, h)` and so on | The same as nested `Compose`, flatter to read. Example: `ExamplePipe3` |
| `Map` | `Map(f Arrow[A, B], h func(B) C) Arrow[A, C]` | Post-processes the output with a pure function |
| `PreMap` | `PreMap(h func(A) B, f Arrow[B, C]) Arrow[A, C]` | Pre-processes the input with a pure function |

## Products: running in parallel

| Function | Signature | Notes |
|---|---|---|
| `Par` | `Par(f Arrow[A, B], g Arrow[A, C]) Arrow[A, Pair[B, C]]` | Runs both concurrently on the same input. Both errors are joined with `errors.Join`. Example: `ExamplePar` |
| `ParStrict` | same as `Par` | Cancels the sibling as soon as one fails, and returns that first error |
| `Par3` | three arrows → `Triple` | Three-way `Par` |
| `Fanout` | same as `Par` | Identical to `Par`; the name signals intent |
| `Fst`, `Snd` | `Arrow[Pair[A, B], A]`, `Arrow[Pair[A, B], B]` | Projections |
| `MakePair` | `MakePair(a, b) Pair[A, B]` | Constructor |

## Coproducts: routing

| Function | Signature | Notes |
|---|---|---|
| `Sum` | `Sum(f Arrow[A, C], g Arrow[B, C]) Arrow[Either[A, B], C]` | Runs whichever branch the input selects. Example: `ExampleSum` |
| `Sum4` | four branches → `R` | Four-way routing over `FourWay` |
| `Left`, `Right` | `Left[A, B](a) Either[A, B]` | Constructors. `IsLeft()` and `IsRight()` inspect the value |
| `Fallback` | `Fallback(primary, backup Arrow[A, B]) Arrow[A, B]` | Runs `backup` if `primary` fails. Cancellation is never recovered. Example: `ExampleFallback` |
| `OnSentinel` | `OnSentinel(sentinel, backup) Transform` | Falls back only on one specific error |

## Collections

| Function | Signature | Notes |
|---|---|---|
| `Traverse` | `Traverse(f Arrow[A, B], opts...) Arrow[[]A, []B]` | Applies `f` to each element, keeping input order. Example: `ExampleTraverse` |
| `WithConcurrency` | `WithConcurrency(n int)` | Limits parallelism. Default `1`; values below `1` become `1` |
| `OnError` | `OnError(p ErrorPolicy)` | What happens if any element fails (see below) |

What `Traverse` returns when some elements fail:

| Policy | Result | Error |
|---|---|---|
| `FailFast` (default) | `nil` | the first failure, by input position |
| `CollectErrors` | all slots; failed ones hold zero values | every failure, joined with `errors.Join` |
| `SkipFailures` | only the successful results, in order | `nil` |
| `PartialResults` | all slots; failed ones hold zero values | `*PartialError` with `Failures` keyed by index, and `Total` |

## Transforms: cross-cutting behaviour

A `Transform[A, B]` is a `func(Arrow[A, B]) Arrow[A, B]`, so it changes behaviour without changing the arrow's type.

| Function | Notes |
|---|---|
| `Apply(a, t1, t2, ...)` | Applies transforms. The **first is outermost**, so `Apply(a, WithRetry(...), WithTimeout(...))` sets a timeout per attempt. Example: `ExampleApply` |
| `WithRetry(maxAttempts, backoff)` | Retries transient errors; doesn't retry `ClassPermanent`. Respects cancellation during backoff |
| `LinearBackoff(d)`, `ExponentialBackoff(base)` | `BackoffFunc` values for `WithRetry` |
| `WithTimeout(d)` | Cancels the context after `d`. The error wraps `context.DeadlineExceeded` |
| `WithTap(func(in, out, err))` | Observes each call without changing the result, for logging or metrics |

## Errors

| Name | Notes |
|---|---|
| `ArrowError` | Wraps an error with a `Class`. Supports `errors.Is` and `errors.As` through `Unwrap` |
| `Class` | `ClassUnknown`, `ClassTransient`, `ClassPermanent`, `ClassBudget`, `ClassUserCancelled` |
| `Classify(err)` | Gets the class of any error. Plain errors count as `ClassUnknown`, which `WithRetry` treats as retryable |
| `PartialError` | Returned by `Traverse` under `PartialResults` |
