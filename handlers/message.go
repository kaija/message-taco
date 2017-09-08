package handlers

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "strings"
    "github.com/Jeffail/gabs"
)


func PushMessage(w http.ResponseWriter, r *http.Request) {

    var status int = http.StatusOK

    body, _ := ioutil.ReadAll(r.Body)

    jsonParsed, _ := gabs.ParseJSON(body)

    provider, ok := jsonParsed.Path("provider").Data().(string)

    if ok {
        switch strings.ToLower(provider) {
        case "apns":
            fmt.Println("APNS")
        case "gcm":
            fmt.Println("APNS")
        }
    } else {
        status = http.StatusBadRequest
    }

    w.WriteHeader(status)
    fmt.Fprintf(w, "OK");
}
