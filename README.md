# set [![CircleCI](https://circleci.com/gh/rocketbitz/env/tree/master.svg?style=svg)](https://circleci.com/gh/rocketbitz/env/tree/master)
an easier to use environment variable package for gophers

# installation

as with all other go packages, you know what to do:
```
go get github.com/rocketbitz/env
```

# usage

go should probably have its own, but it doesn't, so here you are:

```go
package main

import (
	"fmt"

	"github.com/rocketbitz/env"
)

func main() {
    s := set.New()

    s.Add("gopher_0")
    s.Add("gopher_1")

    s.Contains("gopher_0")  // true
    s.Contains("gopher_2")  // false

    s.Slice()  // []interface{}{"gopher_0", "gopher_2"}

    s.Remove("gopher_0")

    s.Len()  // 1

    s.At(0)  // "gopher_2"
    s.Index("gopher_2")  // 0
}
```

# contribute

pr's are welcome. if they're awesome, they'll get reviewed and merged. if they're not, they'll get reviewed and closed, hopefully with a kind comment as to the reason.

# license

[MIT](https://github.com/rocketbitz/env/blob/master/LICENSE) ...move along, move along.
