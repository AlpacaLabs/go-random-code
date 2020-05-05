# go-random-code
Go generator for random codes

## Getting started
```bash
go get -v github.com/AlpacaLabs/go-random-code
```

and in your Go code
```go
import (
    code "github.com/AlpacaLabs/go-random-code"
)

func main() {
    // Generate a random, six-char code
    code.New()
}
```
