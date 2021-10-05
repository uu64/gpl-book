package cmd

import "fmt"

// Run executes this command
func Run(args []string) error {
	cmd, err := newCmd()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if len(args) < 1 {
		cmd.help()
		return fmt.Errorf("check the usage")
	}

	switch args[0] {
	case "list":
		state := "open"
		if len(args) == 2 {
			state = args[1]
		}
		if err := cmd.list(state); err != nil {
			return fmt.Errorf("list: %w", err)
		}
	case "show":
		if len(args) < 2 {
			return fmt.Errorf("show: issue number is required")
		}
		if err := cmd.show(args[1]); err != nil {
			return fmt.Errorf("show: %w", err)
		}
	case "create":
		cmd.create()
	case "edit":
	case "close":
		if len(args) < 2 {
			return fmt.Errorf("close: issue number is required")
		}
		if err := cmd.close(args[1]); err != nil {
			return fmt.Errorf("close: %w", err)
		}
	case "search":
		if len(args) < 2 {
			return fmt.Errorf("search: terms are required")
		}
		if err := cmd.search(args[1:]); err != nil {
			return fmt.Errorf("search: %w", err)
		}
	default:
		cmd.help()
		return fmt.Errorf("check the usage")
	}

	return nil
}
