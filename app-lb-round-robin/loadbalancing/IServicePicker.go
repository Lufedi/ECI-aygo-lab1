package loadbalancing

type IServicePicker interface {
	GetService() string
}
