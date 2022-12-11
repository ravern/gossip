package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Fprintln(os.Stderr, "fatal error: failed to initialise logger")
		os.Exit(1)
	}
	defer logger.Sync()
	logger.Sugar()

	router := chi.NewRouter()

	logger.Info("server listening on :8787...")
	err = http.ListenAndServe(":8787", router)
	if err != nil {
		logger.Error("failed to start server", zap.String("err", err.Error()))
	}
}
