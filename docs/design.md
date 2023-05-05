# WeiwuCache Design

## How to use
**Interactive interface**
```bash
./cache --port=6666
127.0.0.1:6666> set weiwu caozei
OK
127.0.0.1:6666> get weiwu
caozei
```

**As a third-party library**
```go
import (
    "fmt"
	
    wcache "github.com/Uuq114/WeiwuCache"
)
func main() {
    cache := wcache.New[string, string](1000)
    cache.Set("weiwu", "caozei")
    value, ok := cache.Get("weiwu")
    if ok {
        fmt.Println(value)	
    }
}
```

## Internal Design
**Overview**

Request => Storage => Response

**Request Handler**

//todo

**Underlying Storage**

//todo

**Response exporter**

//todo

## 