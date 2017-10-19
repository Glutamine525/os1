package jcb

import "time"

type Job struct {
	order    int           //作业序号，从0开始记录
	in       time.Time     //进入时间
	duration time.Duration //运行时间，单位min
	status   int           //状态，0：Wait；1：Run；2：Finish
	priority int           //优先级
	resource int           //所需资源，单位KB
	begin    time.Time     //开始时间
	out      time.Time     //结束时间
	ti       time.Duration //周转时间，单位min
	wi       float64       //带权周转时间
}

type prioritySort struct {
	index    int
	priority int
}

type Jobs []Job

type prioritySortSlice []prioritySort
