package states_test

import (
	"log"
	"testing"

	"github.com/Irooniam/msg/internal/states"
)

func TestRegisterDealerHappyPath(t *testing.T) {
	out := make(chan [][]byte)
	err := states.AddDealer([]byte("test-dealer"), out)
	if err != nil {
		t.Errorf("didnt expect error when adding new dealer to map %s", err)
	}

	d, ok := states.DEALERS.Load("test-dealer")
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
	var i int = 0
	_ = states.AddDealer([]byte("test-dealer"), out)

	states.DEALERS.Range(func(k, v any) bool {
		log.Println(k, v)
		i++
		return true
	})

	if i != 1 {
		t.Errorf("Expecting 1 key in map but got %d", i)
	}

	states.RemoveDealer([]byte("test-dealer"), out)

	i = 0
	states.DEALERS.Range(func(k, v any) bool {
		log.Println(k, v)
		i++
		return true
	})

	if i != 0 {
		t.Errorf("Expected no dealers in map but got %d", i)
	}
}
