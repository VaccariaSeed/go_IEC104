package go_IEC104

import "testing"

func TestServer(t *testing.T) {
	server := BuildIEC104Server(2404, 1, "./doc/信号点配置表.xlsx")
	err := server.Open()
	if err != nil {
		panic(err)
	}
}
