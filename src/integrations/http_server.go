package integrations

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func loghandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/log" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		fmt.Fprintf(w, "Only POST method is supported")
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "err : %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{status : "ok"}`)
	go func(http.ResponseWriter) {
		config.StreamLogToAll(string(body))
		if err != nil {
			fmt.Fprintf(w, `{status: "not ok", "error": %v}`, err)
			return
		}
	}(w)

}

var config, _ = ReadLogConfig("./conf/datasage.yaml")

func RunServer() {

	http.HandleFunc("/log", loghandler)
	fmt.Printf("HTTP Server listening on :8080\n")
	fmt.Printf("Send logs to 127.0.0.1:8080/log endpoint\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
