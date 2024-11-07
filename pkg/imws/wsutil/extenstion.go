package imwsutil

import "MyIM/pkg/imws"

// RecvExtension is an interface for clearing fragment header RSV bits.
type RecvExtension interface {
	UnsetBits(imws.Header) (imws.Header, error)
}

// RecvExtensionFunc is an adapter to allow the use of ordinary functions as
// RecvExtension.
type RecvExtensionFunc func(imws.Header) (imws.Header, error)

// BitsRecv implements RecvExtension.
func (fn RecvExtensionFunc) UnsetBits(h imws.Header) (imws.Header, error) {
	return fn(h)
}

// SendExtension is an interface for setting fragment header RSV bits.
type SendExtension interface {
	SetBits(imws.Header) (imws.Header, error)
}

// SendExtensionFunc is an adapter to allow the use of ordinary functions as
// SendExtension.
type SendExtensionFunc func(imws.Header) (imws.Header, error)

// BitsSend implements SendExtension.
func (fn SendExtensionFunc) SetBits(h imws.Header) (imws.Header, error) {
	return fn(h)
}
