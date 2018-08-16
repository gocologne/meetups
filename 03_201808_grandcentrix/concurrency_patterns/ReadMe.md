
# Concurrency Patterns

## Simple
### Generator

* Funktion generiert einen Channel und liefert diesen zurück

### Pipeline
* Ein Input und ein Output Channel
* Kann auch als Generator funktionieren

```go
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
```

```go
func sq2(in <-chan int, out chan<- int){
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
}
```
https://play.golang.org/p/V5ihV4qkZVj

### Fan In

* Input: mehrere Channels 
* Output: ein Channel
* Umsetzung mit mehreren Goroutingen
* Umsetzung mit select

### Fan Out
* mehrere Funktionen lesen von einem Channel

### Wait Channel

* Channel blockiert bis eine Nachricht kommt

## Select

### for-select loop

### Timeout mit select

### Quit channel

## More advanced

### Worker

### Semaphore

# Weiterführende Links

* [Visualizing Concurrency in Go](https://divan.github.io/posts/go_concurrency_visualize/)
* [Share Memory By Communicating](https://blog.golang.org/share-memory-by-communicating)
* [Go Concurrency Patterns: Timing out, moving on](https://blog.golang.org/go-concurrency-patterns-timing-out-and)
* [Go Concurrency Patterns: Rob Pike](https://talks.golang.org/2012/concurrency.slide#1)
* [Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns)
* [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
* [Go Concurrency Patterns: Context](https://blog.golang.org/context)
* https://rodaine.com/2018/08/x-files-sync-golang/
* https://godoc.org/golang.org/x/sync