package stringreverse

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s1 := []byte("hello world")
	ReverseString(s1)
	fmt.Println(string(s1))

	s2 := []byte("i love world")
	ReverseString(s2)
	fmt.Println(string(s2))
}
