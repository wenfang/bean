package client

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	rsp, err := Get("http://www.baidu.com", ContentJSON(), Timeout(10e9))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rsp.Body))
}
