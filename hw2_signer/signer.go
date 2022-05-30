package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

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

func cRc32Data(data string) chan string {
	ch1 := make(chan string)
	go func() {
		tmp := DataSignerCrc32(data)
		ch1 <- tmp
	}()
	return ch1
}

func verder(wg *sync.WaitGroup, j interface{}, out chan interface{}, md5 string) {
	defer wg.Done()
	data := strconv.Itoa(j.(int))
	ch1 := cRc32Data(data)
	ch2 := cRc32Data(md5)
	out <- <-ch1 + "~" + <-ch2
}

var SingleHash = func(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for j := range in {
		wg.Add(1)
		md5 := DataSignerMd5(strconv.Itoa(j.(int)))
		go verder(wg, j, out, md5)
	}
	wg.Wait()
}

var MultiHash = func(in, out chan interface{}) {
	var (
		wg    = &sync.WaitGroup{}
		mu    = &sync.Mutex{}
		waitG = &sync.WaitGroup{}
	)
	for v := range in {
		waitG.Add(1)
		hashesDate := make([]string, 6)
		go func(v interface{}) {
			defer waitG.Done()
			for th := 0; th <= 5; th++ {
				wg.Add(1)
				go func(th int) {
					defer wg.Done()
					info := DataSignerCrc32(fmt.Sprintf("%d%s", th, v))
					mu.Lock()
					hashesDate[th] = info
					mu.Unlock()
				}(th)
			}
			wg.Wait()
			out <- strings.Join(hashesDate, "")
		}(v)
	}
	waitG.Wait()
}

var CombineResults = func(in, out chan interface{}) {
	result := make([]string, 0, 5)
	for d := range in {
		result = append(result, d.(string))
	}
	sort.Strings(result)
	results := strings.Join(result, "_")
	out <- results
}
