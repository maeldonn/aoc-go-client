# AOC Go Client

This is a simple client to get the input for a given year and day of Advent of Code.

## Usage

```go
package main

import (
	"fmt"

	"github.com/maeldonn/aoc-go-client"
)

func main() {
	client, err := aocgoclient.NewClient()
	if err != nil {
        log.Fatalf("failed to initialize client: %v", err)
	}

	input, err := client.GetInput(2024, 1)
	if err != nil {
        log.Fatal("failed to get input: ", err)
	}

    // do stuff with the input...
    _ = input
}
```

Then run your script with the `AOC_COOKIE` environment variable set to the value of the `session` cookie.

```sh
$ AOC_COOKIE=your-session-cookie go run main.go
```

## Installation

```sh
$ go get github.com/maeldonn/aoc-go-client
```

## License

MIT
