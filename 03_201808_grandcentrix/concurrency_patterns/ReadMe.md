
# Concurrency Patterns

## Simple
### Generator

* Funktion generiert einen Channel und liefert diesen zurück


### Fan In

* Input: mehrere Channels 
* Output: ein Channel
* Umsetzung mit mehreren Goroutingen
* Umsetzung mit select

### Wait Channel

* Channel blockiert bis eine Nachricht kommt

### Timeout mit select

### Quit channel

# Weiterführende Links

* [Share Memory By Communicating](https://blog.golang.org/share-memory-by-communicating)
* [Go Concurrency Patterns: Timing out, moving on](https://blog.golang.org/go-concurrency-patterns-timing-out-and)
* [Go Concurrency Patterns: Rob Pike](https://talks.golang.org/2012/concurrency.slide#1)
* [Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns)
* [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
* [Go Concurrency Patterns: Context](https://blog.golang.org/context)
* https://rodaine.com/2018/08/x-files-sync-golang/
* https://godoc.org/golang.org/x/sync