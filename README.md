# go-textwrap

Is a port of the [python's library](https://docs.python.org/3/library/textwrap.html).


[![GoDoc](https://godoc.org/github.com/ekalinin/go-textwrap?status.svg)](https://godoc.org/github.com/ekalinin/go-textwrap)
[![Tests](https://github.com/ekalinin/go-textwrap/workflows/Tests/badge.svg)](https://github.com/ekalinin/go-textwrap/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/ekalinin/go-textwrap)](https://goreportcard.com/report/github.com/ekalinin/go-textwrap)

# TOC

    * [How to install](#how-to-install)
    * [API](#api)
        * [Detent](#detent)
        * [Indent](#indent)


# How to install

```sh
$ go get github.com/ekalinin/go-textwrap
```

# API

## Detent

Documentation:

- https://pkg.go.dev/github.com/ekalinin/go-textwrap#Dedent

The code:

```go
// ➜ cat main.go
package main

import (
	"fmt"

	"github.com/ekalinin/go-textwrap"
)

func main() {
	txt := `
		select *
		  from products
		 where price > 100;
	`
	fmt.Println(">> Just text:")
	fmt.Println(txt)

	fmt.Println(">> Width Dedent:")
	fmt.Println(textwrap.Dedent(txt))
}
```

The result:

```sh
➜ go run main.go
>> Just text:

		select *
		  from products
		 where price > 100;

>> Width Dedent:

select *
  from products
 where price > 100;
```

## Indent

Documentation:

- https://pkg.go.dev/github.com/ekalinin/go-textwrap#Indent

The code:

```go
// ➜ cat main.go
package main

import (
	"fmt"

	"github.com/ekalinin/go-textwrap"
)

func main() {
	txt := `
    select *
      from products
     where price > 100;
	`
	fmt.Println(">> Just text:")
	fmt.Println(txt)

	fmt.Println(">> Width Indent:")
	fmt.Println(textwrap.Indent(txt, "++", nil))
}
```

The result:

```sh
➜ go run main.go
>> Just text:

    select *
      from products
     where price > 100;

>> Width Indent:

++    select *
++      from products
++     where price > 100;
```
