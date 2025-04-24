package main

import (
	"github.com/noornee/toolkits/cmd"
	_ "github.com/noornee/toolkits/cmd/redis"
)

func main() {
	cmd.Execute()
}
