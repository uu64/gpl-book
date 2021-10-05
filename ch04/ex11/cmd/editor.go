package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func newEditor() (cmd *exec.Cmd, filename string, err error) {
	name := "vim"
	if e := os.Getenv("EDITOR"); e != "" {
		name = e
	}

	f, err := os.CreateTemp("", "issues.body.*.md")
	defer closeFile(f)
	if err != nil {
		return nil, "", fmt.Errorf("failed to make temp file: %w", err)
	}

	editor := exec.Command(name, f.Name())
	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout
	editor.Stderr = os.Stderr
	return editor, f.Name(), nil
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
