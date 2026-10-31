package weft_test

// Runnable examples. They appear on pkg.go.dev and are checked by `go test`:
// each // Output comment must match exactly. The wiki quotes these, so keep
// the two in sync.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/heddle/weft"
)

func ExampleCompose() {
	greet := weft.Pure(func(name string) string { return "Hello, " + name })
	shout := weft.Pure(strings.ToUpper)

	loud := weft.Compose(greet, shout)
	out, err := loud(context.Background(), "world")
	fmt.Println(out, err)
	// Output: HELLO, WORLD <nil>
}

func ExamplePipe3() {
	parse := weft.ArrowFunc(func(_ context.Context, s string) (int, error) {
		return strconv.Atoi(strings.TrimSpace(s))
	})
	double := weft.Pure(func(n int) int { return 2 * n })
	render := weft.Pure(func(n int) string { return "result=" + strconv.Itoa(n) })

	pipeline := weft.Pipe3(parse, double, render)

	out, err := pipeline(context.Background(), " 21 ")
	fmt.Println(out, err)

	_, err = pipeline(context.Background(), "not a number")
	fmt.Println(err != nil)
	// Output:
	// result=42 <nil>
	// true
}

func ExamplePar() {
	length := weft.Pure(func(s string) int { return len(s) })
	upper := weft.Pure(strings.ToUpper)

	both := weft.Par(length, upper)
	p, err := both(context.Background(), "weft")
	fmt.Println(p.Fst, p.Snd, err)
	// Output: 4 WEFT <nil>
}

func ExampleSum() {
	fromInt := weft.Pure(func(n int) string { return "int " + strconv.Itoa(n) })
	fromStr := weft.Pure(func(s string) string { return "string " + s })

	route := weft.Sum(fromInt, fromStr)
	a, _ := route(context.Background(), weft.Left[int, string](7))
	b, _ := route(context.Background(), weft.Right[int, string]("seven"))
	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// int 7
	// string seven
}

func ExampleTraverse() {
	square := weft.Pure(func(n int) int { return n * n })

	all := weft.Traverse(square, weft.WithConcurrency(4))
	out, err := all(context.Background(), []int{1, 2, 3, 4})
	fmt.Println(out, err)
	// Output: [1 4 9 16] <nil>
}

func ExampleApply() {
	calls := 0
	flaky := weft.ArrowFunc(func(_ context.Context, s string) (string, error) {
		calls++
		if calls < 3 {
			return "", errors.New("temporary failure")
		}
		return "ok: " + s, nil
	})

	robust := weft.Apply(flaky,
		weft.WithRetry[string, string](5, weft.LinearBackoff(time.Millisecond)),
		weft.WithTimeout[string, string](time.Second),
	)
	out, err := robust(context.Background(), "ping")
	fmt.Println(out, err, "after", calls, "calls")
	// Output: ok: ping <nil> after 3 calls
}

func ExampleFallback() {
	primary := weft.ArrowFunc(func(context.Context, string) (string, error) {
		return "", errors.New("primary down")
	})
	backup := weft.Pure(func(s string) string { return "backup handled " + s })

	out, err := weft.Fallback(primary, backup)(context.Background(), "req-1")
	fmt.Println(out, err)
	// Output: backup handled req-1 <nil>
}
