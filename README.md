# Zlog 

## Overview

Zlog is a wrapper to [Zap](https://github.com/uber-go/zap) the structured Go logger
developed by Uber.

It provides Verbosity levels to Debug logging to allow for finer control over
the level of Debug logging you want to see.

A standard lib log Writer is also implemented to allow passing of zlog as the 
HTTP error log.


## Usage

A package level singleton instance of zlog can be used or you can initialize
your own or multiple instances.

### Singleton 

The Singleton logger by default is set to disabled so usage of it must first
initialize the operating mode.

```go
// sets singletone logger to Production mode with human readable output with a 
// verbosity level of zero.
zlog.SetMode(zlog.ModeProduction, zlog.EncodingConsole, 0)

zlog.Info("Hello world")

// Debug message with verbosity set to level 1
zlog.Debug(1, "Sensor activated",
	zlog.Uint32("voltage", 5),
	)
```

### Instance

Create your own or multiple instances of a zlog.

```go
// Create instance of zlog setting verbosity to level 3 with JSON structed
// output
logger := zlog.New(zlog.ModeDevel, zlog.EncodingJson, 3)

logger.Debug(5, "Sensor activated",
    zlog.Uint32("voltage", 5),
    )
```




### Verbosity Levels

To set verbosity levels throughout your code place `zlog.Debug()` calls and set
the verbosity level for each with a value from 0 to 9, where level 0 is the quietest
and level 9 provides the most detail, eg:

```go
zlog.Debug(8, "Wrote DB query",
	zlog.String("query", "SELECT * FROM TABLE"),
	)


zlog.Debug(0, "Sensor Activated",
	zlog.Time("time", time.Now()),
	)
```

Then implement a command line or config switch so when initialising the zlog instance
you pass the desired Verbosity level.

```go
logger := zlog.New(zlog.ModeDevel, zlog.EncodingJson, 3)
```

In the above case Verbosity is set to level 3 so only the `Sensor Activated` message would be shown.

To see both `Wrote DB query` and `Sensor Activated` messages set verbosity to level 8 or 9
when initializing the logger, eg:

```go
logger := zlog.New(zlog.ModeDevel, zlog.EncodingJson, 9)
```




### HTTP Logger

Passing zlog as the Error log to the standard lib HTTP library for logging.

```go
// create instance of zlog
logger := zlog.New(zlog.ModeDevel, zlog.EncodingJson, 3)

// code is abbreviated to show passing to ErrorLog field only 
server = &http.Server{ 	
	ErrorLog:    log.New(&zlog.Writer{Logger: logger, Logtype: zlog.ErrorLog}, "HTTP", 0),
}

server.ListenAndServe()	
```

