package cg

import "fmt"

type Player struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	Exp   int    `json:"exp"`
	Room  int    `json:"room"`

	mq chan *Message // Waiting for new message
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}
	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "Received message:", msg)
		}
	}(player)

	return player
}
