# 100ms Golang Server SDK

An SDK wrapper for the 100ms Server API
This is nowhere near ready. Expect a lot of breaking changes.

How to use

```
package main

import (
	"log"
	"os"

	hmssdk "github.com/Mupati/100ms-golang-server-sdk/sdk"
)

func main() {
	log.Println("==========Trying out the SDK==========")

	roomClient := hmssdk.NewRoomServiceClient(os.Getenv("BASE_URL"), os.Getenv("APP_ACCESS_KEY"), os.Getenv("APP_SECRET"))

	log.Println("....roomClient....: ", roomClient)

	// Trying out room creation

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
}

```
