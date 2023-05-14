go-unidecode
==============

[![Build Status](https://github.com/mozillazg/go-unidecode/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/mozillazg/go-unidecode/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/mozillazg/go-unidecode/badge.svg?branch=master)](https://coveralls.io/r/mozillazg/go-unidecode?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mozillazg/go-unidecode)](https://goreportcard.com/report/github.com/mozillazg/go-unidecode)
[![GoDoc](https://godoc.org/github.com/mozillazg/go-unidecode?status.svg)](https://godoc.org/github.com/mozillazg/go-unidecode)

ASCII transliterations of Unicode text. Inspired by [python-unidecode](https://github.com/avian2/unidecode).


Installation
------------

```
go get github.com/mozillazg/go-unidecode
```

Install CLI tool:

```
$ go install github.com/mozillazg/go-unidecode/cmd/unidecode@latest

$ unidecode 北京kožušček
Bei Jing kozuscek
```


Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/mozillazg/go-unidecode


Usage
------

```go
package main

import (
	"fmt"
	"github.com/mozillazg/go-unidecode"
)

func main() {
	s := "abc"
	fmt.Println(unidecode.Unidecode(s))
	// Output: abc

	s = "北京"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing

	s = "kožušček"
	fmt.Println(unidecode.Unidecode(s))
	// Output: kozuscek
}
```
