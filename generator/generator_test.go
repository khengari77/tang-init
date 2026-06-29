package generator

import (
	"os"
	"testing"

	"github.com/khengari77/tang-init/profiles"
)

func TestGenerate(t *testing.T) {
	dir := t.TempDir()
	origDir, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(origDir)

	board := profiles.All()[0]
	err := Generate("my_blinky", "blinky", board)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"my_blinky/Makefile",
		"my_blinky/.gitignore",
		"my_blinky/constraints/board.cst",
		"my_blinky/src/top.v",
		"my_blinky/tb/top_tb.v",
	}

	for _, p := range expected {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			t.Errorf("expected file %s does not exist", p)
		}
	}

	makefile, _ := os.ReadFile("my_blinky/Makefile")
	if string(makefile) == "" {
		t.Error("Makefile is empty")
	}
}
