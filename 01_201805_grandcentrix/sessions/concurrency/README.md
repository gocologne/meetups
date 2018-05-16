# Einführung in Concurrency


## Nebenläufigkeit in der echten Welt

* [Präsentation](https://docs.google.com/presentation/d/1M7V8AVKFIANK0RYvViI-Hkwp-TNpV-L1kSjrPmfFyB4/edit?usp=sharing)
* Nebenläufigkeit in der realen Welt
* Beispiel mit und ohne Warten


## Channels

* Deadlocks

Beispiele:

* https://play.golang.org/p/GuI1rkWtncc 
* https://play.golang.org/p/asF9lSGyEo1 


## Was sind Go Routinen?

* Go Routinen unterstützen Nebenläufigkeit (concurrency)
* https://de.wikipedia.org/wiki/Nebenl%C3%A4ufigkeit 
* Bei Mehrkernprozessoren werden die Aufgaben auf die Kerne verteilt
* Eine Go Routine ist eine Funktion, welche unabhängig läuft. 
* Keyword: `go`

```go
go meineFunktion()
```

* https://play.golang.org/p/PDnTSFf7JmC 


## Sync Paket - WaitGroup

* main Funktion muss warten, bis alle Go Routinen fertig sind
* Definition einer Variablen WaitGroup
* https://play.golang.org/p/bHDuU4RoNc7 

## Channels

* Exkurs: time.Sleep()

Channels werden verwendet, um Werte zwischen Go Routinen auszutauschen. 

* müssen mit make() erzeugt werden
* Schickt Werte in eine Funktion
* Kommunikation zwischen Go Routinen
* Channel blockiert die Programmausführung
* https://play.golang.org/p/d2O6mq6iQeY 


## Select

* Kontrolliert mehrere Events von Channels
* “Warte auf eine Nachricht aus Channel A oder B”
* https://play.golang.org/p/FnHM72zh4k7 
* https://play.golang.org/p/zzLt80MzYdc 

* Beispiel: 
    * sum Aufgabe: https://play.golang.org/p/HwRoC6adZSU
    * sum Lösung: https://play.golang.org/p/Wmkok6aHXkj


## Buffered Channels - Channel mit Puffer

* Channel ohne Puffer blockiert den Programmfluss
* Channel mit Puffer speichert Werte im Puffer
* Der Puffer kann eine bestimmte Anzahl an Nachrichten speichern 
* Solange im Puffer noch Platz ist, bzw. noch Werte im Puffer sind, blockiert dieser Channel nicht

```go
c := make(chan int, 3) // Buffer = 3
```

* https://play.golang.org/p/snxDh90Jl9H 


## Übung
* Keine Ausgabe Teil 1: https://play.golang.org/p/qCTepL1VOr- 
* Keine Ausgabe Teil 2: https://play.golang.org/p/c4c1EYVIFLJ 
* Join Channels: https://play.golang.org/p/YsPi0ieSjAd 
* WaitGroup mit Channels: https://play.golang.org/p/UIoerxzz3v6 
    * Lösung: https://play.golang.org/p/neWXD2Xmzj- 