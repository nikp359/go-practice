package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

// ExecutePipeline exec
func ExecutePipeline(fun ...job) {
	inChannel := make(chan interface{}, 100)
	outChannel := make(chan interface{}, 100)
	wg := &sync.WaitGroup{}
	for _, f := range fun {
		outChannel = make(chan interface{}, 100)
		wg.Add(1)
		go func(fRun job, in chan interface{}, out chan interface{}) {
			fRun(in, out)
			close(out)
			wg.Done()
		}(f, inChannel, outChannel)
		inChannel = outChannel
	}
	wg.Wait()
}

func crcСount(data string, out chan<- string) {
	out <- DataSignerCrc32(data)
}

// SingleHash Execute
func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for input := range in {
		data := strconv.Itoa(input.(int))
		md5data := DataSignerMd5(data)
		wg.Add(1)
		go func(data string, md5data string, out chan interface{}) {
			crcChan1 := make(chan string)
			crcChan2 := make(chan string)
			go crcСount(data, crcChan1)
			go crcСount(md5data, crcChan2)
			out1 := <-crcChan1
			out2 := <-crcChan2
			out <- out1 + "~" + out2
			wg.Done()
		}(data, md5data, out)
	}
	wg.Wait()
}

// MultiHash Execut
func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for input := range in {
		wg.Add(1)
		go func(data string, out chan interface{}) {
			result := make([]string, 6)
			mu := &sync.Mutex{}
			wgTh := &sync.WaitGroup{}
			for th := 0; th <= 5; th++ {
				wgTh.Add(1)
				go func(rSlice *[]string, th int, mu *sync.Mutex) {
					dataSig := DataSignerCrc32(strconv.Itoa(th) + data)
					mu.Lock()
					r := *rSlice
					r[th] = dataSig
					mu.Unlock()
					wgTh.Done()
				}(&result, th, mu)
			}
			wgTh.Wait()
			out <- strings.Join(result, "")
			wg.Done()
		}(input.(string), out)
	}
	wg.Wait()
}

// CombineResults Execute
func CombineResults(in, out chan interface{}) {
	var combineResult []string
	for input := range in {
		combineResult = append(combineResult, input.(string))
	}
	sort.Strings(combineResult)
	out <- strings.Join(combineResult, "_")
}
