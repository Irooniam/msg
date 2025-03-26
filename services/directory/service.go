package directory

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/Irooniam/msg/conf"
	"github.com/Irooniam/msg/internal/socks"
)

type DirService struct {
	ID       string
	Host     string
	Port     int
	Endpoint string
	Dealers  *sync.Map
	router   *socks.ZRouter
}

/*
*
this struct will not be shared among services
as that will use proto
*
*/
type ServiceInfo struct {
	ID         string
	RouterHost string
	RouterPort int
	Endpoint   string
}

func (d *DirService) RemoveDealer(ID []byte) {
	d.Dealers.Delete(string(ID))
}

func (d *DirService) AddDealer(info ServiceInfo) error {
	d.Dealers.Store(info.ID, info)
	return nil
}

/*
*
wrap router.Run since router is private
in DirService struct
*
*/
func (d *DirService) RouterRun() {
	d.router.Run()
}

func (d *DirService) DealerEvent(ID []byte) {
	_, ok := d.Dealers.Load(string(ID))

	/**
	If we dont have this dealer - dont do anything because
	once dealer sends registration they will be added to map
	**/
	if !ok {
		log.Println("dont have dealer so dont do anything ", string(ID))
		return
	}

	//dealers exists so remove them from DEALER map
	log.Println("already have dealer so remove ", string(ID))
	d.RemoveDealer(ID)
}

func (d *DirService) RecvMsg() {
	for {
		msg := <-d.router.In
		log.Println("Directory got message ", msg)
	}
}

func ChkDirServiceConf() (DirService, error) {
	dir := DirService{}
	var err error

	if dir.ID = os.Getenv(conf.MSG_DIR_ID); dir.ID == "" {
		return dir, errors.New(fmt.Sprintf("env var for directory id %s is not set", conf.MSG_DIR_ID))
	}

	if dir.Host = os.Getenv(conf.MSG_DIR_HOST); dir.Host == "" {
		return dir, errors.New(fmt.Sprintf("env var for directory host %s is not set", conf.MSG_DIR_HOST))
	}

	if os.Getenv(conf.MSG_DIR_PORT) == "" {
		return dir, errors.New(fmt.Sprintf("env var for directory port %s is not set", conf.MSG_DIR_PORT))
	}

	if dir.Port, err = strconv.Atoi(os.Getenv(conf.MSG_DIR_PORT)); err != nil {
		return dir, errors.New(fmt.Sprintf("tried to convert port string to int got %s", err))
	}

	dir.Endpoint = fmt.Sprintf("tcp://%s:%d", dir.Host, dir.Port)
	return dir, nil
}

func New() (*DirService, error) {

	dirConf, err := ChkDirServiceConf()
	if err != nil {
		return &DirService{}, err
	}

	router, err := socks.NewZRouter(dirConf.ID)
	if err != nil {
		return &DirService{}, err
	}

	if err := router.Bind(dirConf.Endpoint); err != nil {
		return &DirService{}, err
	}

	dir := DirService{}
	dir.router = router
	return &dir, nil
}
