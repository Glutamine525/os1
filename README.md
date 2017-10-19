# OS Experiment One: Job Schedule

## Table of Contents

- [Installing](#installing)
- [Monoprogram](#monoprogram)
  - [First Come First Served](#first-come-first-served)
  - [Shortest Job First](#shortest-job-first)
  - [Highest Response-ratio Next](#highest-response-ratio-next)
- [Multiprogram (two-way)](#multiprogram-two-way)
  - [First Come First Served](#first-come-first-served)
  - [Highest Priority Frequency](#highest-priority-frequency)

## Installing

To start using this project, install Go and run `go get`:
```
go get github.com/Glutamine525/os1
```
This will retrieve the library and install the project command line utility into your `$GOPATH`.

## Monoprogram

### First Come First Served
#### Create file `demo.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("FCFS1", 4)
}
```
#### Input(entry time, run time)
```
8:00 120
8:50 50
9:00 10
9:50 20
```

### Shortest Job First
#### Create file `demo.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("SJF1",4)
}
```
#### Input(entry time, run time)
```
8:00 120
8:50 50
9:00 10
9:50 20
```

### Highest Response-ratio Next
#### Create file `demo.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("HRN",4)
}
```
#### Input(entry time, run time)
```
8:00 120
8:50 50
9:00 10
9:50 20
```

## Multiprogram (two-way)

### First Come First Served
#### Create file `demo.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("FCFS2", 4)
}
```
#### Input(entry time, run time, occupied resource)
```
8:00 120 30
8:50 50 90
9:00 10 20
9:50 20 60
```

### Highest Priority Frequency
#### Create file `demo.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("HPF2", 4)
}
```
#### Input(entry time, run time, priority)
```
8:00 120 9
8:50 50 7
9:00 10 5
9:50 20 3
```
