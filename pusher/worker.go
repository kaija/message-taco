package pusher

import (
    "github.com/Sirupsen/logrus"
    "time"
)


func Run() {
    logrus.Infoln("Push worker started.")
    for {
        select {
        case <-time.After(time.Second * 1):
        }
    }
}
