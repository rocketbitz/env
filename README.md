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
    env.SetDefault("DEFAULT_ENV", "default_value")
    env.Get("DEFAULT_ENV")  // "default_value"

    env.Set("DEFAULT_ENV", "explicit_value")
    env.Get("DEFAULT_ENV")  // "explicit_value"

    env.Count() // 1

    env.JSON()
    /*
    [
        {
            "key": "DEFAULT_ENV",
            "value": "explicit_value",
        }
    ]
    */
}
```

# contribute

pr's are welcome. if they're awesome, they'll get reviewed and merged. if they're not, they'll get reviewed and closed, hopefully with a kind comment as to the reason.

# license

[MIT](https://github.com/rocketbitz/env/blob/master/LICENSE) ...move along, move along.
