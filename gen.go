package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// {LibraryName}
	err := os.Mkdir("cjson", 0777)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		major := fmt.Sprintf("v%d.", i)
		for j := 0; j < 10; j++ {
			minor := fmt.Sprintf("%d.", j)
			for k := 1; k < 11; k++ {
				patch := fmt.Sprintf("%d", k)

				version := major + minor + patch
				dirName := filepath.Join("cjson", version)
				// {Version}
				err := os.Mkdir(dirName, 0777)
				if err != nil {
					log.Fatal(err)
				}
				fileName := filepath.Join(dirName, "go.mod")
				goFileName := filepath.Join(dirName, "hello.go")
				goFile := fmt.Sprintf(`package %s
func HelloWorld() {
	println("ni howdy %s")
}
				`, version, version)
				modContent := fmt.Sprintf(`module github.com/MeteorsLiu/nestedmodule/cjson/%s
go 1.23.4
				`, version)
				os.WriteFile(fileName, []byte(modContent), 0777)
				os.WriteFile(goFile, []byte(goFileName), 0777)

			}
		}
	}
}
