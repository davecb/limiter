package internal

import (
    "fmt"
    "sync"
    "time"
)

// Work Items are requests from customers or "wait" directives
type WorkItem struct {
    Name string
    Count int
}
type WorkList []WorkItem

// CustomerQuota contains the values we're limiting to
type customerQuota struct {
    Limit int
    Currently int
}

// OfferWork sends the content of a worklist to a channel
func OfferWork(toWorker chan WorkItem, work WorkList) {

    fmt.Printf("starting OfferWork\n")
    for i, w := range work {
        fmt.Printf("i = %d, i.name = %q, i.count = %d\n", i, w.Name, w.Count)
        toWorker <- w
    }
    fmt.Printf("closing toWorker\n")
    close(toWorker)
}

// DoWork reads workItems and spawns goroutines to do them
func DoWork(input chan WorkItem, wg *sync.WaitGroup, output chan WorkItem) {
    var total int
    var customers = make(map[string]customerQuota)


    fmt.Printf("staring DoWork\n")
    for {
        select {
        // "Done" indications, coming back from workers
        case workDone, ok := <-output:
            if !ok {
                // The channel is closed
                panic("output is closed")
            }
           total -= workDone.Count
           fmt.Printf("done.name = %q, done.count = %d, total = %d \n",
               workDone.Name, workDone.Count, total)
           c := customers[workDone.Name]
           c.Currently -= workDone.Count
           customers[workDone.Name] = c

        // "New" indications, coming from the simulated users
        case newWork, ok := <-input:
            if !ok {
                //fmt.Printf("input is closed, ignore it\n")
                time.Sleep(1 * time.Second) // just waste some time
                continue
            }
            fmt.Printf("newWork.name = %q, newWork.count = %d\n", newWork.Name, newWork.Count)
            c := customers[newWork.Name]
            if (total + newWork.Count) > 450 && (c.Currently + newWork.Count) > c.Limit {
                // Reject this unit of work, it's above both limits
                fmt.Printf("count is too high, looping back\n")
                continue
            }
            total += newWork.Count
            c = customers[newWork.Name]
            c.Currently -= newWork.Count
            customers[newWork.Name] = c

            go func(elem WorkItem, output chan WorkItem) {
                fmt.Printf("work.name  = %q, work.count = %d\n", elem.Name, elem.Count)
                bound := elem.Count
                elem.Count = 1
                for i := 0; i < bound; i++ {
                    time.Sleep(1 * time.Second)
                    output <- elem
                }
            }(newWork, output)

        default:
            fmt.Printf("sleeping\n")
            time.Sleep(1 * time.Second) // just waste some time
        }
    }
    wg.Done()
}
