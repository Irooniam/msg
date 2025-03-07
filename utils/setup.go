package utils

import (
	"github.com/Irooniam/msg/internal/socks"
)

func SetupSocks(routerID string, dealerID string) (socks.ZDealer, socks.ZRouter, error) {
	router, err := socks.NewZRouter(routerID)
	if err != nil {
		return socks.ZDealer{}, socks.ZRouter{}, err
	}
	if err := router.Bind("127.0.0.1:9988"); err != nil {
		return socks.ZDealer{}, socks.ZRouter{}, err
	}

	dealer, err := socks.NewDealer(dealerID)
	if err != nil {
		return socks.ZDealer{}, socks.ZRouter{}, err
	}

	if err := dealer.Connect("127.0.0.1:9988"); err != nil {
		return socks.ZDealer{}, socks.ZRouter{}, err
	}

	return *dealer, *router, nil
}
