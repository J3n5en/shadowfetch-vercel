package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/imroc/req/v3"
)

// 需要忽略的HTTP头
var IGNORE_HEADERS = []string{"Origin", "Sentry-Trace", "Host", "Connection", "Content-Length", "X-Forwarded-Host", "Baggage", "X-Real-Ip", "X-Forwarded-For", "Forwarded"}

// Handler 处理所有进入的HTTP请求并将其代理到目标服务器
func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		TARGET_URL         = os.Getenv("TARGET_URL")
		PWD                = os.Getenv("PASSWORD")
		DEBUG_MODE         = os.Getenv("DEBUG_MODE")
		IMPERSONATE_CHROME = os.Getenv("IMPERSONATE_CHROME")
	)
	// 创建一个新的req客户端
	client := req.C()

	// 如果启用了调试模式，则打印请求和响应的详细信息
	if DEBUG_MODE == "1" {
		client.DevMode()
	}
	if IMPERSONATE_CHROME == "1" {
		IGNORE_HEADERS = append(IGNORE_HEADERS, "User-Agent")
		client.ImpersonateChrome()
	}

	// 构建目标URL
	path := r.URL.Path
	if !strings.HasPrefix(path, "/"+PWD) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 从路径中提取目标URL
	targetPath := strings.TrimPrefix(path, "/"+PWD+"/")

	targetURL := TARGET_URL + targetPath

	// 如果有查询参数，添加到目标URL
	if r.URL.RawQuery != "" {
		targetURL += "?" + r.URL.RawQuery
	}

	// 创建一个新的请求
	request := client.R()

	// 复制原始请求的头信息
	for name, values := range r.Header {
		// 跳过一些特定的头，这些头可能会导致问题
		if slices.Contains(IGNORE_HEADERS, name) || strings.HasPrefix(name, "X-Vercel") {
			continue
		}
		for _, value := range values {
			request.SetHeader(name, value)
		}
	}

	// 如果有请求体，复制请求体
	if r.Body != nil {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		request.SetBody(body)
	}

	// 设置请求方法
	var resp *req.Response
	var err error

	switch r.Method {
	case http.MethodGet:
		resp, err = request.Get(targetURL)
	case http.MethodPost:
		resp, err = request.Post(targetURL)
	case http.MethodPut:
		resp, err = request.Put(targetURL)
	case http.MethodDelete:
		resp, err = request.Delete(targetURL)
	case http.MethodPatch:
		resp, err = request.Patch(targetURL)
	case http.MethodHead:
		resp, err = request.Head(targetURL)
	case http.MethodOptions:
		resp, err = request.Options(targetURL)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	// 检查请求是否成功
	if err != nil {
		log.Printf("Error proxying request: %v", err)
		http.Error(w, "Error proxying request", http.StatusInternalServerError)
		return
	}

	// 复制响应头
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// 设置响应状态码
	w.WriteHeader(resp.StatusCode)

	// 复制响应体
	respBody, err := resp.ToBytes()
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	// 写入响应体
	if _, err := w.Write(respBody); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
