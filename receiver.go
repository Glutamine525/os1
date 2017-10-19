package jcb

import (
	"math"
	"strconv"
	"time"
)

func (this *Job) set(order, priority, resource int, in time.Time, duration time.Duration) {
	this.order = order
	this.in = in
	this.duration = duration
	this.priority = priority
	this.resource = resource
}

func (this Job) String() string {
	round := math.Trunc(this.wi*1e2+0.5) * 1e-2
	var res string
	res += format(this.order+1) + " "
	res += this.in.Format(timeFormat) + " "
	res += format(this.duration) + " "
	if this.priority != 0 {
		res += format(this.priority) + " "
	}
	if this.resource != 0 {
		res += format(this.resource) + " "
	}
	res += this.begin.Format(timeFormat) + " "
	res += this.out.Format(timeFormat) + " "
	res += format(this.ti) + " "
	res += strconv.FormatFloat(round, 'f', 2, 64)
	return res
}

func (this Jobs) Len() int {
	return len(this)
}

func (this Jobs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Jobs) Less(i, j int) bool {
	return this[j].in.After(this[i].in)
}

func (this prioritySortSlice) Len() int {
	return len(this)
}

func (this prioritySortSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this prioritySortSlice) Less(i, j int) bool {
	return this[i].priority < this[j].priority
}
