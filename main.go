package main

import (
	add "github.com/travisperson/ipfs-api-test/v0/add"
	cat "github.com/travisperson/ipfs-api-test/v0/cat"
	"testing"
	"fmt"
	"flag"
)
func strmatch(pat, str string) (bool, error) {
	fmt.Println(pat, str)
	return true, nil
}

func main () {
	flag.Parse()

	test := []testing.InternalTest {
		{"Get", add.Get},
		{"PostEmpty", add.PostEmpty},
		{"PostFile", add.PostFile},
		{"GetCat", cat.Get},
	}

	bench := []testing.InternalBenchmark {}

	examples := []testing.InternalExample {}

	testing.Main(strmatch, test, bench, examples)
}
