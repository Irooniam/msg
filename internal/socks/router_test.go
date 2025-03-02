package socks_test

import (
	"sync"
	"testing"

	"github.com/Irooniam/msg/internal/socks"
)

func TestNewRouterHappyPath(t *testing.T) {
	var wg sync.WaitGroup
	r, err := socks.NewZRouter("router-test")
	if err != nil {
		t.Error(err)
	}

	if err := r.Bind("127.0.0.1:9888"); err != nil {
		t.Error(err)
	}

	go r.Run()
	wg.Add(1)

	r.In <- [][]byte{[]byte("dealer"), []byte("payload")}
	r.Done <- true
	wg.Done()
}
