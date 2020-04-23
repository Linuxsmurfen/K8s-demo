// A webserver that publish...
//	/		Default page the publish "Demo + $APP_TEXT"
//	/data.json	Publish the pod specific enviroment data in json format
//	/client		A simple webpage that will graph the data
//	/_healthz	Healthcheck "OK"

// The K8S_ is automaticlly set by the pod deployment
// The APP_ is manually set

// Linuxsmurfen
// 2020-04-23

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
const loggfilen string = "demo.log"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Demo " + os.Getenv("APP_TEXT"))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    fmt.Fprintf(w, response)
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "demo.html")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "ok")
}

func main() {
	//Setup logging to file
	file, err := os.OpenFile(loggfilen, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	log.SetOutput(file)
	log.SetPrefix(os.Getenv("K8S_POD_NAME") + "::")

	//Build the json respond
    response = "{\n  \"Version\": \"" + os.Getenv("APP_VERSION") +
                "\",\n  \"Text\": \"" + os.Getenv("APP_TEXT") +
                "\",\n  \"Nodename\": \"" + os.Getenv("K8S_NODE_NAME") +
                "\",\n  \"Podname\": \"" + os.Getenv("K8S_POD_NAME") +
                "\",\n  \"Namespace\": \"" + os.Getenv("K8S_POD_NAMESPACE") +
                "\",\n  \"PodIP\": \"" + os.Getenv("K8S_POD_IP") +
                "\",\n  \"HostIP\": \"" + os.Getenv("K8S_HOST_IP") +
                "\",\n  \"Serviceaccount\": \"" + os.Getenv("K8S_POD_SERVICE_ACCOUNT") +
                "\"\n}\n"

	//URI handlers			
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/data.json", jsonHandler)
	http.HandleFunc("/client", clientHandler)
    http.HandleFunc("/_healthz", healthHandler)

	//Logging
	log.Printf("Listening on port: 8080 ...\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
