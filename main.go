package main

import (
	"context"
	"log"
	"os"

	hmssdk "github.com/Mupati/100ms-golang-server-sdk/sdk"
)

func main() {
	// Create a client
	roomClient := hmssdk.NewRoomServiceClient("http://localhost:9090", os.Getenv("APP_ACCESS_KEY"), os.Getenv("APP_SECRET"))

	// Generate the token to create a room
	token, err := roomClient.CreateJoinRoomToken(&hmssdk.JoinRoomParam{
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

	// Get the details of a room
	newRoom, err := roomClient.GetRoom(context.Background(), &hmssdk.GetRoomRequest{
		RoomId: "655356026c4fc3aa925cb9e5",
	})
	if err != nil {
		log.Println("There was an error getting the room: ", err)
		return
	}
	log.Println("This is a new room: ", newRoom.Msg.Room)

	// Get all available rooms
	listOfRooms, err := roomClient.ListRooms(context.Background(), &hmssdk.ListRoomsRequest{})
	if err != nil {
		log.Println("There was an error getting the room: ", err)
		return
	}
	log.Println("This is a list of rooms: ", listOfRooms.Msg.Rooms)

}
