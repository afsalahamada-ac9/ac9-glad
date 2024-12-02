package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Git Tag: %s\nGit Hash: %s\nBuild Date: %s\nBuild Time: %s\n", util.gitTag, util.gitHash, util.buildDate, util.buildTime)
}
