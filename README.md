# log

A simple log library based on github.com/uber-go/zap.

## Installation

```shell
$ go get -u github.com/Kaiser925/log
```

## Quick start

Log message

```go
package main

import "github.com/Kaiser925/log"

func main() {
	defer log.Flush()
	log.Info("Hello world")
}
```

Log with custom options

```go
package main

import "github.com/Kaiser925/log"

func main() {
	defer log.Flush()
	log.Init(&log.Option{
		Level: log.DebugLevel.String(),
		Format: log.JsonFormat,
		EnableCaller: true,
		EnableColor: false,
    })
	log.Info("Hello world")
}
```

Released under the [BSD 3-Clause License](./LICENSE)