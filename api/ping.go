package api

import (
    "net/http"
    "log"
)

func Ping(w http.ResponseWriter, r *http.Request) {
    log.Println("server ok")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"msg":"hello_world"}`))
}