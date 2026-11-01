Every invariant below is guarded by a test or a CI step, and CI checks this page against the code on every push with `bash .github/invariants.sh`. Only the **Guarded by** column counts. Each test named there must exist in exactly one package and is run, and the gate fails if it fails or skips; each `ci:NAME` must be a step in `.github/workflows/ci.yml`. Every `TestLaw_*`, `TestInv_*` and `Fuzz*` test must be listed, and no test file may have build tags, which would hide tests from CI. A mistake isn't fixed until it has a row on this page.

[[/images/invariant-loop.svg|alt=From a mistake to a codified invariant: mistake, named invariant, guard, CI gate]]

| ID | Invariant | Guarded by | Came from |
|---|---|---|---|
| WEFT-1 | Identity: `id ; f = f = f ; id`, including when `f` fails | `TestIdentityLaw_Left`, `TestIdentityLaw_Right`, `TestIdentityLaw_OnFailingArrow` | The category laws |
| WEFT-2 | Associativity: `(f ; g) ; h = f ; (g ; h)`, including errors | `TestAssociativityLaw`, `TestAssociativityLaw_WithErrors` | The category laws |
| WEFT-3 | `Pipe3` and `Pipe4` equal nested `Compose` | `TestPipe3_EquivalentToNestedCompose`, `TestPipe4_EquivalentToNestedCompose` | The category laws |
| WEFT-4 | `Map` is a functor and equals composing with `Pure` | `TestMap_PreservesIdentity`, `TestMap_PreservesComposition`, `TestMap_EquivalentToComposeWithPure` | The category laws |
| WEFT-5 | Product: `Par(f, g) ; Fst = f` and `Par(f, g) ; Snd = g` | `TestLaw_ProductProjections` | A gap the Laws page listed as untested |
| WEFT-6 | Coproduct: `Left ; Sum(f, g) = f` and `Right ; Sum(f, g) = g` | `TestLaw_CoproductInjections` | A gap the Laws page listed as untested |
| WEFT-7 | `ParStrict` returns the error that caused the cancellation, not the sibling's `context canceled` | `TestParStrict_ReturnsTheCauseNotTheCancellation` | Independent review 1, finding 9 |
| WEFT-8 | weft and its tests depend only on Go's standard library | `TestInv_StandardLibraryOnly` | Reuse rule: no dependencies; review 3: test dependencies weren't checked |
| WEFT-9 | Every runnable example's output is exactly as documented | `ExampleCompose`, `ExamplePipe3`, `ExamplePar`, `ExampleSum`, `ExampleTraverse`, `ExampleApply`, `ExampleFallback` | Docs must never claim untested behaviour |
| WEFT-10 | This page is checked strictly against the code: every listed guard exists once and runs without skipping, every `TestLaw_`, `TestInv_` and `Fuzz` test is listed, and no test file has build tags | `ci:invariants` | Review 3: a reviewer satisfied the old gate with a guard named only in prose, a skipped test and a build-tagged test |
