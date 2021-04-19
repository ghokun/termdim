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

