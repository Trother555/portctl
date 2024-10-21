package port

import "testing"

func TestPortRead(t *testing.T) {
	p := port{num: 123}

	t.Run("happy path", func(t *testing.T) {
		_, err := p.Read()
		if err != nil {
			t.Errorf("p.Read() returned error: %s", err)
		}
	})
}

func TestPortWrite(t *testing.T) {
	p := port{num: 123}
	t.Run("happy path", func(t *testing.T) {
		err := p.Write(123, 456)
		if err != nil {
			t.Errorf("p.Write() returned error: %s", err)
		}
	})
}
