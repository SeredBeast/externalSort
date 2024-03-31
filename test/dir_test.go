package test

import (
	"fmt"
	"os"
	"testing"
)

func TestDirLen(t *testing.T) {
	entries, _ := os.ReadDir("chunks")
	fmt.Println(len(entries))
}
