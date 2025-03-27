package directory_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Irooniam/msg/services/directory"
)

func TestAddDealerHappyPath(t *testing.T) {
	var wg sync.WaitGroup
	dconf := directory.DirService{}
	dconf.Host = "127.0.0.1"
	dconf.Port = 9888
	dconf.ID = "directory1"
	dconf.Endpoint = fmt.Sprintf("tcp://%s:%d", dconf.Host, dconf.Port)

	dir, err := directory.New(dconf)
	if err != nil {
		t.Fatal(err)
	}

	go dir.RouterRun()
	go dir.RecvMsg()
	wg.Add(2)

	action := directory.ACTIONS["REGISTER-DEALER"]
	dir.RouterIn() <- [][]byte{[]byte("magical-dealer"), []byte(action), []byte(time.Now().String())}

	dealer, ok := dir.Dealers.Load("magical-dealer")
	if !ok {
		t.Fatal("expected to find key magical-dealer but key not found")
	}

	if dealer.(directory.ServiceInfo).ID != "magical-dealer" {
		t.Fatalf("expecting dealer ID magical-dealer but got %s", dealer.(directory.ServiceInfo).ID)
	}

	wg.Done()
}

func TestRemoveDealerHappyPath(t *testing.T) {
	var wg sync.WaitGroup
	dconf := directory.DirService{}
	dconf.Host = "127.0.0.1"
	dconf.Port = 9888
	dconf.ID = "directory1"
	dconf.Endpoint = fmt.Sprintf("tcp://%s:%d", dconf.Host, dconf.Port)

	dir, err := directory.New(dconf)
	if err != nil {
		t.Fatal(err)
	}

	go dir.RouterRun()
	go dir.RecvMsg()
	wg.Add(2)

	action := directory.ACTIONS["REGISTER-DEALER"]
	dir.RouterIn() <- [][]byte{[]byte("magical-dealer"), []byte(action), []byte(time.Now().String())}

	dealer, ok := dir.Dealers.Load("magical-dealer")
	if !ok {
		t.Fatal("expected to find key magical-dealer but key not found")
	}

	if dealer.(directory.ServiceInfo).ID != "magical-dealer" {
		t.Fatalf("expecting dealer ID magical-dealer but got %s", dealer.(directory.ServiceInfo).ID)
	}

	//now lets remove dealer
	dir.RemoveDealer([]byte("magical-dealer"))

	if _, ok = dir.Dealers.Load("magical-dealer"); ok {
		t.Fatalf("didnt expect to find key in dealers map %t", ok)
	}

	wg.Done()
}
