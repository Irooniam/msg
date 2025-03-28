package directory

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/Irooniam/msg/conf"
	"github.com/Irooniam/msg/internal/socks"
	pb "github.com/Irooniam/msg/protos"
	"google.golang.org/protobuf/proto"
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
	RouterID   string
	Endpoint   string
}

func (d *DirService) RemoveDealer(ID []byte) {
	d.Dealers.Delete(string(ID))
}

/*
*
AddDealer stores the serviceinfo struct into map.sync
It also has to broadcast current snapshot of all connected services
*
*/
func (d *DirService) AddDealer(info ServiceInfo) error {
	d.Dealers.Store(info.ID, info)

	services := make([]*pb.Service, 0)
	dealers := []string{}
	//@todo for each dealer - send proto msg with dealer router conn info
	d.Dealers.Range(func(key, value interface{}) bool {
		service := &pb.Service{}
		v := value.(ServiceInfo)
		service.DealerId = v.ID
		service.RouterId = v.RouterID
		service.RouterHost = v.RouterHost
		service.RouterPort = int32(v.RouterPort)
		service.RouterEndpoint = v.Endpoint
		services = append(services, service)
		dealers = append(dealers, v.ID)
		fmt.Printf("Key: %s, Value: %v", key, value.(ServiceInfo))
		return true // Continue iterating
	})

	payload, err := proto.Marshal(&pb.BroadcastServices{Services: services})
	if err != nil {
		return err
	}

	for k := range dealers {
		d.router.Out <- [][]byte{[]byte(dealers[k]), []byte("DR"), payload}
	}

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
	log.Println(string(ID))
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

/*
*
Router zmq socket is not exposed directly
Only way to send / recv messages is via channels
*
*/
func (d *DirService) RouterIn() chan [][]byte {
	return d.router.In
}

func (d *DirService) RouterOut() chan [][]byte {
	return d.router.Out
}

func (d *DirService) RecvMsg() {
	for {
		msg := <-d.router.In

		//have to receive 3 parts: header, action, payload
		if len(msg) != 3 {
			log.Printf("expected msg received to be 3 parts but is: %d", len(msg))
			continue
		}

		/*
			determine if this is dealer dis / connect msg
			slices are pointers so cant do straight compare
		*/
		if bytes.Equal(msg[2], []byte{0}) {
			msg[1] = []byte(ACTIONS["DEALER-EVENT"])
		}

		action, err := TranslateAction(msg[1])
		if err != nil {
			log.Printf("tried getting action from msg but got %s", err)
			continue
		}

		//all messages go / recv in []byte
		saction := string(action)

		/*
		   We are matching what the action message TRANSLATES to not actual msg (e.g. DR)
		*/
		switch saction {
		case "DEALER-EVENT": //connect/disconnect
			log.Println("dealer event ", msg)
			d.DealerEvent(msg[0])

		case "REGISTER-DEALER": //new dealer sent msg
			log.Println("we have new dealer with ID", string(msg[0]))
			sinfo := ServiceInfo{}
			sinfo.ID = string(msg[0])
			err := d.AddDealer(sinfo)
			if err != nil {
				log.Printf("Error registering dealer: %s", err)
			}

		default:
			log.Printf("actions is %s - and we dont have case math", action)

		}

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

func New(dirConf DirService) (*DirService, error) {
	router, err := socks.NewZRouter(dirConf.ID)
	if err != nil {
		return &DirService{}, err
	}

	if err := router.Bind(dirConf.Endpoint); err != nil {
		return &DirService{}, err
	}

	dir := DirService{
		ID:       dirConf.ID,
		Host:     dirConf.Host,
		Port:     dirConf.Port,
		Endpoint: dirConf.Endpoint,
		Dealers:  &sync.Map{},
		router:   router,
	}
	return &dir, nil
}
