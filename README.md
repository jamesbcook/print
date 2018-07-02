# Print Package

## Example

```go
package main

import (
    "fmt"
    "io"
    "os"

    "github.com/jamesbcook/print"
)

func main() {
    print.Goodln("Good Line")
    print.Goodf("Good printf %s\n", "Some string")
    print.Statusln("Status Line")
    print.Statusf("Status printf %s\n", "Some string")
    print.Warningln("Warning Line")
    print.Warningf("Warning printf %s\n", "Some string")
    var output io.Writer
    output = os.Stdout
    debugWriter := print.Debugln(true, &output)
    debugWriter("Hello World")
    errorWriter := print.Error(&output)
    errorWriter(fmt.Errorf("Hello World"))
    print.Badln("Bad Line")
}
```

## Terminal Output

![terminal](https://github.com/jamesbcook/print/raw/master/media/example-usage.png)