package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/gorilla/websocket"

	"github.com/kyoukaya/subsonar/pkg/subsonar"
)

var target = flag.String("target", "wss://api.ffxivsonar.com/ws", "sonar's url")
var logfile = flag.String("logfile", "", "logfile")

func main() {
	flag.Parse()
	log.SetFlags(0)

	if *logfile != "" {
		f, err := os.OpenFile(*logfile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatal("failed to open logfile:", err)
		}
		log.SetOutput(io.MultiWriter(os.Stdout, f))
	}

	ws, _, err := websocket.DefaultDialer.Dial(*target, nil)
	if err != nil {
		log.Fatalf("failed to dial %s: %v", *target, err)
	}

	sc := subsonar.New(ws)
	defer ws.Close()
	sc.Run()
}
