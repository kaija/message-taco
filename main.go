package main

import (
    "github.com/Sirupsen/logrus"
    "github.com/spf13/viper"
    "net/http"
    "context"
    "os"
    "os/signal"
    "syscall"

    "github.com/kaija/message-taco/application"
    "github.com/kaija/message-taco/pusher"
)

func newConfig() (*viper.Viper, error) {
    c := viper.New()
    c.SetDefault("cookie_secret", "b3KnGMmYxp8Ly4Wy")
    c.SetDefault("http_addr", ":8888")
    c.SetDefault("http_cert_file", "")
    c.SetDefault("http_key_file", "")
    c.SetDefault("http_drain_interval", "1s")

    c.AutomaticEnv()

    return c, nil
}

func main() {
    config, err := newConfig()
    if err != nil {
        logrus.Fatal(err)
    }
    // Launch worker for message consume
    go pusher.Run()

    app, err := application.New(config)
    if err != nil {
        logrus.Fatal(err)
    }

    middle, err := app.MiddlewareStruct()
    if err != nil {
        logrus.Fatal(err)
    }

    serverAddress := config.Get("http_addr").(string)

    certFile := config.Get("http_cert_file").(string)
    keyFile := config.Get("http_key_file").(string)

    if err != nil {
        logrus.Fatal(err)
    }

    sigstop := make(chan os.Signal, 1)
    shutdown := make(chan int, 1)
    signal.Notify(sigstop, syscall.SIGINT)
    signal.Notify(sigstop, syscall.SIGTERM)

    server := &http.Server{Addr: serverAddress, Handler: middle}

    go func() {
        logrus.Infoln("Running HTTP server on " + serverAddress)
        if certFile != "" && keyFile != "" {
            err = server.ListenAndServeTLS(certFile, keyFile)
        } else {
            err = server.ListenAndServe()
        }
    }()

    go func() {
        sig := <-sigstop
        logrus.Infoln("catch signal: ", sig)
        shutdown<-1
    }()

    <-shutdown
    logrus.Infoln("Shutting down server...")
    server.Shutdown(context.Background())
/*
    srv := &graceful.Server{
        Timeout: drainInterval,
        Server:  &http.Server{Addr: serverAddress, Handler: middle},
    }
*/


    if err != nil {
        logrus.Fatal(err)
    }
}
