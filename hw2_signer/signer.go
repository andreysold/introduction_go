package main

import (
	"fmt"
	"strconv"
	"sync"
)

//type job func(in, out chan interface{})

// сюда писать код

func ExecutePipeline(freeFlowJobs ...job) {
	var wg sync.WaitGroup
	in := make(chan interface{})
	out := make(chan interface{})
	for _, jobs := range freeFlowJobs {
		wg.Add(1)

		out = make(chan interface{})
		go func(jobs job, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			jobs(in, out)
		}(jobs, in, out)
		in = out
	}
	wg.Wait()
}

//type SingleHash func(in, out interface{})

/*func SingleHash(in interface{}, out interface{}) {
	fmt.Println(in, out)
}
*/

func fun(data string, ch chan interface{}) {
	val := DataSignerMd5(data)
	ch <- val
}

var SingleHash = func(in, out chan interface{}) {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := range in {
		wg.Add(1)
		ch := make(chan interface{})
		data := strconv.Itoa(i.(int))
		go func(i interface{}, out chan interface{}, mu *sync.Mutex, wg *sync.WaitGroup) {
			//go func(i interface{}, out chan interface{}, data string) {
			//	//go func(data string, ch chan interface{}) {
			//	//out <- DataSignerMd5(data)
			//	ch <- DataSignerCrc32(data)
			//}(data, ch)
			//}(i, out, data)

		}(i, out, &mu, &wg)
		fmt.Printf("%s SingleHash data %s\n", data, data)
		fmt.Printf("%s SingleHash md5(data) %s\n", data, <-ch)
		fmt.Printf("%s SingleHash crc32(md5(data)) %s\n", data, <-ch)
	}
	wg.Wait()
}

//func main() {
//	in := make(chan interface{})
//	out := make(chan interface{})
//	inputData := []int{0,1,2,3,4}
//	for _, fibNum := range inputData {
//		out <- fibNum
//	}
//
//}
