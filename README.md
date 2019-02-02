# zadapters
Some simple adapters for Zerolog so it can be used with other packages that might define their own Logger interface

## Couchbase

```go
package main

import(
	"os"
	
	"github.com/couchbase/gocbcore"
	"github.com/dcarbone/zadapters/zgocbcore"
	"github.com/rs/zerolog"
)

func main() {
	// init a logger to use 
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Fields(map[string]interface{}{"source": "couchbase"}).
		Logger()
	
	// create the adapter.
	adapter := zgocbcore.NewDefault(log)
	
	// set logger
	gocbcore.SetLogger(adapter)
	
	// presumably do something
}
```

### LevelMap

The [LevelMap](zgocbcore/adapter.go#L10) maps a `gocb.LogLevel` to a `zerolog.Level`.  If you wish a given
Couchbase log level to be ignored, you may either simply not set it or set it to `zerolog.Disabled`.


There is a [DefaultLevelMap](zgocbcore/adapter.go#L22) that will be used if `nil` is passed as first arg to `New()` or
if the [Adapter](zgocbcore/adapter.go#L13) is constructed using [NewDefault()](zgocbcore/adapter.go#L46)

And that's pretty much it.

## Generic

Also included is a generic [Adapter](zstdlog/adapter.go) that allows a ZeroLog logger to be used with 
packages expecting a `*log.Logger` type.  This will also satisfy those packages who create their own `Logger` interface
type with the standard `Print`, `Println`, and `Printf` methods on them.

```go
package main

import(
    "os"

    "github.com/dcarbone/zadapters/zstdlog"
    "github.com/rs/zerolog"
)

func main() {
    // init a logger to use 
    log := zerolog.New(os.Stdout)

    // creates a *log.Logger with no level prefix
    stdLogger := zstdlog.NewStdLogger(log)
    stdLogger.Println("hello world")
    // results in {"message":"hello world"}

    // creates a *log.Logger with a level prefix
    stdLeveledLogger := zstdlog.NewStdLoggerWithLevel(log, zerolog.InfoLevel)
    stdLeveledLogger.Println("hello world")
    // results in {"level":"info","message":"hello world"}
}
```