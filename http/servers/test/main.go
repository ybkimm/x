package main

import (
	"fmt"
	"net/http"

	"go.ybk.im/x/http/servers"
)

var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It Works!"))
}

func main() {
	fmt.Println("Starting server at port 7089")
	err := servers.HTTP(":7089", handler)
	if err != nil {
		panic(err)
	}
}
