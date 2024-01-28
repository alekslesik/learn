package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))

	ts.Start()
	defer ts.Close()

	client := ts.Client()

	res1, err := client.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}

	r, err := io.ReadAll(res1.Body)
	res1.Body.Close()

	fmt.Println(string(r))

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}
