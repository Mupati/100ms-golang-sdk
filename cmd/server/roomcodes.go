package main

import (
	"context"
	"encoding/json"
	"net/url"
	"os"
	"strconv"

	"connectrpc.com/connect"
	roomcodesv1 "github.com/Mupati/100ms-golang-server-sdk/gen/roomcodes/v1"
)

type RoomCodesServer struct{}

func (s *RoomCodesServer) GetRoomCodes(ctx context.Context, req *connect.Request[roomcodesv1.GetRoomCodesRequest]) (*connect.Response[roomcodesv1.GetRoomCodesResponse], error) {

	qs := url.Values{}
	qs.Add("role", req.Msg.Filters.Role)
	qs.Add("enabled", strconv.FormatBool(req.Msg.Filters.Enabled))
	qs.Add("limit", strconv.FormatInt(int64(req.Msg.Filters.Limit), 10))
	qs.Add("start", req.Msg.Filters.Start)

	resp, err := PerformHTTPCall(os.Getenv("BASE_URL")+"room-codes/room/"+req.Msg.RoomId+"?"+qs.Encode(), "GET", nil)
	if err != nil {
		return nil, err
	}

	var roomcodes roomcodesv1.GetRoomCodesResponse
	err = json.Unmarshal([]byte(resp), &roomcodes)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(
		&roomcodesv1.GetRoomCodesResponse{
			Limit: roomcodes.Limit,
			Data:  roomcodes.Data,
			Last:  roomcodes.Last,
		},
	)

	return res, nil
}

func (s *RoomCodesServer) CreateRoomCode(ctx context.Context, req *connect.Request[roomcodesv1.CreateRoomCodeRequest]) (*connect.Response[roomcodesv1.CreateRoomCodeResponse], error) {
	return nil, nil
}

func (s *RoomCodesServer) CreateRoomCodeForRole(ctx context.Context, req *connect.Request[roomcodesv1.CreateRoomCodeForRoleRequest]) (*connect.Response[roomcodesv1.CreateRoomCodeForRoleResponse], error) {
	return nil, nil
}

func (s *RoomCodesServer) UpdateRoomCode(ctx context.Context, req *connect.Request[roomcodesv1.UpdateRoomCodeRequest]) (*connect.Response[roomcodesv1.UpdateRoomCodeResponse], error) {
	return nil, nil
}

func (s *RoomCodesServer) CreateShortCodeAuthToken(ctx context.Context, req *connect.Request[roomcodesv1.CreateShortCodeAuthTokenRequest]) (*connect.Response[roomcodesv1.CreateShortCodeAuthTokenResponse], error) {
	return nil, nil
}
