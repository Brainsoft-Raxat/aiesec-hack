package main

import (
	"flag"
	"github.com/Brainsoft-Raxat/aiesec-hack/internal/app"
)

var (
	httpAddr string
	envParse bool
	envPath  string
)

func init() {
	flag.StringVar(&httpAddr, "http.addr", ":8080", "HTTP listen address only port :8080")
	flag.BoolVar(&envParse, "env.parse", true, "Whether parse envs from file or not")
	flag.StringVar(&envPath, "env.path", "internal/app/config/local.env", "Path to env file")
}

func main() {
	flag.Parse()

	files := make([]string, 0)
	if envParse {
		files = append(files, envPath)
	}

	app.Run(files...)
}