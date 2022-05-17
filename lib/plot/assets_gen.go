//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/innius/vegeta/v12/lib/plot"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(plot.Assets, vfsgen.Options{
		PackageName:  "plot",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})

	if err != nil {
		log.Fatalln(err)
	}
}
