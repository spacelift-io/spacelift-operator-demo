package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := ":8888"
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.Method, request.URL.String())
		if request.URL.Path != "/" {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		secrets := os.Environ()
		body := ""
		for i := 0; i < len(secrets); i++ {
			body += fmt.Sprintf("%s\n", secrets[i])
		}
		_, _ = writer.Write([]byte(body))
	})
	log.Printf("Listening on %s\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}
}
