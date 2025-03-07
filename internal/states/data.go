package states

import (
	"sync"
)

type DealerInfo struct {
	ID   string
	Host string
	Port int
}

// use sync.map so goroutine safe
var DEALERS sync.Map
