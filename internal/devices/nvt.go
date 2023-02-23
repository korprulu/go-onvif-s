package devices

type NVT struct {
	Address string
	Scopes  map[string]string
}

func NewNVT() *NVT {
	return &NVT{
		Scopes: make(map[string]string),
	}
}
