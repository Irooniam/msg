package socks_test

import (
	"testing"

	"github.com/Irooniam/msg/internal/socks"
)

func TestNewRouterHappyPath(t *testing.T) {
	r, err := socks.NewZRouter("router-test")
	if err != nil {
		t.Error(err)
	}

	if err := r.Bind("127.0.0.1", 9876); err != nil {
		t.Error(err)
	}
}
