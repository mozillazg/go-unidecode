go-unidecode
==============

[![Build Status](https://travis-ci.org/mozillazg/go-unidecode.svg?branch=master)](https://travis-ci.org/mozillazg/go-unidecode)
[![Coverage Status](https://coveralls.io/repos/mozillazg/go-unidecode/badge.png?branch=master)](https://coveralls.io/r/mozillazg/go-unidecode?branch=master)
[![GoDoc](https://godoc.org/github.com/mozillazg/go-unidecode?status.svg)](https://godoc.org/github.com/mozillazg/go-unidecode)

ASCII transliterations of Unicode text.


Installation
------------

```
go get -u github.com/mozillazg/go-unidecode
```

install CLI tool:

```
go get -u github.com/mozillazg/go-unidecode/unidecode
$ unidecode 北京
Bei Jing 
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
	s = "abc"
	fmt.Println(unidecode.Unidecode(s))
	// Output: abc

	s = "北京"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing

	s = "30 𝗄𝗆/𝗁"
	fmt.Println(unidecode.Unidecode(s))
	// Output: 30 km/h

	s = "kožušček"
	fmt.Println(unidecode.Unidecode(s))
	// Output: kozuscek
}
```
