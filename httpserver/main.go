package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
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

	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	for name, values := range r.Header {
		for _, value := range values {
			w.Header().Add(name, value)
			io.WriteString(w, fmt.Sprintf("%s=%s\n", name, value))
		}
	}

	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	var VERSION string = os.Getenv("VERSION")
	io.WriteString(w, fmt.Sprintf("VERSION: %s\n", VERSION))

	// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	wc := &ResponseWithRecorder{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}

	fmt.Printf("StatusCode: %d\n", wc.statusCode)

	addr, _ := net.InterfaceAddrs()
	fmt.Printf("IP: %s\n", addr)
}

// 4.当访问 localhost/healthz 时，应返回200
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
