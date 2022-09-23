package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gnunu/elb/usecase"
	"github.com/golang/gddo/httputil/header"
	"k8s.io/klog"
)

func sendRequest(r *usecase.Request, url string) error {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		klog.Error(err.Error())
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.Error(err.Error())
		return err
	}

	req.Header.Set("X-Custom-Header", "request")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		klog.Error(err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}
	return fmt.Errorf("StatusCode=%d", resp.StatusCode)
}

func pickEndpoint(u usecase.Usecase, r *usecase.Request) (string, error) {
	return "", fmt.Errorf("no endpoint found")
}

//// route request to the right endpoint
func serveUsecase(r *usecase.Request) bool {
	klog.Info(fmt.Sprintf("serving request: %v", r))
	klog.Info(fmt.Sprintf("usecases: %v", usecases))
	u, ok := usecases.LookUp(r.Name)
	if !ok {
		klog.Info(fmt.Sprintf("usecase %s not found", r.Name))
		return false
	}
	klog.Info(fmt.Sprintf("serving usecase: %v", u))
	e, err := pickEndpoint(u, r)
	if err == nil {
		sendRequest(r, e)
		return true
	}
	return false
}

//// serve one request to target
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
	klog.Info(fmt.Sprintf("%v", m["usecase"]))
	klog.Info(fmt.Sprintf("%v", m["uri"]))
	if m["usecase"] != "" {
		req := usecase.NewRequest(m["usecase"], m["device"], m["policy"], m["uri"])
		serveUsecase(req)
	}
}

/// receive user request and serve
func StartHttpServer(done chan (struct{})) {
	mux := http.NewServeMux()
	mux.HandleFunc("/request", serveRequest)

	klog.Info(fmt.Sprintf("request server listening at %d", request_port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", request_port), mux)
	klog.Error(fmt.Sprintf("Something fatal happened: %v", err))
	close(done)
}
