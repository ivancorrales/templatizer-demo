// Application which greets you.
package main

import (
	"fmt"
	"github.com/{{.RepoOrg}}/{{.RepoName}}/internal/{{.RepoName}}"
)

func main() {
	fmt.Println(greet())
}

func greet() string {
	return {{.RepoName}}.GetMessage()
}
