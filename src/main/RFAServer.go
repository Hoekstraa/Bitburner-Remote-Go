package main

import (
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/grafov/bcast"
	"log"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func websocket(outMessages *bcast.Group) {
	ln, err := net.Listen("tcp", "localhost:"+strconv.Itoa(configuration.RFAPort))
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
			// Allow only localhost, regardless of port
			if strings.Split(string(host), ":")[0] == "localhost" {
				return nil
			}

			fmt.Println(string(host), " rejected")
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

		go func() {
			defer conn.Close()
			messageBus := outMessages.Join()
			for {
				recv := messageBus.Recv()
				var msg = recv.(Message)
				msg.Jsonrpc = "2.0"
				marshal, err := json.Marshal(msg)
				if err != nil {
					log.Printf("JSON parsing error: ", err)
				}
				err = wsutil.WriteServerMessage(conn, ws.OpCode(1), marshal)
				if err != nil {
					log.Printf("Message send error: ", err)
				}
			}
		}()
	}
}
