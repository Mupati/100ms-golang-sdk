package hmssdk

import (
	"context"
	"net/http"
	"time"

	roomv1 "github.com/Mupati/100ms-golang-server-sdk/gen/room/v1"
	"github.com/Mupati/100ms-golang-server-sdk/gen/room/v1/roomv1connect"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"connectrpc.com/connect"
)

type RoomServiceClient struct {
	room roomv1connect.RoomServiceClient
}

func NewRoomClient(config *HMSConfig, httpClient *http.Client) *RoomServiceClient {
	return &RoomServiceClient{
		room: roomv1connect.NewRoomServiceClient(httpClient, config.BaseUrl),
	}

}

func (s *HMSService) Room() *RoomServiceClient {
	return NewRoomClient(s.config, s.httpClient)
}

type JoinRoomParam struct {
	UserId    string
	RoomId    string
	Role      string
	ExpiresIn int
}

func (s *HMSService) CreateJoinRoomToken(rb *JoinRoomParam) (string, error) {

	var expiresIn uint32

	if rb.ExpiresIn == 0 {
		expiresIn = uint32(24 * 3600)
	} else {
		expiresIn = uint32(rb.ExpiresIn)
	}

	mySigningKey := []byte(s.config.AppSecret)
	now := uint32(time.Now().UTC().Unix())
	exp := now + expiresIn
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": s.config.AppAccessKey,
		"type":       "app",
		"version":    2,
		"room_id":    rb.RoomId,
		"user_id":    rb.UserId,
		"role":       rb.Role,
		"jti":        uuid.New().String(),
		"iat":        now,
		"exp":        exp,
		"nbf":        now,
	})

	return token.SignedString(mySigningKey)
}

type HMSRoom = roomv1.HMSRoom
type GetRoomRequest = roomv1.GetRoomRequest

func (c *RoomServiceClient) GetRoom(ctx context.Context, data *GetRoomRequest) (*connect.Response[roomv1.GetRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.room.GetRoom(ctx, req)
}

type ListRoomsRequest = roomv1.ListRoomsRequest
type ListRoomsFilters = roomv1.ListRoomsFilters

func (c *RoomServiceClient) ListRooms(ctx context.Context, data *ListRoomsRequest) (*connect.Response[roomv1.ListRoomsResponse], error) {
	req := connect.NewRequest(data)
	return c.room.ListRooms(ctx, req)
}

type CreateRoomRequest = roomv1.CreateRoomRequest

func (c *RoomServiceClient) CreateRoom(ctx context.Context, data *CreateRoomRequest) (*connect.Response[roomv1.CreateRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.room.CreateRoom(ctx, req)

}

type UpdateRoomRequest = roomv1.UpdateRoomRequest

func (c *RoomServiceClient) UpdateRoom(ctx context.Context, data *UpdateRoomRequest) (*connect.Response[roomv1.UpdateRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.room.UpdateRoom(ctx, req)

}

type EnableRoomRequest = roomv1.EnableRoomRequest

func (c *RoomServiceClient) EnableRoom(ctx context.Context, data *EnableRoomRequest) (*connect.Response[roomv1.EnableRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.room.EnableRoom(ctx, req)

}

type DisableRoomRequest = roomv1.DisableRoomRequest

func (c *RoomServiceClient) DisableRoom(ctx context.Context, data *DisableRoomRequest) (*connect.Response[roomv1.DisableRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.room.DisableRoom(ctx, req)
}
