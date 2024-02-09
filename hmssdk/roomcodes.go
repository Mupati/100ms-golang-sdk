package hmssdk

import (
	"net/http"

	"github.com/Mupati/100ms-golang-server-sdk/gen/roomcodes/v1/roomcodesv1connect"
)

type RoomCodesServiceClient struct {
	roomcodes roomcodesv1connect.RoomCodesServiceClient
}

func NewRoomCodesClient(config *HMSConfig, httpClient *http.Client) *RoomCodesServiceClient {
	return &RoomCodesServiceClient{
		roomcodes: roomcodesv1connect.NewRoomCodesServiceClient(httpClient, config.BaseUrl),
	}

}

func (s *HMSService) RoomCodes() *RoomCodesServiceClient {
	return NewRoomCodesClient(s.config, s.httpClient)
}
