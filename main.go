package main

import (
	"context"
	"log"
	"os"

	// roomv1 "github.com/Mupati/100ms-golang-server-sdk/gen/room/v1"
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

	// Create a new room
	aNewRoom, err := roomClient.CreateRoom(context.Background(), &hmssdk.CreateRoomRequest{
		Room: &hmssdk.HMSRoom{
			Name:        "mupati-new-room-from-sdk",
			Description: "This is a new room I'm creating from my SDK",
		},
	})

	if err != nil {
		log.Panic("There was an error creating the room: ", err)
	}
	log.Println("I am creating a new room: ", aNewRoom.Msg.Room)

	updateRoom, err := roomClient.UpdateRoom(context.Background(), &hmssdk.UpdateRoomRequest{
		RoomId: "65b03920e01621cac2c13210",
		Room: &hmssdk.HMSRoom{
			Description: "This is an updated description",
		},
	})

	if err != nil {
		log.Panic("There was an error updating the room: ", err)
	}
	log.Println("I am updating a new room: ", updateRoom.Msg.Room)

}
