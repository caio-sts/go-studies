package tests

import (
    "net/http"
    "testing"
    "time"
    "log"
    "basics/api"
)

func startTestServer() *http.Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/api", api.Ping)
    server := &http.Server{
        Addr:    ":8081",
        Handler: mux,
    }
    go func() {
        if err := server.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe(): %s", err)
        }
    }()
    time.Sleep(100 * time.Millisecond)
    return server
}

func TestServerResponds(t *testing.T) {
    server := startTestServer()
    defer server.Close()

    resp, err := http.Get("http://localhost:8081/api")
    if err != nil {
        t.Fatalf("Erro ao fazer requisição: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("esperado status 200, obteve %d", resp.StatusCode)
    }
}