# Logger

[![GoDoc](https://godoc.org/github.com/air-gases/logger?status.svg)](https://godoc.org/github.com/air-gases/logger)

A useful gas that used to log every request for the web applications built using
[Air](https://github.com/aofei/air).

## Installation

Open your terminal and execute

```bash
$ go get github.com/air-gases/logger
```

done.

> The only requirement is the [Go](https://golang.org), at least v1.12.

## Usage

Create a file named `main.go`

```go
package main

import (
	"github.com/air-gases/logger"
	"github.com/aofei/air"
)

func main() {
	a := air.Default
	a.DebugMode = true
	a.Pregases = []air.Gas{
		logger.Gas(logger.GasConfig{}),
	}
	a.GET("/", func(req *air.Request, res *air.Response) error {
		return res.WriteString("Go and see what your terminal outputs.")
	})
	a.Serve()
}
```

and run it

```bash
$ go run main.go
```

then visit `http://localhost:8080`.

## Community

If you want to discuss Logger, or ask questions about it, simply post questions
or ideas [here](https://github.com/air-gases/logger/issues).

## Contributing

If you want to help build Logger, simply follow
[this](https://github.com/air-gases/logger/wiki/Contributing) to send pull
requests [here](https://github.com/air-gases/logger/pulls).

## License

This project is licensed under the Unlicense.

License can be found [here](LICENSE).
