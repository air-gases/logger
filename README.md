# Logger

A useful gas that used to log every request for the web applications built using
[Air](https://github.com/sheng/air).

## Installation

Open your terminal and execute

```bash
$ go get github.com/air-gases/logger
```

done.

> The only requirement is the [Go](https://golang.org), at least v1.8.

## Usage

Create a file named `main.go`

```go
package main

import (
	"github.com/air-gases/logger"
	"github.com/sheng/air"
)

func main() {
	air.Gases = []air.Gas{
		logger.Gas(logger.GasConfig{}),
	}
	air.GET("/", func(req *air.Request, res *air.Response) error {
		return res.String("Go and see what your terminal outputs.")
	})
	air.Serve()
}
```

and run it

```bash
$ go run main.go
```

then visit `http://localhost:2333`.

## Community

If you want to discuss this gas, or ask questions about it, simply post
questions or ideas [here](https://github.com/air-gases/logger/issues).

## Contributing

If you want to help build this gas, simply follow
[this](https://github.com/air-gases/logger/wiki/Contributing) to send pull
requests [here](https://github.com/air-gases/logger/pulls).

## License

This gas is licensed under the Unlicense.

License can be found [here](LICENSE).
