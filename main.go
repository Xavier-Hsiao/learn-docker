package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	multiplexer := http.NewServeMux()

	const addr = ":8000"

	multiplexer.HandleFunc("/", handlePage)

	server := http.Server {
		Handler: multiplexer,
		Addr: addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	fmt.Println("server started on port ", addr)
	err := server.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	
	const page = `
	<html>
		<body>
			<p>Hello Gophoer!</p>
		</body>
	</html>
	`

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(page))
}