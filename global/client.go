package global

import (
	"net"
	"net/http"
	"time"

	"github.com/pathfinder-cm/pathfinder-go-client/ext"
)

func NewClient() (ext.Client, error) {
	p, err := GetCurrentProfile()
	if err != nil {
		return nil, err
	}

	h := &http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 60 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 60 * time.Second,
		},
	}
	paths := map[string]string{
		"GetNodes":      "api/v1/ext_app/nodes",
		"GetNode":       "api/v1/ext_app/nodes",
		"GetContainers": "api/v1/ext_app/containers",
		"GetContainer":  "api/v1/ext_app/container",
	}
	return ext.NewClient(p.Cluster, p.Token, h, p.Server, paths), nil
}
