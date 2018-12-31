# zerolog-gocb
Simple adapter so that Zerolog can be used with GO Couchbase

## Basic Usage

```go
package main

import(
	"os"
	
	"github.com/couchbase/gocb"
	"github.com/dcarbone/zerolog-gocb"
	"github.com/rs/zerolog"
)

func main() {
	// init a logger to use 
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Fields(map[string]interface{}{"source": "couchbase"}).
		Logger()
	
	// create the adapter.
	adapter := zerologgocb.NewDefault(log)
	
	// set logger
	gocb.SetLogger(adapter)
	
	// presumably do something
}
```

## LevelMap

The [LevelMap](compat.go#L10) maps a `gocb.LogLevel` to a `zerolog.Level`.  If you wish a given
Couchbase log level to be ignored, you may either simply not set it or set it to `zerolog.Disabled`.


There is a [DefaultLevelMap](compat.go#L22) that will be used if `nil` is passed as first arg to `New()` or
if the [Adapter](compat.go#L13) is constructed using [NewDefault()](compat.go#L46)

And that's pretty much it.