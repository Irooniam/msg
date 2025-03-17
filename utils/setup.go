package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Irooniam/msg/conf"
	"github.com/Irooniam/msg/internal/socks"
)

func ChkDirectoryConf() (string, int, error) {
	var host string
	var port int
	if host = os.Getenv(conf.MSG_DIR_HOST); host == "" {
		return "", 0, errors.New(fmt.Sprintf("env var for directory host %s is not set", conf.MSG_DIR_HOST))
	}

	if os.Getenv(conf.MSG_DIR_PORT) == "" {
		return "", 0, errors.New(fmt.Sprintf("env var for directory port %s is not set", conf.MSG_DIR_PORT))
	}

	port, err := strconv.Atoi(os.Getenv(conf.MSG_DIR_PORT))
	if err != nil {
		return "", 0, errors.New(fmt.Sprintf("tried to convert port string to int got %s", err))
	}

	return host, port, nil
}

func ChkWSRouterConf() (string, int, error) {
	var host string
	var port int
	if host = os.Getenv(conf.MSG_DEALER_HOST); host == "" {
		return "", 0, errors.New(fmt.Sprintf("env var for ws zmq router host %s is not set", conf.MSG_DIR_HOST))
	}

	if os.Getenv(conf.MSG_DIR_PORT) == "" {
		return "", 0, errors.New(fmt.Sprintf("env var for directory port %s is not set", conf.MSG_DIR_PORT))
	}

	port, err := strconv.Atoi(os.Getenv(conf.MSG_DIR_PORT))
	if err != nil {
		return "", 0, errors.New(fmt.Sprintf("tried to convert port string to int got %s", err))
	}

	return host, port, nil
}

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
