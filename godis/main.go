package main

import (
	"net"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// Configure logger

	// Equivalent of Lshortfile
	// https://github.com/rs/zerolog?tab=readme-ov-file#customize-automatic-field-names
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	log.Logger = log.With().Caller().Logger()

	// Colorize output for terminal
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	// Initialize TCP Server
	addr := "localhost:7000"
	log.Info().Msgf("starting tcp server on %s", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal().Msgf("Error starting TCP Server: %s", err)
	}
	defer l.Close()

	// Keep the server running forever
	for {
		// Accept incoming connection
		log.Info().Msgf("waiting for incoming connection")

		conn, err := l.Accept()
		log.Info().Msgf("incoming connection from %s", conn.RemoteAddr().String())

		if err != nil {
			log.Error().Msgf("error accepting incoming connection: %s", err)
		}

		go handleConnection(conn)
	}
}

// handleConnection will handle the incoming TCP connection
func handleConnection(conn net.Conn) {
	// Read data from the connection
	var r []byte
	_, err := conn.Read(r)
	if err != nil {
		log.Error().Msgf("error reading data from connection: %s", err)
		return
	}
	defer conn.Close()
}
