package main

import (
    "fmt"
    "sync"
    "time"

    "golang.org/x/crypto/scrypt"
)

// takes ~18.9s on macbook pro without -race
// takes ~3m9s on macbook pro with -race
func main()  {
    start := time.Now()
    defer func() {
       fmt.Println(time.Since(start))
    }()

    var wg sync.WaitGroup

    const loadFactor = 100
    wg.Add(loadFactor)
    for i := 0; i < loadFactor; i++ {
        go func(w *sync.WaitGroup) {
            password := []byte{10, 145, 184, 122, 252, 127, 236, 174, 244, 174, 64, 102, 14, 68, 244, 45, 43, 201, 124, 195, 192, 242, 69, 245, 67, 155, 71, 96, 15, 245, 243, 54}
            keySlice, err := scrypt.Key(password, []byte{}, 65536, 1, 32, 32)
            if err != nil {
                panic(err)
            }
            fmt.Println(keySlice)
            w.Done()
        }(&wg)
    }

    wg.Wait()
    fmt.Println("Done")

}

