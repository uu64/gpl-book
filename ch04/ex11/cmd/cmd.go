package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/uu64/gpl-book/ch04/ex11/github/issue"
	"gopl.io/ch4/github"
)

func search(terms []string) error {
	result, err := github.SearchIssues(terms)
	if err != nil {
		return err
	}

	fmt.Printf("%d issues:\n", len(result.Items))
	for _, item := range result.Items {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02"), item.User.Login, item.Title)
	}

	return nil
}

func list(owner, repo string) error {
	result, err := issue.List(owner, repo)
	if err != nil {
		return err
	}

	fmt.Printf("%d issues:\n", len(*result))
	for _, item := range *result {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02"), item.User.Login, item.Title)
	}

	return nil
}

func show() {}

func create() {}

func edit() {}

func close() {}

// Run executes this command
func Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("command is required")
	}

	owner, repo, err := getRepoInfo()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	command := args[0]
	switch command {
	case "search":
		if len(args) < 2 {
			return fmt.Errorf("search: terms are required")
		}
		if err := search(args[1:]); err != nil {
			return fmt.Errorf("search: %w", err)
		}
	case "list":
		if err := list(owner, repo); err != nil {
			return fmt.Errorf("list: %w", err)
		}
	case "create":
		fmt.Println(command)
	case "show":
		fmt.Println(command)
	case "edit":
		fmt.Println(command)
	case "close":
		fmt.Println(command)
	default:
		return fmt.Errorf("invalid command: %s", command)
	}

	return nil
}

func getRepoInfo() (string, string, error) {
	var owner, repo string
	out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return owner, repo, err
	}

	// remove \n
	url := string(out[:len(out)-1])
	switch {
	case strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "ssh://"):
		tmp := strings.Split(url, "/")
		owner = tmp[len(tmp)-2]
		repo = tmp[len(tmp)-1]
	case strings.HasPrefix(url, "git@"):
		tmp := strings.Split(url, ":")
		tmp = strings.Split(tmp[1], "/")
		owner = tmp[0]
		repo = tmp[1]
	default:
		return owner, repo, fmt.Errorf("unexpected url format: %s", url)
	}
	return owner, repo, nil
}
