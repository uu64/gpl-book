package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/uu64/gpl-book/ch04/ex11/github/issue"
	"gopl.io/ch4/github"
)

const usage = "The simple cli tool to manage gh issues.\n\n" +
	"USAGE\n" +
	"  List issues in this repository. State is either open or closed. Default is open.\n" +
	"  $ issues list [state]\n\n" +
	"  Show the detail of the specified issue.\n" +
	"  $ issues show <issue_number>\n\n" +
	"  Create a new issue.\n" +
	"  $ issues create\n\n" +
	"  Edit an issue.\n" +
	"  $ issues edit <issue_number>\n\n" +
	"  Close an issue.\n" +
	"  $ issues close <issue_number>\n\n" +
	"  Search for issues from github.\n" +
	"  $ issues search [...terms]\n\n"

type Cmd struct {
	owner string
	repo  string
}

func (cmd *Cmd) help() {
	fmt.Println(usage)
}

func (cmd *Cmd) search(terms []string) error {
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

func (cmd *Cmd) list(state string) error {
	result, err := issue.List(cmd.owner, cmd.repo, state)
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

func (cmd *Cmd) show(number string) error {
	result, err := issue.Show(cmd.owner, cmd.repo, number)
	if err != nil {
		return err
	}

	dateFmt := "2006-01-02 15:04:05 MST"
	fmt.Printf("%s #%d\n", result.Title, result.Number)
	fmt.Printf("state: %s (%s opened at %s)\n\n",
		result.State, result.User.Login, result.CreatedAt.Format(dateFmt))
	fmt.Printf("%s\n", result.Body)

	return nil
}

func (cmd *Cmd) create() {}

func (cmd *Cmd) edit() {}

func (cmd *Cmd) close(number string) error {
	result, err := issue.Close(cmd.owner, cmd.repo, number)
	if err != nil {
		return err
	}
	fmt.Printf("#%d is closed.\n", result.Number)
	return nil
}

func newCmd() (*Cmd, error) {
	var owner, repo string
	out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return nil, fmt.Errorf("get git config failed: %w", err)
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
		return nil, fmt.Errorf("unexpected url format: %s", url)
	}

	return &Cmd{owner, repo}, nil
}

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

	command := args[0]
	switch command {
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
		fmt.Println(command)
	case "edit":
		fmt.Println(command)
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
		return fmt.Errorf("invalid command: %s", command)
	}

	return nil
}
