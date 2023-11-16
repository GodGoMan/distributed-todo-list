package client

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/GodGoMan/dts/dependencies/task"
)

type client struct {
	addr  string
	owner string
	conn  net.Conn
	tsk   task.Task
}

func NewClient(addr, owner string) *client {
	return &client{
		addr:  addr,
		owner: owner,
	}
}

func (c *client) ConnecToServe() {
	var err error
	c.conn, err = net.Dial("tcp", c.addr)
	if err != nil {
		fmt.Printf("DIAL ERROR --> %e", err)
		return
	}
	c.startANewTask()
	c.sendTaskToServer()
}

func (c *client) startANewTask() {
	rd := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task name :")
	tn_r, err := rd.ReadString('\n')
	if err != nil {
		fmt.Printf("READ ERROR --> %e", err)
		return
	}
	tn := strings.TrimSuffix(tn_r, "\n")
	fmt.Print("Enter Task Description :")
	td_r, err := rd.ReadString('\n')
	if err != nil {
		fmt.Printf("READ ERROR --> %e", err)
		return
	}
	td := strings.TrimSuffix(td_r, "\n")
	c.tsk = *task.NewTask(c.owner, tn, td)
}

func (c *client) sendTaskToServer() {
	encdr := gob.NewEncoder(c.conn)
	err := encdr.Encode(c.tsk)
	if err != nil {
		fmt.Printf("ENCODER ERROR --> %e", err)
	}
}
