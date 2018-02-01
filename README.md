[![](https://godoc.org/github.com/ravernkoh/deepl?status.svg)](http://godoc.org/github.com/ravernkoh/deepl)

# DeepL Golang Client
Unofficial Golang client for the undocumented API of [deepl.com](http://deepl.com).

## Usage
```go
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
```

## Installation
```
$ go get github.com/ravernkoh/deepl
```
