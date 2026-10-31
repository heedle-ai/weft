## 1. Install Go 1.23+

**macOS (Apple Silicon)**

```bash
brew install go
go version
```

**Linux (amd64 or arm64)**

```bash
GO_VERSION=1.24.7
ARCH=$(uname -m | sed -e 's/x86_64/amd64/' -e 's/aarch64/arm64/')
curl -fsSLO "https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz"
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf "go${GO_VERSION}.linux-${ARCH}.tar.gz"
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile && source ~/.profile
go version
```

Distribution packages (`apt install golang`) are often too old. Check `go version` reports 1.23 or later.

## 2. Use weft in your module

```bash
go get github.com/heedle-ai/weft
```

```go
package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/heedle-ai/weft"
)

func main() {
	greet := weft.Pure(func(name string) string { return "Hello, " + name })
	shout := weft.Pure(strings.ToUpper)

	loud := weft.Compose(greet, shout)
	out, err := loud(context.Background(), "world")
	fmt.Println(out, err) // HELLO, WORLD <nil>
}
```

This is `ExampleCompose` in `doc_examples_test.go`, so `go test` checks it.

## 3. Work on weft itself

heddle's repos refer to each other by relative path, so clone them as siblings:

```bash
mkdir -p ~/src/heddle && cd ~/src/heddle
git clone https://github.com/heedle-ai/weft.git
cd weft
go vet ./...
go test -race ./...
```

Run only the laws, or only the runnable examples:

```bash
go test -run 'Law|Equivalent|Preserves' -v ./...
go test -run Example -v ./...
```

Next: [Combinators](Combinators).
