package ipc

import "testing"

type EchoServer struct{}

func (server *EchoServer) Handle(method, param string) *Response {
	return &Response{"Echo:" + method, param}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("From client1", "")
	resp2, _ := client2.Call("From client2", "")

	if resp1.Code != "Echo:From client1" || resp2.Code != "Echo:From client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
