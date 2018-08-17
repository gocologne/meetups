
# Concurrency Patterns

Diese Session baut auf [Einführung in Concurrency](https://github.com/gocologne/meetups/tree/master/01_201805_grandcentrix/sessions/concurrency) auf.

## Wichtig bei Verwendung von Goroutinen
* Wenn man eine Goroutine startet, dann sollte man immer auch an das Ende dieser Goroutine denken!
* Leaking Goroutines
    * Goroutinen welche **ungewollt** unendlich lang laufen
    * Ursache ist meistens blockierender code
    * Problem ist, wenn langsam immer mehr blockierende Groroutinen hinzukommen
    * Garbage Collector kann diese nicht aufräumen
    * Können auch durch Fremdpakete erzeugt werden


## Simple
### Generator
* Die Funktion erzeugt eine Goroutine
* und liefert einen Channel als Return-Wert

```go
func gen() chan string {
	ch := make(chan string)
	go func() {
		fmt.Println("Generator läuft")
		for {
			ch <- "ich warte noch zwei Sekunden"
			time.Sleep(time.Second * 2)
		}
	}()
	return ch
}
```

https://play.golang.org/p/TxGhUNfgE4a

### Pipeline
* Ein Input und ein Output Channel


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
* Kann auch als Generator funktionieren
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

* Übung: [Sieb des Eratosthenes](https://de.wikipedia.org/wiki/Sieb_des_Eratosthenes)
    * Erstelle einen Generator, welcher `int` an einen Channel schickt
    * Erstelle eine Pipeline, welche zu einer Primzahl die Vielfachen rausfiltert
    * Erzeuge für jede gefundene Primzahl einen neuen Filter
    * die erste Zahl, welche durch alle Filter durchkommt ist eine neu gefundene Primzahl

### Fan In

* Input: mehrere Channels 
* Output: ein Channel
* Umsetzung mit mehreren Goroutinen
    * https://play.golang.org/p/oqUck7rxnUy
    * https://play.golang.org/p/78VY_ttsWl2

```go
func fanIn(ch1, ch2, ch3 chan string) chan string {
	out := make(chan string)
	go func() { out <- <-ch1 }()
	go func() { out <- <-ch2 }()
	go func() { out <- <-ch3 }()
	return out
}
```
* Übung: Fan In für bliebig viele Channels mit einer variadischen Funktion (variadic function) 

* Umsetzung mit select

### Fan Out
* mehrere Funktionen lesen von einem Channel

```go
func main() {
	c := make(chan string)
	go myGoroutine(c, "Routine 1")
	go myGoroutine(c, "Routine 2")
	go myGoroutine(c, "Routine 3")
	go myGoroutine(c, "Routine 4")
	c <- "Test1"
	time.Sleep(time.Second)
	c <- "Test2"
	time.Sleep(time.Second)
	c <- "Test3"
	time.Sleep(time.Second)
	c <- "Test4"
	time.Sleep(time.Second)
}

func myGoroutine(ch chan string, name string) {
	fmt.Println("Starte: ", name)
	for {
		msg := <-ch
		fmt.Printf("%s: %s\n", name, msg)
	}
}
```
https://play.golang.org/p/VxQPtHsQEcP

### Wait Channel
* Channel blockiert bis eine Nachricht kommt
* Synchrinisierung von Goroutinen

## Select

### for-select loop
* Standardfall

```go
func myGoroutine(ch1, ch2 chan string) {
	for {
		select {
		case s := <-ch1:
			doThis(s)
		case s := <-ch2:
			doThat(s)
		default:
			// should not block
		}
	}
}
```

### range über Channel
* Verwendung von `range`
* loop endet, wenn channel geschlossen wird

```go
func myGoroutine(ch chan string, name string) {
	fmt.Println("Starte: ", name)
	for msg := range ch {
		fmt.Printf("%s: %s\n", name, msg)
	}
	fmt.Println("Beende: ", name)
}
```
https://play.golang.org/p/hNQkQLPgXsM


### Timeout mit select

### Quit channel

### Leaking Goroutine durch select
* kein Pattern, sondern negativ Beispiel
* oft nicht sofort erkennbar


```go
func myFunc(ctx context.Context) error {
	errc := make(chan error)
	go func() {
		errc <- doSomething()
	}()
	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
```

## Zusammengesetzte Patterns

### Context
* https://golang.org/pkg/context/
* https://blog.golang.org/context

### Worker
* ein oder mehrere Worker werdengestartet
* über einen Channel werden die Aufgaben an die Worker gesendet
* Scheduler verteilt die Aufgaben an die Worker

### Semaphores


* Mit Buffered Channels (https://golang.org/doc/effective_go.html#channels)

```go
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // Create new instance of req for the goroutine.
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}
```


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
* [Golangpatterns](http://www.golangpatterns.info/concurrency)