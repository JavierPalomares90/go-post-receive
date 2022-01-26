Example of how to handle post request in Go.

To Run the server:


go run receive.go &

To send a post request with 2 parameters.

 curl -d "string1=foo&string2=bar" -X POST localhost:8090