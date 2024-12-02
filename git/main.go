package main

import (
	"fmt"
	"sudhagar/glad/pkg/util"
)

func main() {
	fmt.Printf("Git Tag: %s\nGit Hash: %s\nBuild Date: %s\nBuild Time: %s\n", util.GitTag, util.GitHash, util.BuildDate, util.BuildTime)
}
