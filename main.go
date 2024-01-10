package main

import (
"fmt"
"log"
"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker!

`)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

