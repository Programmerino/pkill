# pkill for Golang!

### Extremely simple, but cross-platform function to kill a process by it's process name
```go
import "github.com/SkyrisBactera/pkill"
// Kills chrome
output, err := pkill.Pkill("chrome")
if err != nil {
    fmt.Println(err)
}
fmt.Println(output)
```