// String operations logics are same for UTF-8 encoded text. This is not true for
// other encodings.
package main

import "fmt"

func main() {
	fmt.Printf("HasPrefix(%s, %s) = %v\n", "world, hello", "hello", HasPrefix("world, hello", "hello"))
	fmt.Printf("HasPrefix(%s, %s) = %v\n", "hello, world", "hello", HasPrefix("hello, world", "hello"))

	fmt.Printf("HasSuffix(%s, %s) = %v\n", "world, hello", "hello", HasSuffix("world, hello", "hello"))
	fmt.Printf("HasSuffix(%s, %s) = %v\n", "hello, world", "hello", HasSuffix("hello, world", "hello"))

	fmt.Printf("Contains(%s, %s) = %v\n", "world, hello", "hello", Contains("world, hello", "hello"))
	fmt.Printf("Contains(%s, %s) = %v\n", "hello, world", "hello", Contains("hello, world", "hello"))

	fmt.Printf("Contains(%s, %s) = %v\n", "你haha们好，世界", "们好", Contains("你haha们好，世界", "haha们好"))
	fmt.Printf("HasPrefix(%s, %s) = %v\n", "你们好，世界", "你们", Contains("你们好，世界", "你们"))
	fmt.Printf("HasSuffix(%s, %s) = %v\n", "你们好，世界", "世界", Contains("你们好，世界", "世界"))

}
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
