package main

import (
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net"
	"net/http"
	"runtime"
)

func websocket(c chan Message) {
	ln, err := net.Listen("tcp", "localhost:12525")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare handshake header writer from http.Header mapping.
	header := ws.HandshakeHeaderHTTP(http.Header{
		"X-Go-Version":  []string{runtime.Version()},
		"X-RFA-Version": []string{"0.1"},
		"X-creator":     []string{"Zoë Hoekstra"},
	})

	u := ws.Upgrader{
		OnHost: func(host []byte) error {
			if string(host) == "localhost" || true {
				return nil
			}
			return ws.RejectConnectionError(
				ws.RejectionStatus(403),
				ws.RejectionHeader(ws.HandshakeHeaderString(
					"X-Want-Host: localhost\r\n",
				)),
			)
		},
		OnBeforeUpgrade: func() (ws.HandshakeHeader, error) {
			return header, nil
		},
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		_, err = u.Upgrade(conn)
		if err != nil {
			log.Printf("upgrade error: %s", err)
		}

		for {
			msg := <-c
			msg.Jsonrpc = "2.0"
			marshal, err := json.Marshal(msg)
			if err != nil {
				log.Printf("JSON parsing error: #{err}")
			}
			err = wsutil.WriteServerMessage(conn, ws.OpCode(1), marshal)
			if err != nil {
				log.Printf("Message send error: #{err}")
			}
		}
	}
}
