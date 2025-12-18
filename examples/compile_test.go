package examples

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestExamplesCompile(t *testing.T) {
	examplesDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Find all main.go files in subdirectories
	err = filepath.Walk(examplesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == "main.go" {
			dir := filepath.Dir(path)
			relDir, _ := filepath.Rel(examplesDir, dir)

			t.Run(relDir, func(t *testing.T) {
				cmd := exec.Command("go", "build", "-o", "/dev/null", ".")
				cmd.Dir = dir
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("Failed to compile %s:\n%s", relDir, out)
				}
			})
		}
		return nil
	})

	if err != nil {
		t.Fatalf("Error walking examples directory: %v", err)
	}
}
