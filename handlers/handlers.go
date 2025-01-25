package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HttpFunc(endpoint string, f func(data any) any) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dataByte, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}

		value := f(dataByte)
		fmt.Fprintln(w, value)
	})

	http.HandleFunc(endpoint, handler)
}

func ServerStart(addr string) {
	http.ListenAndServe(addr, nil)
}
