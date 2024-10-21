package port

import (
	"math/rand"

	"log"
)

type InPort interface {
	Read() (int64, error)
}

type OutPort interface {
	Write(transactionId int64, val int64) error
}

type port struct {
	num int64
}

func (i *port) Read() (int64, error) {
	return int64(rand.Int() % 2), nil
}

func (o *port) Write(transactionId int64, val int64) error {
	log.Printf("Port: %d, transaction: %d, value: %d", o.num, transactionId, val)
	return nil
}

func NewInPort(num int64) (InPort, error) {
	log.Printf("In Port %d created", num)
	return &port{num: num}, nil
}

func NewOutPort(num int64) (OutPort, error) {
	log.Printf("Out Port %d created", num)
	return &port{num: num}, nil
}
