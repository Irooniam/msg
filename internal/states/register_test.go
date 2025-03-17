package states_test

import (
	"log"
	"sync"
	"testing"

	"github.com/Irooniam/msg/internal/states"
)

func TestRegisterDealerHappyPath(t *testing.T) {
	out := make(chan [][]byte)
	var m sync.Map
	err := states.AddDealer(&m, []byte("test-dealer"), out)
	if err != nil {
		t.Errorf("didnt expect error when adding new dealer to map %s", err)
	}

	d, ok := m.Load("test-dealer")
	if !ok {
		t.Error("Expected DEALERS to have test-dealer but cant key test-dealer")
	}

	if d.(states.DealerInfo).ID != "test-dealer" {
		t.Errorf("expected ID of dealer to be test-dealer but got %s", d.(states.DealerInfo).ID)
	}

	log.Println(d.(states.DealerInfo))
}

func TestRemoveDealerHappyPath(t *testing.T) {
	out := make(chan [][]byte)
	var m sync.Map
	var i int = 0
	_ = states.AddDealer(&m, []byte("test-dealer"), out)

	m.Range(func(k, v any) bool {
		log.Println(k, v)
		i++
		return true
	})

	if i != 1 {
		t.Errorf("Expecting 1 key in map but got %d", i)
	}

	states.RemoveDealer(&m, []byte("test-dealer"), out)

	i = 0
	m.Range(func(k, v any) bool {
		log.Println(k, v)
		i++
		return true
	})

	if i != 0 {
		t.Errorf("Expected no dealers in map but got %d", i)
	}
}
