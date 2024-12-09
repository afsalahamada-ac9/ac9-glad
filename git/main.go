package main

import (
	"fmt"
	"ac9/glad/pkg/util"
)

var (
	gitTag  string
	gitHash string
)

func main() {
	fmt.Printf("Git Tag: %s\nGit Hash: %s\nBuild Date: %s\nBuild Time: %s\n", gitTag, gitHash, util.BuildDate, util.BuildTime)
}
