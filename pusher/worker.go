package pusher

import (
    "fmt"
    "time"
)


func Run() {

    for {
        select {
        case <-time.After(time.Second * 1):
            fmt.Println("yap")
        }
    }
}
