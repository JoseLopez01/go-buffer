<p align="center">
  <img src="gopher.png">
</p>
<p align="center">
  <img src="https://img.shields.io/github/workflow/status/globocom/go-buffer/Go?style=flat-square">
  <a href="https://github.com/globocom/go-buffer/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/globocom/go-buffer?color=blue&style=flat-square">
  </a>
  <img src="https://img.shields.io/github/go-mod/go-version/globocom/go-buffer?style=flat-square">
  <a href="https://pkg.go.dev/github.com/globocom/go-buffer">
    <img src="https://img.shields.io/badge/Go-reference-blue?style=flat-square">
  </a>
</p>

# go-buffer

`go-buffer` represents a buffer that asynchronously flushes its contents. It is useful for applications that need to aggregate data before writing it to an external storage. A buffer is flushed manually, or automatically when it becomes full or after an interval has elapsed, whichever comes first.

## Installation

    go get github.com/globocom/go-buffer

## Examples

### Size-triggered flush

```golang
package main

import (
    "fmt"
    "time"

    "github.com/globocom/go-buffer/v3"
)

func main() {
    buff := buffer.New[string](
        // the Flusher.Write method is called when the buffer need to be flushed
        func(items []string) {
            for _, item := range items {
                fmt.Println(item)
            }
        },
        // buffer cab hold up to 5 items
        buffer.WithSize(5),
    )
    defer buff.Close()

    buff.Push("item 1")
    buff.Push("item 2")
    buff.Push("item 3")
    buff.Push("item 4")
    buff.Push("item 5")

    time.Sleep(3 * time.Second)
    fmt.Println("done")
}

```

### Interval-triggered flush

```golang
package main

import (
    "fmt"
    "time"

    "github.com/globocom/go-buffer/v3"
)

func main() {
    buff := buffer.New[string](
        func(items []string) {
            for _, item := range items {
                fmt.Println(item)
            }
        },
        buffer.WithSize(5),
        // buffer will be flushed every second, regardless of
        // how many items were pushed
        buffer.WithFlushInterval(time.Second),
    )
    defer buff.Close()

    buff.Push("item 1")
    buff.Push("item 2")
    buff.Push("item 3")

    time.Sleep(3 * time.Second)
    fmt.Println("done")
}

```

### Manual flush

```golang
package main

import (
    "fmt"
    "time"

    "github.com/globocom/go-buffer/v3"
)

func main() {
    buff := buffer.New[string](
        func(items []string) {
            for _, item := range items {
                fmt.Println(item)
            }
        },
        buffer.WithSize(5),
    )
    defer buff.Close()

    buff.Push("item 1")
    buff.Push("item 2")
    buff.Push("item 3")

    time.Sleep(3 * time.Second)
    buff.Flush()
    fmt.Println("done")
}

```

## Documentation

Visit [Pkg.go.dev](https://pkg.go.dev/github.com/globocom/go-buffer) for full documentation.

## License

[MIT License](https://github.com/globocom/go-buffer/blob/master/LICENSE)
