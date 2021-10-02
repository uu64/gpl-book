package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	lt1month := []*github.Issue{}
	lt1year := []*github.Issue{}
	ge1year := []*github.Issue{}

	for _, item := range result.Items {
		d := now.Sub(item.CreatedAt)
		switch {
		case d.Hours() < 24*31:
			lt1month = append(lt1month, item)
		case d.Hours() < 24*365:
			lt1year = append(lt1year, item)
		default:
			ge1year = append(ge1year, item)
		}
	}

	fmt.Printf("issues created within 1 month:\n")
	show(lt1month)

	fmt.Printf("issues created within 1 year:\n")
	show(lt1year)

	fmt.Printf("issues created before 1 year:\n")
	show(ge1year)
}

func show(items []*github.Issue) {
	for _, item := range items {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02 15:04:05 MST"), item.User.Login, item.Title)
	}
	fmt.Printf("%d issues\n\n", len(items))
}
