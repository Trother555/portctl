package app

import (
	"errors"
	"log"

	"github.com/Trother555/portctl/port"
)

type App interface {
	Read(portNum int64) (int64, error)
	Write(portNum int64, transactionId int64, val int64) error
}

type app struct {
	inPorts  []port.InPort
	outPorts []port.OutPort
}

type Config struct {
	InPorts  int64
	OutPorts int64
}

var ErrPortDoesNotExist = errors.New("port does not exist")

func New(cfg *Config) App {
	inPorts := make([]port.InPort, 0, cfg.InPorts)
	var i int64
	for i = 0; i < cfg.InPorts; i++ {
		p, err := port.NewInPort(i)
		if err != nil {
			log.Fatalf("App: failed to init port: %d", i)
		}
		inPorts = append(inPorts, p)
	}

	outPorts := make([]port.OutPort, 0, cfg.OutPorts)
	for i = 0; i < cfg.OutPorts; i++ {
		p, err := port.NewOutPort(i)
		if err != nil {
			log.Fatalf("App: failed to init port: %d", i)
		}
		outPorts = append(outPorts, p)
	}

	return &app{inPorts: inPorts, outPorts: outPorts}
}

func (a *app) Read(portNum int64) (int64, error) {
	if portNum >= int64(len(a.inPorts)) {
		log.Printf("error: port %d does not exist", portNum)
		return -1, ErrPortDoesNotExist
	}
	return a.inPorts[portNum].Read()
}

func (a *app) Write(portNum int64, transactionId int64, val int64) error {
	if portNum >= int64(len(a.outPorts)) {
		log.Printf("error: port %d does not exist", portNum)
		return ErrPortDoesNotExist
	}
	return a.outPorts[portNum].Write(transactionId, val)
}
