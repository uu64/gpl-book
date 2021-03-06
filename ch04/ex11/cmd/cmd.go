package cmd

import (
	"bufio"
	"fmt"
	"os"
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
	"  $ issues search [...terms]\n\n" +
	"ENV VARS\n" +
	"  EDITOR: Set the command name to use for authoring text. Default is \"vim\".\n" +
	"  GH_ACCESS_TOKEN: Set the github personal access token.\n\n"

// Cmd is issues command
type Cmd struct {
	owner string
	repo  string
}

func newCmd() (*Cmd, error) {
	var owner, repo string
	out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return nil, fmt.Errorf("get git config failed: %w", err)
	}

	// remove '\n'
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

func (cmd *Cmd) create() error {
	var title, body []byte
	var err error

	fmt.Printf("Title: ")
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		title = in.Bytes()
	}

	fmt.Printf("Body: (please press ENTER to launch editor)")
	if in.Scan() {
		body, err = edit(nil)
	}
	if err != nil {
		return err
	}

	result, err := issue.Create(cmd.owner, cmd.repo, title, body)
	if err != nil {
		return err
	}
	fmt.Printf("#%d is opened.\n", result.Number)
	return nil
}

func (cmd *Cmd) edit(number string) error {
	var title, body []byte
	var err error

	current, err := issue.Show(cmd.owner, cmd.repo, number)
	if err != nil {
		return err
	}

	fmt.Printf("Title (%s): ", current.Title)
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		title = in.Bytes()
	}
	if len(title) == 0 {
		title = []byte(current.Title)
	}

	fmt.Printf("Body: (please press ENTER to launch editor)")
	if in.Scan() {
		body, err = edit(&current.Body)
	}
	if err != nil {
		return err
	}

	result, err := issue.Update(cmd.owner, cmd.repo, fmt.Sprintf("%d", current.Number), title, body)
	if err != nil {
		return err
	}
	fmt.Printf("#%d is updated.\n", result.Number)
	return nil
}

func (cmd *Cmd) close(number string) error {
	result, err := issue.Close(cmd.owner, cmd.repo, number)
	if err != nil {
		return err
	}
	fmt.Printf("#%d is closed.\n", result.Number)
	return nil
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
