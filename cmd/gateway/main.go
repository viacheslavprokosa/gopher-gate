package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	port := ":8080"
	logger.Info("Gateway starting", "port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
