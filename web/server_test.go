package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type appMock struct {
	readMock  func(portNum int64) (int64, error)
	writeMock func(portNum int64, transactionId int64, val int64) error
}

func (a *appMock) Read(portNum int64) (int64, error) {
	return a.readMock(portNum)
}

func (a *appMock) Write(portNum int64, transactionId int64, val int64) error {
	return a.writeMock(portNum, transactionId, val)
}

func TestServerRead(t *testing.T) {
	app := &appMock{
		readMock:  func(portNum int64) (int64, error) { return 1, nil },
		writeMock: func(portNum, transactionId, val int64) error { return nil },
	}
	s := Server{app: app}

	t.Run("happy path", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/read?portNum=0", nil)
		w := httptest.NewRecorder()
		s.ReadHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("/read expected status 200, got: %s", res.Status)
		}

		dataRaw, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("io.ReadAll failed: %s", err)
		}

		data := map[string]string{}
		err = json.Unmarshal(dataRaw, &data)
		if err != nil {
			t.Errorf("json.Unmarshal failed: %s", err)
		}

		if val := data["val"]; val != "1" {
			t.Errorf("server must return 1")
		}
	})

	t.Run("missing param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/read", nil)
		w := httptest.NewRecorder()
		s.ReadHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("/read expected status 400, got: %s", res.Status)
		}
	})

	t.Run("read failed", func(t *testing.T) {

		app := &appMock{
			readMock:  func(portNum int64) (int64, error) { return 1, errors.New("failed") },
			writeMock: func(portNum, transactionId, val int64) error { return nil },
		}
		s := Server{app: app}

		req := httptest.NewRequest(http.MethodGet, "/read?portNum=1", nil)
		w := httptest.NewRecorder()
		s.ReadHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("/read expected status 500, got: %s", res.Status)
		}
	})
}

func TestServerWrite(t *testing.T) {
	var writtenVal, writtenTransactionId int64
	app := &appMock{
		readMock: func(portNum int64) (int64, error) { return 1, nil },
		writeMock: func(portNum, transactionId, val int64) error {
			writtenVal = val
			writtenTransactionId = transactionId
			return nil
		},
	}
	s := Server{app: app}

	t.Run("happy path", func(t *testing.T) {
		var expectedVal int64 = 111
		var expectedTransactionId int64 = 222
		req := httptest.NewRequest(
			http.MethodPost,
			fmt.Sprintf("/write?portNum=0&transactionId=%d&val=%d", expectedTransactionId, expectedVal),
			nil,
		)
		w := httptest.NewRecorder()
		s.WriteHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Errorf("/read expected status 200, got: %s", res.Status)
		}
		if writtenVal != expectedVal {
			t.Errorf("expected val to be equal %d, got %d", expectedVal, writtenVal)
		}
		if writtenTransactionId != expectedTransactionId {
			t.Errorf("expected transactionId to be equal %d, got %d", expectedTransactionId, writtenTransactionId)
		}
	})

	t.Run("missing param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/write", nil)
		w := httptest.NewRecorder()
		s.WriteHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("/write expected status 400, got: %s", res.Status)
		}
	})

	t.Run("write failed", func(t *testing.T) {
		app := &appMock{
			readMock:  func(portNum int64) (int64, error) { return 1, nil },
			writeMock: func(portNum, transactionId, val int64) error { return errors.New("failed") },
		}
		s := Server{app: app}

		req := httptest.NewRequest(http.MethodPost, "/write?portNum=0&transactionId=1&val=1", nil)
		w := httptest.NewRecorder()
		s.WriteHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("/write expected status 500, got: %s", res.Status)
		}
	})
}
