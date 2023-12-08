package contract

type ILogEvent interface {
	Create(event, key string) error
}
