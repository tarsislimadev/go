
package main

import "os"
import "net/http"

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func RequestHandler() (string, error) {
	return "{}", nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := RequestHandler()
		check(err)
		w.Write([]byte(response))
	})

	host := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")
	err := http.ListenAndServe(host + ":" + port, nil)
	check(err)
}
