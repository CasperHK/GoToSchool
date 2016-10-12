package sys

import (
    "net/http"
    "webapi"
    "db"
)

func main() {
    // Serving a HTTP web server.
    http.Handle("/", http.FileServer(http.Dir("./")))

    // Registrate required web api.
    webapi.Registrate()

    // Init database
    db.Init()

    http.ListenAndServe(":3000", nil)
}
