package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

func verder(wg *sync.WaitGroup, mu *sync.Mutex, j interface{}, out chan interface{}) {
	defer wg.Done()
	tmp := strconv.Itoa(j.(int))

	fmt.Printf("%s SingleHash data %s\n", tmp, tmp)
	mu.Lock()
	d1 := DataSignerMd5(tmp)
	mu.Unlock()
	fmt.Printf("%s SingleHash md5(data) %s\n", tmp, d1)
	dCM := DataSignerCrc32(d1)
	fmt.Printf("%s SingleHash crc32(md5(data)) %s\n", tmp, dCM)
	dC := DataSignerCrc32(tmp)
	fmt.Printf("%s SingleHash crc32(data) %s\n", tmp, dC)
	fmt.Printf("%s SingleHash result %s~%s\n", tmp, dC, dCM)
	out <- fmt.Sprintf("%s~%s", dC, dCM)
}

var SingleHash = func(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for j := range in {
		wg.Add(1)
		go verder(wg, mu, j, out)
	}
	wg.Wait()
}

var MultiHash = func(in, out chan interface{}) {
	hashesDate := make([]string, 6)
	wg := &sync.WaitGroup{}
	for i := range in {
		dt := i.(string)
		wg.Add(1)
		go func(dt string, wg *sync.WaitGroup) {
			defer wg.Done()
			hashWg := &sync.WaitGroup{}
			for th := 0; th < 6; th++ {
				hashWg.Add(1)
				go func(th int, hashWg *sync.WaitGroup) {
					defer hashWg.Done()
					hashesDate[th] = DataSignerCrc32(fmt.Sprintf("%d%s", th, dt))
				}(th, hashWg)
			}
			hashWg.Wait()
			out <- strings.Join(hashesDate, "")
		}(dt, wg)
	}
	wg.Wait()
}

var CombineResults = func(in, out chan interface{}) {
	result := make([]string, 0, 5)
	for d := range in {
		result = append(result, d.(string))
	}
	sort.Strings(result)
	results := strings.Join(result, "_")
	fmt.Printf(results)
	//out <- results
	//results := make([]string, 0, 5)
	//for data := range in {
	//	results = append(results, data.(string))
	//}
	//sort.Strings(results)
	//result := strings.Join(results, "_")
	//out <- result
}
