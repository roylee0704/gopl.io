package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	const (
		httpPrefix = "http://"
	)

	for _, url := range os.Args[1:] {

		//exercise1.8: add prefix if its missing from url
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		fmt.Printf("fetching %s\n", url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//fmt.Println(resp) //for fun, ouput in original format (byte slice): %v

		//exercise1.9: print http status code
		fmt.Printf("Resp Status Code: %s", resp.Status)

		//original(url, resp)
		exercise1dot7(url, resp)
	}
}

func original(url string, resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", b) //for fun, ouput in original format (byte slice): %v
}

func exercise1dot7(url string, resp *http.Response) {
	n, err := io.Copy(os.Stdout, resp.Body) //save buffer streams

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	}
	fmt.Printf("%d bytes is read.\n", n)

}
