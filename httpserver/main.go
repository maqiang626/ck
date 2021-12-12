package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering rootHandler...")
	fmt.Fprintf(w, "Hello\n")

	// 1. 为 HTTPServer 添加 0-2 秒的随机延时
	rand.Seed(time.Now().UnixNano())
	delayed := rand.Intn(2000)
	time.Sleep(time.Duration(delayed) * time.Millisecond)

	for name, values := range r.Header {
		for _, value := range values {
			w.Header().Add(name, value)
			io.WriteString(w, fmt.Sprintf("%s=%s\n", name, value))
		}
	}

	var VERSION string = os.Getenv("VERSION")
	io.WriteString(w, fmt.Sprintf("VERSION: %s\n", VERSION))

	wc := &ResponseWithRecorder{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}

	fmt.Printf("StatusCode: %d\n", wc.statusCode)

	addr, _ := net.InterfaceAddrs()
	fmt.Printf("IP: %s\n", addr)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthzHandler)

	server := &http.Server{
		Addr:    "localhost:9001",
		Handler: mux,
	}
	server.ListenAndServe()
}
