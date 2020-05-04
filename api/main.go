package main

import (
	"github.com/vallikkumar/todo/api/app"
)

func main() {
	// init app server
	s := app.NewServer()
	// serve server
	s.Serve()
}
