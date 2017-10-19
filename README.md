#OS Experiment One: Job Schedule

#### Donwload and install

	go get github/Glutamine525/os1

#### Create file `JobSchedule.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("FCFS1", 4)
	//jcb.Calc("SJF1",4)
	//jcb.Calc("HRN",4)
}
```

#### Input
```
8:00 120
8:50 50
9:00 10
9:50 20
```

#### Create file `JobSchedule.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("FCFS2", 4)
}
```

#### Input
```
8:00 120 30
8:50 50 90
9:00 10 20
9:50 20 60
```

#### Create file `JobSchedule.go`
```go
package main

import "github.com/Glutamine525/os1"

func main() {
	jcb.Calc("HPF2", 4)
}
```

#### Input
```
8:00 120 9
8:50 50 7
9:00 10 5
9:50 20 3
```
