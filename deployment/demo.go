// Compile with...
// go build -o demo demo.go

package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
)

var response string

func defaultHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Demo " + os.Getenv("K8S_TEXT"))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	    w.Header().Set("Content-Type", "application/json")
	    w.Header().Set("Access-Control-Allow-Origin", "*")
	    
        fmt.Fprintf(w, response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "ok")
}

func main() {
        response = "{\n  \"Version\": \"" + os.Getenv("K8S_VERSION") +
                    "\",\n  \"Text\": \"" + os.Getenv("K8S_TEXT") +
                    "\",\n  \"Nodename\": \"" + os.Getenv("K8S_NODE_NAME") +
                    "\",\n  \"Podname\": \"" + os.Getenv("K8S_POD_NAME") +
                    "\",\n  \"Namespace\": \"" + os.Getenv("K8S_POD_NAMESPACE") +
                    "\",\n  \"PodIP\": \"" + os.Getenv("K8S_POD_IP") +
                    "\",\n  \"HostIP\": \"" + os.Getenv("K8S_HOST_IP") +
                    "\",\n  \"Serviceaccount\": \"" + os.Getenv("K8S_POD_SERVICE_ACCOUNT") +
                    "\"\n}\n"

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/data.json", jsonHandler)
        http.HandleFunc("/_healthz", healthHandler)

        log.Printf("Listening on :8080 at ...\n%s", response)
        log.Fatal(http.ListenAndServe(":8080", nil))
}
