package main

import (
	"net/http"

	"github.com/njhale/buffalo/gen/buffalo/brick/v1/brickv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type BrickServer struct {
	brickv1connect.UnimplementedBrickServiceHandler
}

func main() {
	server := &BrickServer{} // TODO: Replace with working implementation.
	mux := http.NewServeMux()
	path, handler := brickv1connect.NewBrickServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
