package main

import (
	"bitel-cetm/config"
	"bitel-cetm/server"
	"fmt"
	"net/http"
)

// func init() {
// 	config.Init()
// }

func main() {
	s := server.NewServer()
	fmt.Println("Start Listen and server ", config.GetHostPort())
	err := http.ListenAndServe(config.GetHostPort(), *s)
	if nil != err {
		fmt.Println("Listen And Server ERROR: ", err.Error())
	}
}
