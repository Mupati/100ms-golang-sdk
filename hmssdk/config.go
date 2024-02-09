package hmssdk

import (
	"net/http"
)

type HMSConfig struct {
	BaseUrl      string
	AuthBaseUrl  string
	AppAccessKey string
	AppSecret    string
}

type HMSService struct {
	config     *HMSConfig
	httpClient *http.Client
}

func NewHMSService(config *HMSConfig, httpClient *http.Client) *HMSService {
	return &HMSService{
		config:     config,
		httpClient: httpClient,
	}
}
