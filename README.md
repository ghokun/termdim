# termdim

Retrieves terminal dimensions.

## Usage
```go
import (
  "github.com/ghokun/termdim"
)

func MyFunc() {
  width, height, err := termdim.GetSize(int(os.Stdout.Fd()))
}
```
## Installation
```bash
go get github.com/ghokun/termdim
```

## LICENSE
MIT

## Authors
ghokun
