package busybee

func NewStaticServiceDiscovery() StaticServiceDiscovery {
	s := StaticServiceDiscovery{}
	s.Services = make(map[Service]string)
	return s
}

type StaticServiceDiscovery struct {
	Services map[Service]string
}

func (discov StaticServiceDiscovery) AddService(url string, svc Service) {
	discov.Services[svc] = url
}

func (discov StaticServiceDiscovery) Discover(svc Service) string {
	return discov.Services[svc]
}
