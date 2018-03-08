package main

import (
	"flag"
	"fmt"
	"gitlab.xinghuolive.com/Backend-Go/kangaroo/wiki"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var tomlFile = flag.String("toml", "", "the toml file path")

func main() {

	flag.Parse()

	if *tomlFile == "" {
		fmt.Println("need toml file path")
		return
	}

	if strings.HasPrefix(*tomlFile, "$GOPATH") {

		gopath := os.Getenv("GOPATH")
		dirs := strings.Split(gopath, ":")

		if runtime.GOOS == "windows" {
			dirs = strings.Split(gopath, ";")
		}

		for _, d := range dirs {
			filePath := strings.Replace(*tomlFile, "$GOPATH", d, 1)
			if _, err := os.Stat(filePath); err == nil {
				*tomlFile = filePath
				break
			}
		}
	}

	fmt.Println("tomlFile: ", *tomlFile)

	err := wiki.ParseToml(*tomlFile)
	if err != nil {
		fmt.Println(" parse toml file error ")
	}

	if wiki.WikiOutDir != "" {
		wiki.ParseToOutDir(wiki.Pages, wiki.WikiOutDir)
	} else if wiki.WikiPackageName != "" {
		tomlPath, err := filepath.Abs(*tomlFile)
		if err != nil {
			fmt.Println("get current directory error.")
		}
		dir := wiki.GetParentDirectory(tomlPath) + "/markdown/"
		wiki.ParseToOutDir(wiki.Pages, dir)
	} else {
		fmt.Println("must set out dir")
	}

}
