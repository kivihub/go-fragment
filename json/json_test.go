package json

import (
	"encoding/json"
	"testing"
)

func TestUnmarshal(test *testing.T) {
	tar := []NodeMsg{}
	err := json.Unmarshal([]byte("[{\"status\":3,\"message\":\"success\",\"url\":\"https://baidu.com\"}]"), &tar)
	if err != nil {
		test.Error(err)
	} else {
		test.Logf("%+v", tar)
	}
}

type NodeMsg struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	url     string `json:"url"`
}
