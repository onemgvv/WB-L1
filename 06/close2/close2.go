package close2

import "sync"

func Main() {
    var wg sync.WaitGroup
    wg.Add(1)

    ch := make(chan int)
    go func() {
        for {
            foo, ok := <- ch
            if !ok {
                println("done")
                wg.Done()
                return
            }
            println(foo)
        }
    }()
    ch <- 1
    ch <- 2
    close(ch)

    wg.Wait()
}