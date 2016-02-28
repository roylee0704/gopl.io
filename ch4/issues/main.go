package main

import (
	"fmt"
	"log"
	"os"

	"github.com/roylee0704/gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, issue := range result.Items {
		//%9.9s
		fmt.Printf("#%-5d|%9.9s|%.55s\n", issue.Number, issue.User.Login, issue.Title)
	}

}
