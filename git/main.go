package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Git Tag: %s\nGit Hash: %s\nBuild Date: %s\nBuild Time: %s\n", util.GitTag, util.GitHash, util.BuildDate, util.BuildTime)
}
