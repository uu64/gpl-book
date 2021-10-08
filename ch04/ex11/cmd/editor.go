package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func edit(content *string) ([]byte, error) {
	name := "vim"
	if e := os.Getenv("EDITOR"); e != "" {
		name = e
	}

	f, err := os.CreateTemp("", "issues.body.*.md")
	defer closeFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to make temp file: %w", err)
	}

	if content != nil {
		_, err = f.WriteString(*content)
		if err != nil {
			return nil, fmt.Errorf("failed to write to temp file: %w", err)
		}
	}

	editor := exec.Command(name, f.Name())
	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout
	editor.Stderr = os.Stderr

	err = editor.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to edit: %w", err)
	}

	b, err := os.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return b, nil
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
