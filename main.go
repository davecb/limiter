package main

import (
    "fmt"
    "github.com/davecb/limiter/internal"
    "sync"
)

//// Work Items are requests from customers or "wait" directives
//type workItem struct {
//    Name string
//    Count int
//}
//type workList []workItem

// In the main function, we send work to workers and see what happens
func main() {
    var toWorker, fromWorker chan internal.WorkItem
    var wg sync.WaitGroup

    fmt.Printf("staring main\n")
    toWorker = make(chan internal.WorkItem,1)
    fromWorker = make(chan internal.WorkItem,1)
    internal.OfferWork(toWorker, internal.WorkList{
        {"C", 300},
    })
    wg.Add(1)
    go internal.DoWork(toWorker, &wg, fromWorker)
    wg.Wait()
}
//
//// offerWork sends the content of a worklist to a channel
//func offerWork(toWorker chan workItem, work workList) {
//
//    fmt.Printf("staring offerWork\n")
//    for i, w := range work {
//        fmt.Printf("i = %d, i.name = %q, i.count = %d\n", i, w.Name, w.Count)
//        toWorker <- w
//    }
//    close(toWorker)
//}
//
//// doWork reads workItems and spawns goroutines to do them
//func doWork(input chan workItem, wg *sync.WaitGroup, output chan workItem) {
//
//    fmt.Printf("staring doWork\n")
//    for elem := range input {
//        fmt.Printf("j.name  = %q, j.count = %d\n", elem.Name, elem.Count)
//    }
//    wg.Done()
//}
