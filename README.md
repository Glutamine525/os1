# OS Experiment 1: Job Schedule

## Table of Contents

- [Installing](#installing)
- [Monoprogram](#monoprogram)
  - [First Come First Served](#first-come-first-served)
  - [Shortest Job First](#shortest-job-first)
  - [Highest Response-ratio Next](#highest-response-ratio-next)
- [Multiprogram (two-way)](#multiprogram-two-way)
  - [First Come First Served](#first-come--first-served)
  - [Highest Priority First](#highest-priority-first)

## Installing

To start using this project, install Go and run `go get`:

	go get github.com/Glutamine525/os1
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
#### Input Example (entry time, run time)
	8:00 120
	8:50 50
	9:00 10
	9:50 20
#### Output Example (order, entry time, run time, start time, end time, turnaround time, weighted time)
	1 08:00 120 08:00 10:00 120 1.00
	2 08:50  50 10:00 10:50 120 2.40
	3 09:00  10 10:50 11:00 120 12.00
	4 09:50  20 11:00 11:20  90 4.50
		T: 112.500min	W: 4.975

### Shortest Job First
#### Create file `demo.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("SJF1",4)
}
```
#### Input Example (entry time, run time)
	8:00 120
	8:50 50
	9:00 10
	9:50 20
#### Output Example (order, entry time, run time, start time, end time, turnaround time, weighted time)
	1 08:00 120 08:00 10:00 120 1.00
	2 08:50  50 10:30 11:20 150 3.00
	3 09:00  10 10:00 10:10  70 7.00
	4 09:50  20 10:10 10:30  40 2.00
		T: 95.000min	W: 3.250

### Highest Response-ratio Next
#### Create file `demo.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("HRN1",4)
}
```
#### Input Example (entry time, run time)
	8:00 120
	8:50 50
	9:00 10
	9:50 20
#### Output Example (order, entry time, run time, start time, end time, turnaround time, weighted time)
	1 08:00 120 08:00 10:00 120 1.00
	2 08:50  50 10:10 11:00 130 2.60
	3 09:00  10 10:00 10:10  70 7.00
	4 09:50  20 11:00 11:20  90 4.50
		T: 102.500min	W: 3.775

## Multiprogram (two-way)

### First Come  First Served
#### Create file `demo.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("FCFS2", 4)
}
```
#### Input Example (entry time, run time, occupied resource)
	8:00 120 30
	8:50 50 90
	9:00 10 20
	9:50 20 60
#### Output Example (order, entry time, run time, occupied resource, start time, end time, turnaround time, weighted time)
	1 08:00 120  30 08:00 10:00 120 1.00
	2 08:50  50  90 10:30 11:20 150 3.00
	3 09:00  10  20 09:00 10:10  70 7.00
	4 09:50  20  60 10:00 10:30  40 2.00
		T: 95.000min	W: 3.250

### Highest Priority First
#### Create file `demo.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("HPF2", 4)
}
```
#### Input Example (entry time, run time, priority)
	8:00 120 9
	8:50 50 7
	9:00 10 5
	9:50 20 3
#### Output Example (order, entry time, run time, priority, start time, end time, turnaround time, weighted time)
	1 08:00 120   9 08:00 11:20 200 1.67
	2 08:50  50   7 08:50 09:40  50 1.00
	3 09:00  10   5 09:40 09:50  50 5.00
	4 09:50  20   3 09:50 10:10  20 1.00
		T: 80.000min	W: 2.167
