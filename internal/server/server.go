package server

import (
	"CICDRef/internal/config"
	"CICDRef/internal/handlers"
	"net/http"
)

func NewServer() *http.Server {
	cfg := config.Load()
	h := handlers.New(cfg.Message)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/message", h.MessageHandler)

	return &http.Server{Addr: ":" + cfg.Port, Handler: mux}
}
