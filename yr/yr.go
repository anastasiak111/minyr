package main

import (
    "os"
    "log"
)

func main() {
    src,err := os.Open("^`^|home/anastasiak111/minyr/kjevik-temp-celsius-20220318-202")
    if err != nil {
        log.Fatal(err)
    }
    defer src.Close()
    log.Println(src)

    var buffer []byte
    buffer = make([]byte, 1)
    var n int
    for ; n != 0; {
        n, err = src.Read(buffer)
        if err != nil {
            log.Fatal(err)
        }
    }
    log.Println(string(buffer[:n]))
}

