package main

import (
	"ekz/tools"
	"ekz/tree"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")
var fileName = "cmd\\server\\tree.json"

func processorProvider() tree.Processor {
	return &tree.MultiplyNodes{}
}

func writerProvider() io.Writer {
	return &tools.FileWriter{FileName: fileName}
}

func main() {
	// Parse command line arguments. Port number may be defined with "-p" flag.
	flag.Parse()

	// Create the server.
	if server, err := ComposeApiServer(HTTPPortNumber(*httpPortNumber)); err == nil {
		// Start it.
		go func() {
			log.Println("Starting server...")

			err := server.start()
			if err == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}
		}()

		// Wait for Ctrl-C signal.
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if err := server.stop(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	} else {
		log.Fatalf("Cannot initialize server: %s", err)
	}
}
