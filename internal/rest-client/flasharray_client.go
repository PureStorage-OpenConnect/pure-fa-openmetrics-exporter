package client

import (
	"crypto/tls"
	"errors"

	"github.com/go-resty/resty/v2"
)

var UserAgentVersion string = "development"

var FARestUserAgentBase string = "Dev_Pure_FA_OpenMetrics_exporter"

var FARestUserAgent string = FARestUserAgentBase + "/" + UserAgentVersion

type Client interface {
	GetAlerts(filter string) *AlertsList
	GetArrays() *ArraysList
	GetArraysPerformance() *ArraysPerformanceList
	GetConnections() *ConnectionsList
	GetDirectories() *DirectoriesList
	GetDirectoriesPerformance() *DirectoriesPerformanceList
	GetHosts() *HostsList
	GetHostsPerformance() *HostsPerformanceList
	GetHostsBalance() *HostsBalanceList
	GetHardware() *HardwareList
	GetDrive() *DriveList
	GetPods() *PodsList
	GetPodsPerformance() *PodsPerformanceList
	GetVolumes() *VolumesList
	GetVolumesPerformance() *VolumesPerformanceList
}

type FAClient struct {
	EndPoint   string
	ApiToken   string
	RestClient *resty.Client
	ApiVersion string
	XAuthToken string
	XRequestID string
	Error      error
}

func NewRestClient(endpoint string, apitoken string, apiversion string, uagent string, rid string, debug bool) *FAClient {
	type ApiVersions struct {
		Versions []string `json:"version"`
	}
	fa := &FAClient{
		EndPoint:   endpoint,
		ApiToken:   apitoken,
		XRequestID: rid,
		RestClient: resty.New(),
		XAuthToken: "",
	}
	fa.RestClient.SetBaseURL("https://" + endpoint + "/api")
	fa.RestClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	fa.RestClient.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"X-Request-ID": fa.XRequestID,
		"User-Agent":   FARestUserAgent + " (" + uagent + ")",
	})
	if debug {
		fa.RestClient.SetDebug(true)
	}
	//	fa.RestClient.OnRequestLog(func(rl *resty.RequestLog) error {
	//		fmt.Fprintln(os.Stderr, rl)
	//		return nil
	//	})

	result := new(ApiVersions)
	res, err := fa.RestClient.R().
		SetResult(&result).
		Get("/api_version")
	if err != nil {
		fa.Error = err
		return fa
	}
	if res.StatusCode() != 200 {
		fa.Error = errors.New("not a valid FlashArray REST API server")
		return fa
	}
	if len(result.Versions) == 0 {
		fa.Error = errors.New("not a valid FlashArray REST API version")
		return fa
	}
	if apiversion == "latest" {
		fa.ApiVersion = result.Versions[len(result.Versions)-1]
	} else {
		fa.ApiVersion = apiversion
	}
	fa.XRequestID = res.Header().Get("X-Request-ID")
	fa.RestClient.SetBaseURL("https://" + endpoint + "/api/" + fa.ApiVersion)
	fa.RestClient.SetHeader("X-Request-ID", fa.XRequestID)
	res, err = fa.RestClient.R().
		SetHeader("api-token", apitoken).
		Post("/login")
	if err != nil {
		fa.Error = err
		return fa
	}
	if res.StatusCode() != 200 {
		fa.Error = errors.New("failed to login to FlashArray, check API Token")
		return fa
	}
	fa.XAuthToken = res.Header().Get("x-auth-token")
	fa.RestClient.SetHeader("x-auth-token", fa.XAuthToken)
	return fa
}

func (fa *FAClient) Close() *FAClient {
	if fa.XAuthToken == "" {
		return fa
	}
	_, err := fa.RestClient.R().
		SetHeader("x-auth-token", fa.XAuthToken).
		SetHeader("X-Request-ID", fa.XRequestID).
		Post("/logout")
	if err != nil {
		fa.Error = err
	}
	return fa
}

func (fa *FAClient) RefreshSession() *FAClient {
	res, err := fa.RestClient.R().
		SetHeader("api-token", fa.ApiToken).
		SetHeader("X-Request-ID", fa.XRequestID).
		Post("/login")
	if err != nil {
		fa.Error = err
		return fa
	}
	fa.XAuthToken = res.Header().Get("x-auth-token")
	fa.RestClient.SetHeader("x-auth-token", fa.XAuthToken)
	return fa
}
