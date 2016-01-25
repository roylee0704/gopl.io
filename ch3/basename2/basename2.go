// basename2 removes directory components and a .suffix
package main

import "fmt"

func main() {
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("a.go"))     // "a"
	fmt.Println(basename("abc"))      // "abc"
}

func basename(s string) string {
	// Discard last '/' and everything before
	slash := lastIndex(s, "/") //-1 if "/" not found
	s = s[slash+1:]            // if -1, nothing happens

	if dot := lastIndex(s, "."); dot != -1 {
		s = s[:dot]
	}
	return s

	// Preserve everything before last '.'

}

// LastIndex returns the index of last instance of sep in S or -1 if sep is not
// present in s.
func lastIndex(s, sep string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if hasSuffix(s[i:i+len(sep)], sep) {
			return i
		}
	}
	return -1
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
