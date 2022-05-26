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

	//wg := &sync.WaitGroup{}
	//var df string
	//var result []string
	result := make([]string, 6)
	for i := range in {
		//wg.Add(1)
		//go func(i interface{}, wg *sync.WaitGroup, df string) {
		//	defer wg.Done()
		for th := 0; th <= 5; th++ {
			t := DataSignerCrc32(fmt.Sprintf("%d%s", th, i))
			//df += t
			result = append(result, t)
			fmt.Printf("%s MultiHash: crc32(th+step1)) %d %s\n", i, th, t)
			t = ""
		}
		//fmt.Println(df)
		//result = append(result, df)
		//df = ""

		//}(i, wg, df)
	}
	//wg.Wait()
	//l := strings.Join(df, "")
	//fmt.Println("|", l, "|")
	//fmt.Println(strings.Join(result, " "))
	out <- strings.Join(result, "")
}

var CombineResults = func(in, out chan interface{}) {
	//it := <-in
	var df []string
	for data := range in {
		df = append(df, fmt.Sprintf("%v", data))
	}
	sort.Strings(df)
	fmt.Printf("%v\n", df)
	//for i := range df {
	//	fmt.Println(df[i])
	//}
	out <- strings.Join(df, "_")
}
