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

type HMSRoomService struct {
	config      *HMSConfig
	roomService roomv1connect.RoomServiceClient
}

func NewRoomServiceClient(url, appAccessKey, appSecret string) *HMSRoomService {
	client := roomv1connect.NewRoomServiceClient(http.DefaultClient, url)
	return &HMSRoomService{
		roomService: client,
		config: &HMSConfig{
			BaseUrl:      url,
			AppAccessKey: appAccessKey,
			AppSecret:    appSecret,
		},
	}
}

type JoinRoomParam struct {
	UserId    string
	RoomId    string
	Role      string
	ExpiresIn int
}

func (c *HMSRoomService) CreateJoinRoomToken(rb *JoinRoomParam) (string, error) {

	var expiresIn uint32

	if rb.ExpiresIn == 0 {
		expiresIn = uint32(24 * 3600)
	} else {
		expiresIn = uint32(rb.ExpiresIn)
	}

	mySigningKey := []byte(c.config.AppSecret)
	now := uint32(time.Now().UTC().Unix())
	exp := now + expiresIn
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": c.config.AppAccessKey,
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

func (c *HMSRoomService) GetRoom(ctx context.Context, data *GetRoomRequest) (*connect.Response[roomv1.GetRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.roomService.GetRoom(ctx, req)
}

type ListRoomsRequest = roomv1.ListRoomsRequest
type ListRoomsFilters = roomv1.ListRoomsFilters

func (c *HMSRoomService) ListRooms(ctx context.Context, data *ListRoomsRequest) (*connect.Response[roomv1.ListRoomsResponse], error) {
	req := connect.NewRequest(data)
	return c.roomService.ListRooms(ctx, req)
}

type CreateRoomRequest = roomv1.CreateRoomRequest

func (c *HMSRoomService) CreateRoom(ctx context.Context, data *CreateRoomRequest) (*connect.Response[roomv1.CreateRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.roomService.CreateRoom(ctx, req)

}

type UpdateRoomRequest = roomv1.UpdateRoomRequest

func (c *HMSRoomService) UpdateRoom(ctx context.Context, data *UpdateRoomRequest) (*connect.Response[roomv1.UpdateRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.roomService.UpdateRoom(ctx, req)

}

type EnableRoomRequest = roomv1.EnableRoomRequest

func (c *HMSRoomService) EnableRoom(ctx context.Context, data *EnableRoomRequest) (*connect.Response[roomv1.EnableRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.roomService.EnableRoom(ctx, req)

}

type DisableRoomRequest = roomv1.DisableRoomRequest

func (c *HMSRoomService) DisableRoom(ctx context.Context, data *DisableRoomRequest) (*connect.Response[roomv1.DisableRoomResponse], error) {
	req := connect.NewRequest(data)
	return c.roomService.DisableRoom(ctx, req)

}
