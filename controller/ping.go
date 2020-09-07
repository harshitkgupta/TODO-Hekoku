package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"first/views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
		t2 := time.Now()

		fmt.Println(t2.Nanosecond() - t1.Nanosecond())
	}
}
