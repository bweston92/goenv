[![Build Status](https://travis-ci.org/bweston92/goenv.svg?branch=master)](https://travis-ci.org/bweston92/goenv)
[![codecov](https://codecov.io/gh/bweston92/goenv/branch/master/graph/badge.svg)](https://codecov.io/gh/bweston92/goenv)

# Environment helpers for Golang

Uses the `os.Getenv` but supports file resolving and base64 encoded values out the box.

# Example

```go
package main

import (
    "fmt"
    "github.com/bweston92/goenv"
)

func main() {
    fmt.Println("it was meant to be secret but oh well: ", goenv.Get("MYSQL_PASSWORD"))
}
```

# License

MIT License

Copyright (c) 2017 Bradley Weston

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.