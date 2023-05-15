# WeiwuCache Design

## How to use
**Interactive interface**
```bash
./cache
> set weiwu caozei
OK
> get weiwu
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

**Communication Protocol**

//todo

**Underlying Storage**
- Data Structure
  - `Cache`: User can declare a `Cache` and handle it
  - `List`: `Cache` can handle either a `List` or a `Hash`, depending on the amount of cached elements
  - `Hash`: When the number of cached elements grows up to `1024`, a `List` will transform into a `Hash`
- Element operation
  - Ordinary CRUD
- Hashtable Expansion
  - Progressive hashing
  - More rehashing patterns are coming!


## 