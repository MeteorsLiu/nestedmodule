package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func genVersion(fn func(string)) {
	for i := 0; i < 10; i++ {
		major := fmt.Sprintf("%d.", i)
		for j := 0; j < 10; j++ {
			minor := fmt.Sprintf("%d.", j)
			for k := 1; k < 11; k++ {
				patch := fmt.Sprintf("%d", k)

				fn(major + minor + patch)

			}
		}

	}
}

func genDir(version string) {
	dirName := filepath.Join("cjson", version)
	// {Version}
	os.Mkdir(dirName, 0777)

	fileName := filepath.Join(dirName, "go.mod")
	goFileName := filepath.Join(dirName, "hello.go")
	goFile := fmt.Sprintf(`package cjson
func HelloWorld() {
	println("ni howdy %s")
}
				`, version)
	modContent := fmt.Sprintf(`module github.com/MeteorsLiu/nestedmodule/cjson/%s
go 1.23.4
				`, version)
	os.WriteFile(fileName, []byte(modContent), 0777)
	os.WriteFile(goFileName, []byte(goFile), 0777)

	exec.Command("git", "tag", fmt.Sprintf("cjson/%s/v0.1.0", version)).Run()
}

func removeTag(version string) {
	exec.Command("git", "tag", "-d", fmt.Sprintf("cjson@%s", version)).Run()
}

func main() {
	// {LibraryName}
	os.Mkdir("cjson", 0777)

	genVersion(genDir)
}
