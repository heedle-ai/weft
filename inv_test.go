package weft

import (
	"os/exec"
	"strings"
	"testing"
)

// TestInv_StandardLibraryOnly: weft depends on nothing outside Go's
// standard library (asked of the toolchain with go list).
func TestInv_StandardLibraryOnly(t *testing.T) {
	out, err := exec.Command("go", "list", "-deps", "-f", "{{if not .Standard}}{{.ImportPath}}{{end}}", ".").Output()
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range strings.Fields(string(out)) {
		if p != "github.com/heedle-ai/weft" {
			t.Errorf("non-standard dependency: %s", p)
		}
	}
}
