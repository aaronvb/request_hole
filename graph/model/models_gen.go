// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ServerInfo struct {
	RequestAddress string            `json:"request_address"`
	RequestPort    int               `json:"request_port"`
	WebPort        int               `json:"web_port"`
	ResponseCode   int               `json:"response_code"`
	BuildInfo      map[string]string `json:"build_info"`
}
