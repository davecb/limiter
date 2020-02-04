package main

import (
	"fmt"
	"github.com/davecb/limiter/internal"
	. "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
)

func TestLimiter(t *testing.T) {
	var toWorker, fromWorker chan internal.WorkItem
	var wg sync.WaitGroup
	
	Convey("limiter", t, func() {
		fmt.Printf("staring main\n")
		toWorker = make(chan internal.WorkItem, 1)
		fromWorker = make(chan internal.WorkItem, 1)
		internal.OfferWork(toWorker, internal.WorkList{
			{"C", 300},
		})
		wg.Add(1)
		go internal.DoWork(toWorker, &wg, fromWorker)
		wg.Wait()
	})
}
