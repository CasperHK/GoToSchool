package main

import (
    "net/http"
    "sys/webapi"
    "sys/db"
)

func main() {
    // Serving a HTTP web server.
    http.Handle("/", http.FileServer(http.Dir("./")))

    // Registrate required web api.
    webapi.Registrate()

    // Connect database
    db.Connect()
    defer db.Disconnect()

    http.ListenAndServe(":3000", nil)
}
