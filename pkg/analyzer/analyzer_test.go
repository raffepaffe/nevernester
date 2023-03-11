package analyzer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNeverNesterPass(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "examples")
	analysistest.Run(t, testdata, New(), "below_five")
}

func TestShouldFail(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "examples")

	results := analysistest.Run(t, testdata, New(), "above_four")
	assert.NotNil(t, results)
	assert.Contains(t, results[0].Diagnostics[0].Message, "max")
}
