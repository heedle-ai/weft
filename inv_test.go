package weft

import (
	"os/exec"
	"strings"
	"testing"
)

// TestInv_StandardLibraryOnly: weft depends on nothing outside Go's
// standard library, tests included, (asked of the toolchain with go list).
func TestInv_StandardLibraryOnly(t *testing.T) {
	out, err := exec.Command("go", "list", "-deps", "-test", "-f", "{{if not .Standard}}{{.ImportPath}}{{end}}", "./...").Output()
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range strings.Fields(string(out)) {
		// The package itself and its test variants ("…/weft […]", "…/weft_test", "…/weft.test").
		if p != "github.com/heedle-ai/weft" && !strings.HasPrefix(p, "github.com/heedle-ai/weft_test") && !strings.HasPrefix(p, "github.com/heedle-ai/weft.test") && !strings.HasPrefix(p, "[") {
			t.Errorf("non-standard dependency: %s", p)
		}
	}
}
