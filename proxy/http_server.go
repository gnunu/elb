package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/gddo/httputil/header"
	"k8s.io/klog"
)

func serveRequest(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	m := make(map[string]string)
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&m)
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		http.Error(w, msg, http.StatusBadRequest)
	}
	klog.Info(fmt.Sprintf("%v", m))
}

func StartHttpServer(done chan (struct{})) {
	mux := http.NewServeMux()
	mux.HandleFunc("/request", serveRequest)

	klog.Info(fmt.Sprintf("request server listening at %d", request_port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", request_port), mux)
	klog.Error(fmt.Sprintf("Something fatal happened: %v", err))
	close(done)
}
