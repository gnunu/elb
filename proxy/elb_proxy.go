package proxy

import (
	"fmt"
	"net/http"

	"k8s.io/klog"
)

func ServeRequest(w http.ResponseWriter, r *http.Request) {
	klog.Info("uri", r.RequestURI)
	fmt.Fprint(w, "OK")
}

func start_proxy() {
	http.HandleFunc("/", ServeRequest)
}
