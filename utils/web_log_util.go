package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type WriterProxy struct {
	http.ResponseWriter
	StatusCode int
}

func NewWriterProxy(w http.ResponseWriter) *WriterProxy {
	return &WriterProxy{w, http.StatusOK}
}

func (proxy *WriterProxy) WriteHeader(statusCode int) {
	proxy.StatusCode = statusCode
	proxy.ResponseWriter.WriteHeader(statusCode)
}

// I did not found a function transfer Duration to Millisecond , oh my!
func durationToMs(d time.Duration) string {
	ms := d / time.Millisecond
	nsec := d % time.Millisecond
	return fmt.Sprintf("%.4fms", float64(ms)+float64(nsec)*1e-6)
}

func WebLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		XiaoUserAgent := r.Header.Get("Xiao-User-Agent")
		if XiaoUserAgent == "" {
			XiaoUserAgent = "unknown"
		}
		start := time.Now()
		writerProxy := NewWriterProxy(w)
		handler.ServeHTTP(writerProxy, r)
		log.Println(
			XiaoUserAgent,
			r.Proto,
			r.Method,
			r.URL,
			writerProxy.StatusCode,
			durationToMs(time.Since(start)),
		)
	})
}
