package main

import (
	"fmt"

	"github.com/ravernkoh/deepl"
)

func main() {
	// Instantiate client and configure
	cli := deepl.NewClient()

	// Translate some simple text
	res, _ := cli.Translate([]string{"Hello, world!"}, deepl.English, deepl.German)
	fmt.Println(res[0])
}
