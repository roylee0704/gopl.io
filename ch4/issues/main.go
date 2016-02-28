package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/roylee0704/gopl.io/ch4/github"
)

type ageReport time.Duration

func (a ageReport) String() string {

	days := int(time.Duration(a).Hours() / 24)

	var ret string
	switch {
	case days < 31:
		ret = "< 1 month"
	case days < 365:
		ret = "< 1 year"
	default:
		ret = "> 1 year"
	}
	return ret
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, issue := range result.Items {

		var a ageReport = ageReport(time.Since(issue.CreatedAt))
		fmt.Printf("#%-5d|%9s|%9.9s|%.55s\n", issue.Number, a, issue.User.Login, issue.Title)
	}

}
