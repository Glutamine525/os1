package jcb

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const timeFormat string = "15:04"

func resetStatus(jobs Jobs) {
	for i := range jobs {
		jobs[i].status = 0
	}
}

func format(v interface{}) string {
	var result int
	switch value := v.(type) {
	case int:
		result = value
	case time.Duration:
		result = int(value / time.Minute)
	}
	switch {
	case result < 10:
		return "  " + strconv.Itoa(result)
	case result < 100:
		return " " + strconv.Itoa(result)
	default:
		return strconv.Itoa(result)
	}
}

func Calc(remark string, num int) {
	var T, W float64
	fmt.Printf("Input:\n")
	jobs := get(remark, num)
	switch remark {
	case "FCFS1":
		T, W = FCFS_Single(jobs)
	case "SJF1":
		T, W = SJF_Single(jobs)
	case "HRN1":
		T, W = HRN_Single(jobs)
	case "FCFS2":
		T, W = FCFS_Double(jobs)
	case "HPF2":
		T, W = HPF_Double(jobs)
	}
	fmt.Printf("\nOutput:\n")
	for _, v := range jobs {
		fmt.Println(v)
	}
	fmt.Printf("     T: %.3fmin       W: %.3f\n", T, W)
}

func get(remark string, num int) Jobs {
	switch remark {
	case "FCFS1":
		break
	case "SJF1":
		break
	case "HRN1":
		break
	case "FCFS2":
		break
	case "HPF2":
		break
	default:
		fmt.Println("Unknown method.")
		return nil
	}
	jobs := make([]Job, num)
	var in string
	var duration, priority, resource int
	for i := 0; i < num; i++ {
		fmt.Scan(&in, &duration)
		switch remark {
		case "FCFS2":
			fmt.Scan(&resource)
		case "HPF2":
			fmt.Scan(&priority)
		}
		tmp := strings.Split(in, ":")
		hour, _ := strconv.Atoi(tmp[0])
		min, _ := strconv.Atoi(tmp[1])
		inTime := time.Date(2000, time.January, 1, hour, min, 0, 0, time.UTC)
		jobs[i].set(i, priority, resource, inTime, time.Duration(duration)*time.Minute)
	}
	return jobs
}

func FCFS_Single(jobs Jobs) (T float64, W float64) {
	num := len(jobs)
	if num == 0 {
		return -1, -1
	}
	sort.Sort(jobs)
	var now time.Time
	now = jobs[0].in
	for i, v := range jobs {
		if now.Before(v.in) {
			now = v.in
		}
		jobs[i].begin = now
		jobs[i].out = now.Add(jobs[i].duration)
		jobs[i].ti = jobs[i].out.Sub(jobs[i].in)
		jobs[i].wi = float64(jobs[i].ti) / float64(jobs[i].duration)
		now = jobs[i].out
		T += float64(jobs[i].ti / time.Minute)
		W += jobs[i].wi
	}
	return T / float64(num), W / float64(num)
}

func SJF_Single(jobs Jobs) (T float64, W float64) {
	num := len(jobs)
	if num == 0 {
		return -1, -1
	}
	sort.Sort(jobs)
	resetStatus(jobs)
	var count int
	var now time.Time
	now = jobs[0].in
	for count < num {
		min := int64(1<<63 - 1)
		index := -1
		for i, v := range jobs {
			if now.Before(v.in) {
				break
			} else if v.status == 0 {
				if int64(v.duration) < min {
					min = int64(v.duration)
					index = i
				}
			}
		}
		jobs[index].begin = now
		jobs[index].out = now.Add(jobs[index].duration)
		jobs[index].ti = jobs[index].out.Sub(jobs[index].in)
		jobs[index].wi = float64(jobs[index].ti) / float64(jobs[index].duration)
		now = jobs[index].out
		T += float64(jobs[index].ti / time.Minute)
		W += jobs[index].wi
		jobs[index].status = 2
		count++
	}
	return T / float64(num), W / float64(num)
}

func HRN_Single(jobs Jobs) (T float64, W float64) {
	num := len(jobs)
	if num == 0 {
		return -1, -1
	}
	sort.Sort(jobs)
	resetStatus(jobs)
	var count int
	var now time.Time
	now = jobs[0].in
	for count < num {
		var ratio float64
		index := -1
		for i, v := range jobs {
			if now.Before(v.in) {
				break
			} else if v.status == 0 {
				if float64(now.Sub(v.in)/v.duration+1) > ratio {
					ratio = float64(now.Sub(v.in)/v.duration + 1)
					index = i
				}
			}
		}
		jobs[index].begin = now
		jobs[index].out = now.Add(jobs[index].duration)
		jobs[index].ti = jobs[index].out.Sub(jobs[index].in)
		jobs[index].wi = float64(jobs[index].ti) / float64(jobs[index].duration)
		now = jobs[index].out
		T += float64(jobs[index].ti / time.Minute)
		W += jobs[index].wi
		jobs[index].status = 2
		count++
	}
	return T / float64(num), W / float64(num)
}

func FCFS_Double(jobs Jobs) (T float64, W float64) {
	num := len(jobs)
	if num == 0 {
		return -1, -1
	}
	sort.Sort(jobs)
	resetStatus(jobs)
	var now time.Time
	now = jobs[0].in
	index := []int{}
	done := make([]time.Duration, num)
	count, pipe, resource := 0, 2, 100
	for count < num {
		for i, v := range jobs {
			if now.Before(v.in) || pipe == 0 {
				break
			}
			if resource >= v.resource && v.status == 0 {
				pipe--
				resource -= v.resource
				index = append(index, i)
				jobs[i].begin = now
				jobs[i].status = 1
			}
		}
		sort.Ints(index)
		done[index[0]] += time.Minute
		now = now.Add(time.Minute)
		if done[index[0]] == jobs[index[0]].duration {
			pipe++
			resource += jobs[index[0]].resource
			jobs[index[0]].out = now
			jobs[index[0]].status = 2
			jobs[index[0]].ti = jobs[index[0]].out.Sub(jobs[index[0]].in)
			jobs[index[0]].wi = float64(jobs[index[0]].ti) / float64(jobs[index[0]].duration)
			T += float64(jobs[index[0]].ti / time.Minute)
			W += jobs[index[0]].wi
			index = index[1:]
			count++
		}
	}
	return T / float64(num), W / float64(num)
}

func HPF_Double(jobs Jobs) (T float64, W float64) {
	num := len(jobs)
	if num == 0 {
		return -1, -1
	}
	sort.Sort(jobs)
	resetStatus(jobs)
	var now time.Time
	now = jobs[0].in
	record := prioritySortSlice{}
	done := make([]time.Duration, num)
	count, pipe := 0, 2
	for count < num {
		for i, v := range jobs {
			if now.Before(v.in) || pipe == 0 {
				break
			}
			if v.status == 0 {
				pipe--
				record = append(record, prioritySort{i, v.priority})
				jobs[i].begin = now
				jobs[i].status = 1
			}
		}
		sort.Sort(record)
		done[record[0].index] += time.Minute
		now = now.Add(time.Minute)
		if done[record[0].index] == jobs[record[0].index].duration {
			pipe++
			jobs[record[0].index].out = now
			jobs[record[0].index].status = 2
			jobs[record[0].index].ti = jobs[record[0].index].out.Sub(jobs[record[0].index].in)
			jobs[record[0].index].wi = float64(jobs[record[0].index].ti) / float64(jobs[record[0].index].duration)
			T += float64(jobs[record[0].index].ti / time.Minute)
			W += jobs[record[0].index].wi
			record = record[1:]
			count++
		}
	}
	return T / float64(num), W / float64(num)
}
