// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(
		http.Dir("./list"),
		vfsgen.Options{
			Filename:     "./vfsdata_disposable.go",
			PackageName:  "disposable",
			VariableName: "asset",
		})
	if err != nil {
		log.Fatalln(err)
	}
}
