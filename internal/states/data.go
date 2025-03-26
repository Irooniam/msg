package states

type ServiceInfo struct {
	ID         string
	RouterHost string
	RouterPort int
	Endpoint   string
}

type RouterInfo struct {
	ID       string
	Host     string
	Port     int
	Endpoint string
}

type DealerInfo struct {
	ID string
}
