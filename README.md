# 100ms Golang Server SDK

An SDK wrapper for the 100ms Server API
This is nowhere near ready. Expect a lot of breaking changes.

How to use

```
package main

import (
	"log"
	"os"

	hmssdk "github.com/Mupati/100ms-golang-server-sdk/hmssdk"
)

func main() {

	hmsService := hmssdk.NewHMSService(&hmssdk.HMSConfig{
		BaseUrl:      "http://localhost:9090",
		AppSecret:    os.Getenv("APP_SECRET"),
		AppAccessKey: os.Getenv("APP_ACCESS_KEY"),
	}, http.DefaultClient)

	// Trying out room creation
	token, err := hmsService.CreateJoinRoomToken(&hmssdk.JoinRoomParam{
		UserId:    "bra-kofi-mupati",
		RoomId:    "655356026c4fc3aa925cb9e5",
		Role:      "guest",
		ExpiresIn: 84000,
	})

	if err != nil {
		log.Println("There was an error during token creation: ", err)
		return
	}
	log.Println("This is the token: ", token)


	roomClient := hmsService.Room()

	// Get the details of a room
	newRoom, err := roomClient.GetRoom(context.Background(), &hmssdk.GetRoomRequest{
		RoomId: "655356026c4fc3aa925cb9e5",
	})
	if err != nil {
		log.Println("There was an error getting the room: ", err)
		return
	}
	log.Println("This is a new room: ", newRoom.Msg.Room)


}

```
