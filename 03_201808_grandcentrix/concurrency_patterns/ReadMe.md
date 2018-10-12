



<!-- TOC -->

- [1. Concurrency Patterns](#1-concurrency-patterns)
	- [1.1. Wichtig bei Verwendung von Goroutinen](#11-wichtig-bei-verwendung-von-goroutinen)
	- [1.2. Simple](#12-simple)
		- [1.2.1. Generator](#121-generator)
		- [1.2.2. Pipeline](#122-pipeline)
		- [1.2.3. Fan In](#123-fan-in)
		- [1.2.4. Fan Out](#124-fan-out)
		- [1.2.5. Wait Channel](#125-wait-channel)
		- [1.2.6. Channel of Channel](#126-channel-of-channel)
	- [1.3. Select](#13-select)
		- [1.3.1. for-select loop](#131-for-select-loop)
		- [1.3.2. range über Channel](#132-range-%C3%BCber-channel)
		- [1.3.3. Timeout mit select](#133-timeout-mit-select)
		- [1.3.4. Quit channel](#134-quit-channel)
		- [1.3.5. Leaking Goroutine durch select](#135-leaking-goroutine-durch-select)
	- [1.4. Zusammengesetzte Patterns](#14-zusammengesetzte-patterns)
		- [1.4.1. Context](#141-context)
		- [1.4.2. Worker](#142-worker)
		- [1.4.3. Semaphores](#143-semaphores)
		- [1.4.4. Quit signal](#144-quit-signal)
		- [1.4.5. State Machnie](#145-state-machnie)
- [2. Weiterführende Links](#2-weiterf%C3%BChrende-links)

<!-- /TOC -->

# 1. Concurrency Patterns

Diese Session baut auf [Einführung in Concurrency](https://github.com/gocologne/meetups/tree/master/01_201805_grandcentrix/sessions/concurrency) auf.

## 1.1. Wichtig bei Verwendung von Goroutinen
* Wenn man eine Goroutine startet, dann sollte man immer auch an das Ende dieser Goroutine denken!
* Leaking Goroutines
    * Goroutinen welche **ungewollt** unendlich lang laufen
    * Ursache ist meistens blockierender code
    * Problem ist, wenn langsam immer mehr blockierende Groroutinen hinzukommen
    * Garbage Collector kann diese nicht aufräumen
    * Können auch durch Fremdpakete erzeugt werden


## 1.2. Simple
### 1.2.1. Generator
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

### 1.2.2. Pipeline
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

* Ein Input und ein Output Channel
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

### 1.2.3. Fan In

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

### 1.2.4. Fan Out
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

### 1.2.5. Wait Channel
* Channel blockiert bis eine Nachricht kommt
* Synchrinisierung von Goroutinen
* Wird im time Paket verwendet
	* https://golang.org/pkg/time/#After
	* https://golang.org/pkg/time/#Tick

### 1.2.6. Channel of Channel
* Problemstellung
	* Goroutine 1 möchte in Goroutine 2 etwas auslösen 
	* Goroutine 1 benötigt jedoch Rückmeldung von Goroutine 2
* Lösung
	* Goroutine 2 behandelt den Typ `chan chan Type`
	* Goroutine 1 schickt somit einen `chan Type` über `chan chan Type`

```go
var wg sync.WaitGroup

func main() {
	count := make(chan int)
	reset := make(chan chan int)
	counter := make(chan int)
	wg.Add(1)
	go routine1(count, reset)
	count <- 1
	count <- 2
	reset <- counter
	fmt.Println(<-counter)
	close(count)
	wg.Wait()
}

func routine1(count chan int, reset chan chan int) {
	defer wg.Done()
	counter := 0
	for {
		select {
		case i, ok := <-count:
			if !ok {
				return
			}
			counter += i
		case c := <-reset:
			c <- counter
			counter = 0
		}
	}
}
```
https://play.golang.org/p/ibJAWo86H50

## 1.3. Select

### 1.3.1. for-select loop
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

### 1.3.2. range über Channel
* Verwendung von `range`
* loop endet, wenn channel geschlossen wird
* wenn möglich immer verwenden

```go
func myGoroutine(ch chan string, name string) {
	fmt.Println("Starte: ", name)
	// setup()
	for msg := range ch {
		fmt.Printf("%s: %s\n", name, msg)
	}
	// cleanup()
	fmt.Println("Beende: ", name)
}
```
https://play.golang.org/p/hNQkQLPgXsM


### 1.3.3. Timeout mit select

```go
timeout := time.Second * 3
result := make(chan int, 1)
go func() {
	// r := doSomething()
	result <- r
}()

select {
case r := <-result:
	//useResult(r)
case <-time.After(timeout):
	fmt.Println("timeout")
}
```

### 1.3.4. Quit channel

```go
select {
case r := <-result:
	//useResult(r)
case <-quitc:
	fmt.Println("timeout")
}
```

### 1.3.5. Leaking Goroutine durch select
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

## 1.4. Zusammengesetzte Patterns

### 1.4.1. Context
* https://golang.org/pkg/context/
* https://blog.golang.org/context

### 1.4.2. Worker
* ein oder mehrere Worker werdengestartet
* über einen Channel werden die Aufgaben an die Worker gesendet
* Scheduler verteilt die Aufgaben an die Worker

### 1.4.3. Semaphores


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

```go
	limit := 5
	sem := make(chan bool, limit)
	hugeSlice := make([]bool, 10)
	for i, task := range hugeSlice {
		sem <- true
		go func(task bool, nr int) {
			fmt.Println("Working: ", nr)
			time.Sleep(time.Second * time.Duration(nr))
			fmt.Println("Ready: ", nr)
			<-sem
		}(task, i)
	}
	for n := limit; n > 0; n-- {
		sem <- true
		fmt.Println(n, "to go")
	}
```
https://play.golang.org/p/2jMeNZ7Arg6 	


### 1.4.4. Quit signal

* starte eine Goroutine, welche das Ctrl+C verarbeitet
* Quit channel wird geschlossen
* vor `os.Exit()` sollte gewartet werden bis alle Aufräumarbeiten abgeschlossen sind
* [Notify benötigt einen buffered channel](https://golang.org/pkg/os/signal/#Notify) 

```go
c := make(chan os.Signal, 1)
quitc := make(chan struct{})

signal.Notify(c, os.Interrupt)
go func() {
	for s := range c {
		fmt.Println("Got signal:", s)
		fmt.Println("closing quitc")
		close(quitc)
		cleanup()
		os.Exit(0)
	}
}()
```

### 1.4.5. State Machnie
* Peter Bourgon: https://www.youtube.com/watch?v=LHe1Cb_Ud_M
* Verwendung, wenn Zugriffe auf den Typ über mehrere Goroutinen erfolgen
* Alle lesenden und schreibenen Zugriffe können über die Methode `loop()` erfolgen
* Vermeindung von Dataraces, da ja nur `loop()` lesen und schreiben darf

```go
type stateMachine struct {
	state   string
	actionc chan func()
	quitc   chan struct{}
}

func (sm *stateMachine) loop() {
	for {
		select {
		case f := <-sm.actionc:
			f()
		case <-sm.quitc:
			return
		}
	}
}

func (sm *stateMachine) foo() int {
	c := make(chan int)
	// Definition der Funktion hier jedoch wird diese
	// im loop ausgeführt
	sm.actionc <- func() {
		if sm.state == "A" {
			sm.state = "B"
		}
		c <- 123
	}
	return <-c
}

// Im Constructor kann hierzu die Goroutine gestartet werden

func New() *stateMachine {
	sm := &stateMachine{
		state:   "initial",
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go sm.loop()
	return sm
}
```

# 2. Weiterführende Links

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