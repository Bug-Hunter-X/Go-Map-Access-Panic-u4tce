```go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int) // Initialize the map
    value, ok := m["key"]
    if ok {
        fmt.Println("Value for key 'key':", value)
    } else {
        fmt.Println("Key 'key' not found in the map")
    }
}
```