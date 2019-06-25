package main

import (
    "fmt"
    "time"

    "golang.org/x/crypto/scrypt"
)

func main()  {
    start := time.Now()
    defer func() {
       fmt.Println(time.Since(start))
    }()

    password := []byte{10, 145, 184, 122, 252, 127, 236, 174, 244, 174, 64, 102, 14, 68, 244, 45, 43, 201, 124, 195, 192, 242, 69, 245, 67, 155, 71, 96, 15, 245, 243, 54}
    keySlice, err := scrypt.Key(password, []byte{}, 65536, 1, 32, 32)
    if err != nil {
        panic(err)
    }
    fmt.Println(keySlice)
}

