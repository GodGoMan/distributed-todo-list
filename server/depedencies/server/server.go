package server

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/GodGoMan/dts-server/depedencies/task"
)

type server struct {
	ln   net.Listener
	conn net.Conn
	addr string
	todo chan task.Task
}

func NewServer(addr string) *server {
	return &server{
		addr: addr,
		todo: make(chan task.Task),
	}
}

func (s *server) Start() {
	var err error
	s.ln, err = net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("LISTNER ERROR --> %e", err)
		return
	}
	defer s.ln.Close()
	s.conn, err = s.ln.Accept()
	if err != nil {
		fmt.Printf("ACCEPT ERROR --> %e", err)
	}
	defer s.conn.Close()
	s.getTask()
	t := <-s.todo
	fmt.Println(t)
}

func (s *server) getTask() {
	dcdr := gob.NewDecoder(s.conn)
	var t task.Task
	dcdr.Decode(&t)
	fmt.Println(t)
	s.todo <- t
}
