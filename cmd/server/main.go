package main

import (
	"net/http"

	"github.com/Mupati/100ms-golang-server-sdk/gen/room/v1/roomv1connect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {

	room := &RoomServer{}

	mux := http.NewServeMux()
	roomPath, roomHandler := roomv1connect.NewRoomServiceHandler(room)

	mux.Handle(roomPath, roomHandler)

	http.ListenAndServe("localhost:9090", h2c.NewHandler(mux, &http2.Server{}))
}
