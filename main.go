package main

import (
	"net/http"
	"os"

	"github.com/olahol/melody"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	m := melody.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		if err := m.HandleRequest(w, r); err != nil {
			log.Error().Msg("Unexpected error has occurred.")
		}
	})

	m.HandleMessage(func(s *melody.Session, b []byte) {
		log.Info().Msg(string(b))
		m.Broadcast(b)
	})

	http.ListenAndServe(":5001", nil)
}
