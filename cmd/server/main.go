package main

import (
	"net/http"

	"github.com/Mupati/100ms-golang-server-sdk/gen/room/v1/roomv1connect"
	"github.com/Mupati/100ms-golang-server-sdk/gen/roomcodes/v1/roomcodesv1connect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {

	room := &RoomServer{}
	roomcodes := &RoomCodesServer{}

	mux := http.NewServeMux()
	roomPath, roomHandler := roomv1connect.NewRoomServiceHandler(room)
	roomcodesPath, roomcodesHandler := roomcodesv1connect.NewRoomCodesServiceHandler(roomcodes)

	mux.Handle(roomPath, roomHandler)
	mux.Handle(roomcodesPath, roomcodesHandler)

	http.ListenAndServe("localhost:9090", h2c.NewHandler(mux, &http2.Server{}))
}
