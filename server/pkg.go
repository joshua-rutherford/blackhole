package server

import (
	"fmt"
	"io"
	"net/http"
)

func Start(host string, port int) error {
	http.HandleFunc("/dev/null", func(w http.ResponseWriter, r *http.Request) {
		bytes := make([]byte, 1024)
		for {
			_, err := r.Body.Read(bytes)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					w.WriteHeader(500)
					w.Write([]byte{})
					return
				}
			}
		}
		w.WriteHeader(200)
		w.Write([]byte{})
	})
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}
