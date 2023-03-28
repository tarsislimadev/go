
package main

import "os"
import "flag"
import "net/url"
import "net/http"

var (
	ROOT = *flag.String("path", os.Getenv("DATA_PATH"), "The path of the data.")
	HOST = *flag.String("host", os.Getenv("HOSTNAME"), "The host name")
	PORT = *flag.String("port", os.Getenv("PORT"), "The host port")

	INDEX_NAME = "private/login"
)

func RequestHandler(query url.Values) (string, error) {
	obj := Index{ ROOT, INDEX_NAME }.New() 
	err := obj.WriteArrayMap(query)
	if err != nil { return "", err }
	return obj.ObjectName, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := RequestHandler(r.URL.Query())
		CheckError(err)
		w.Write([]byte(response))
	})

	CheckError(http.ListenAndServe(HOST + ":" + PORT, nil))
}
