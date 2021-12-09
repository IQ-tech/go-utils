## Installation

```terminal
go get github.com/IQ-tech/go-utils
```

## Executing code before process exits

`AtInterruption` receives a function that will be called once before the process exits.  
`AtInterruption` can be called many times with different functions and each function will be
called once.

```go
utils.AtInterruption(func() {
	log.Printf("Closing DB Connection.")
	db.Close()

	log.Printf("Shutting down New Relic application.")
	nrl.Shutdown(5 * time.Second)
})
```

## Logging and terminating the process on error

`EndAsErr` will log the error message and terminate the process if `err` is not `nil`

```go
func main() {
	di := setup.Start()
	err := server.Start(di)
	utils.EndAsErr(err, "Could not start the server.", di.Log.InfoWriter(), di.Log.ErrorWriter())
}

```
