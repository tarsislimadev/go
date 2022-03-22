
package main

import "os"
import "net/url"
import "net/http"

var (
	ROOT = os.Getenv("DATA_PATH")
	HOST = os.Getenv("HOSTNAME")
	PORT = os.Getenv("PORT")

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
