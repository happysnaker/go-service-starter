package httpserver

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/happysnaker/go-service-starter/internal/buildinfo"
	"github.com/happysnaker/go-service-starter/internal/config"
)

type Server struct {
	cfg    config.Config
	logger *slog.Logger
	http   *http.Server
}

func New(cfg config.Config, logger *slog.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"service": cfg.ServiceName,
			"message": "go-service-starter is running",
		})
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
	})
	mux.HandleFunc("/version", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"service": cfg.ServiceName,
			"version": buildinfo.Version,
			"commit":  buildinfo.Commit,
			"builtAt": buildinfo.BuiltAt,
		})
	})

	h := requestLoggingMiddleware(logger, mux)

	return &Server{
		cfg:    cfg,
		logger: logger,
		http: &http.Server{
			Addr:              cfg.HTTPAddr,
			Handler:           h,
			ReadHeaderTimeout: 2 * time.Second,
			ReadTimeout:       cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       60 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	s.logger.Info("http server starting", "addr", s.cfg.HTTPAddr, "service", s.cfg.ServiceName)
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down http server")
	return s.http.Shutdown(ctx)
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func requestLoggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		logger.Info("http request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rec.status,
			"duration", time.Since(start).String(),
			"remote_addr", r.RemoteAddr,
		)
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
