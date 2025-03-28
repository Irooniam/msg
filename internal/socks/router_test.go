package socks_test

import (
	"sync"
	"testing"
	"time"

	"github.com/Irooniam/msg/internal/socks"
	"github.com/Irooniam/msg/internal/states"
	"github.com/Irooniam/msg/utils"
)

func TestRegisterDealerHappyPath(t *testing.T) {
	var wg sync.WaitGroup
	dealer, router, err := utils.SetupSocks("test-router", "test-dealer")
	if err != nil {
		t.Fatal(err)
	}

	go router.Run()
	go router.ParseIn()
	time.Sleep(time.Millisecond * 100)
	go dealer.Run()
	go dealer.ParseIn()
	wg.Add(4)

	err = states.RegisterDealer("test-dealer", dealer.Out)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 100)
	sd, ok := states.DEALERS.Load("test-dealer")
	if !ok {
		t.Fatal("could load key dealer ", err)
	}

	d := sd.(states.DealerInfo)
	if d.ID != "test-dealer" {
		t.Fatalf("Expecting dealer ID to be test-dealer but got %s", d.ID)
	}

	wg.Done()
}

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
