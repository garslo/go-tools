package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	usage bool
)

func init() {
	flag.StringVar(&usage, "usage", false, "Display usage")
}

func dieIf(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func main() {
	if usage {
		fmt.Printf(`gopkg

Prints the current package relative to $GOPATH
`)
		os.Exit(0)
	}

	gopath := os.Getenv("GOPATH")
	pwd, err := os.Getwd()
	dieIf(err)
	packageName, err := getPackageName(gopath, pwd)
	dieIf(err)
	fmt.Println(packageName)
}

func getPackageName(gopath, pwd string) (string, error) {
	if len(gopath) >= len(pwd) {
		return "", fmt.Errorf("'%s' (pwd) is not a subdir of (GOPATH) '%s'", pwd, gopath)
	}
	for i, ch := range gopath {
		if ch != []rune(pwd)[i] {
			return "", fmt.Errorf("'%s' (pwd) is not a subdir of GOPATH '%s'", pwd, gopath)
		}
	}
	// We get "$GOPATH/src/<package-path>"; trim off the $GOPATH and "/src/"
	return pwd[len(gopath)+5:], nil
}
