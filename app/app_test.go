package app

import (
	"testing"

	"github.com/Trother555/portctl/port"
)

func TestAppRead(t *testing.T) {
	pIn, _ := port.NewInPort(0)
	pOut, _ := port.NewOutPort(0)
	app := app{inPorts: []port.InPort{pIn}, outPorts: []port.OutPort{pOut}}

	t.Run("happy path", func(t *testing.T) {
		_, err := app.Read(0)
		if err != nil {
			t.Errorf("app.Read() returned error: %s", err)
		}
	})

	t.Run("error path", func(t *testing.T) {
		_, err := app.Read(1)
		if err != ErrPortDoesNotExist {
			t.Errorf("app.Read() must return ErrPortDoesNotExist, got: %s", err)
		}
	})
}

func TestAppWrite(t *testing.T) {
	pIn, _ := port.NewInPort(0)
	pOut, _ := port.NewOutPort(0)
	app := app{inPorts: []port.InPort{pIn}, outPorts: []port.OutPort{pOut}}

	t.Run("happy path", func(t *testing.T) {
		err := app.Write(0, 0, 0)
		if err != nil {
			t.Errorf("app.Write() returned error: %s", err)
		}
	})

	t.Run("error path", func(t *testing.T) {
		err := app.Write(1, 0, 0)
		if err != ErrPortDoesNotExist {
			t.Errorf("app.Write() must return ErrPortDoesNotExist, got: %s", err)
		}
	})
}
