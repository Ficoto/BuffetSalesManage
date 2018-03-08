package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResp : application/json common return method
func JSONResp(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	var err = json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
}

// PlainResp : text/plain common return method
func PlainResp(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, v)
}

// VndResp : pplication/vnd.apple.mpegurl common return method
func VndResp(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, v)
}
