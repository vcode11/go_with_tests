package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Racer(urls  ...string) string {
	loadTimes := make([]int64, len(urls))
	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func (i int) {
		 	getLoadTime(url, loadTimes, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	idx, minLoadTime := 0, int64((1 << 63) - 1)
	for i, time := range loadTimes {
		if time <= minLoadTime {
			fmt.Print(minLoadTime)
			minLoadTime = time
			idx = i
		}
	}
	return urls[idx]
}

func getLoadTime(url string, loadTimes []int64, idx int) {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	loadTimes[idx] = int64(duration)
}