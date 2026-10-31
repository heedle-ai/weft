weft treats `Arrow` as the morphisms of a Kleisli category for the "may fail, may be cancelled" effect. The laws below are the library's spec: a change that breaks one is a bug.

Two arrows are **equivalent** when, for the same input, they return equal outputs, or both fail with the same error message.

## Property-tested laws

These use `testing/quick`, checking 200 random inputs per law through `arrowEquiv` in `laws_test.go`.

| Law | Statement | Test |
|---|---|---|
| Left identity | `Compose(Id, f) ≡ f` | `TestIdentityLaw_Left` |
| Right identity | `Compose(f, Id) ≡ f` | `TestIdentityLaw_Right` |
| Identity on failure | identity laws hold when `f` fails | `TestIdentityLaw_OnFailingArrow` |
| Associativity | `Compose(Compose(f, g), h) ≡ Compose(f, Compose(g, h))` | `TestAssociativityLaw`, `TestAssociativityLaw_WithErrors` |
| Pipes are compositions | `Pipe3(f, g, h) ≡ Compose(Compose(f, g), h)`, and the same for `Pipe4` | `TestPipe3_…`, `TestPipe4_…` |
| Functor: identity | `Map(f, id) ≡ f` | `TestMap_PreservesIdentity` |
| Functor: composition | `Map(Map(f, h₁), h₂) ≡ Map(f, h₂ ∘ h₁)` | `TestMap_PreservesComposition` |
| `Map` is post-composition | `Map(f, h) ≡ Compose(f, Pure(h))` | `TestMap_EquivalentToComposeWithPure` |

## Behaviour checked with examples

These are tested on specific cases, not random inputs:

- **Product (`Par`)**: both branches run, run concurrently, and errors are joined. `ParStrict` cancels the sibling on the first failure.
- **Coproduct (`Sum`)**: `Left` goes to the first branch and `Right` to the second. `Id` works as a branch.
- **`Fallback`**: the backup runs only when the primary fails, and never after cancellation.
- **Cancellation**: `Compose` passes cancellation through to later arrows.

## Known gaps

The universal properties of the product and coproduct aren't property-tested yet. For example: `Compose(Par(f, g), Fst) ≡ f` when `g` succeeds, and `Compose(Left-injection, Sum(f, g)) ≡ f`. Adding them is a good first contribution.

## What the laws don't cover

The laws describe values and errors, not side effects. An arrow that sends an email does so every time it runs, and retries or fallbacks can make that happen more than once. Make effectful arrows idempotent, or keep them out of `WithRetry`.
