package types

type TransportationEventType string

const (
	EventTypeTransportationInit   TransportationEventType = "TransportationInit"
	EventTypeTransportationFinish                         = "TransportationFinish"
	EventTypeTransportationAccept                         = "TransportationAccept"
)
