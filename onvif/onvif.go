package onvif

import "time"

type Capabilities struct{}

type Onvifdev struct {
	Xaddr       string
	Model       string
	Seriel      string
	DeviceXaddr string
	EventsXaddr string
	MediaXaddr  string
	TimeOffset  int64
	Capabilities
}

func CheckXaddrsAndGetTime() time.Time {
}
