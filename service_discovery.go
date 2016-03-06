package busybee

type ServiceDiscovery interface {
	Discover(Service) string
}
