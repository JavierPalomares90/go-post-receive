package main

import (
	"fmt"
	"net/http"
)

func receive(w http.ResponseWriter, req *http.Request) {
	// Make sure we have a POST request (not GET)
	switch req.Method {
	case "POST":
		// Parse the body
		fmt.Fprintf(w, "Got your post request\n")
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", req.PostForm)
		string1 := req.FormValue("string1")
		string2 := req.FormValue("string2")
		fmt.Fprintf(w, " you sent a value of %s for string1 and %s for string2\n", string1, string2)

	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.\n")
	}
}

func main() {
	// all requests handled by receive function
	http.HandleFunc("/", receive)
	// Listen on port 8090
	fmt.Println("New Server is listening")
	http.ListenAndServe(":8090", nil)
}
