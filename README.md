# Go API for [dynamon.io](https://dynamon.io)


## Install

```
go get github.com/dynamon-io/dynamon-go
```


## Usage

```go
package main

import "github.com/dynamon-io/dynamon-go"

func main() {
    xs := []float64{1, 2, 3}
    ys := []float64{3, 2, 4}
    dynamon.Line("myNamespace", xs, ys, dynamon.LineOpts{})
}
```
