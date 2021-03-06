package publishers

import (
	"errors"

	"github.com/mchmarny/thingz-commons"
)

const (
	PUB_CONSOLE  = "stdout"
	PUB_INFLUXDB = "influxdb"
	PUB_KAFKA    = "kafka"
	PUB_WS       = "websocket"
)

// Publisher describes the metric publisher functionality
type Publisher interface {

	// Publish metric
	Publish(in <-chan *commons.Metric, err chan<- error)

	// Finalize tells the publisher to close used resources
	// and do any general cleanup it needs
	Finalize()
}

// GetPublisher makes you wish for some generics
func GetPublisher(src, pub, args string) (Publisher, error) {

	switch pub {
	case PUB_CONSOLE:
		return NewConsolePublisher()
	case PUB_INFLUXDB:
		return NewInfluxDBPublisher(args)
	case PUB_KAFKA:
		return NewKafkaPublisher(src, args)
	case PUB_WS:
		return NewWebsocketPublisher(args)
	default:
		return nil, errors.New("Invalid publishing target: " + pub)
	}

}
