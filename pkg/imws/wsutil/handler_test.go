package imwsutil

import (
	"bytes"
	"runtime"
	"testing"

	"github.com/luantao/IM-base/pkg/imws"
)

func TestControlHandler(t *testing.T) {
	for _, test := range []struct {
		name  string
		state imws.State
		in    imws.Frame
		out   imws.Frame
		noOut bool
		err   error
	}{
		{
			name: "ping",
			in:   imws.NewPingFrame(nil),
			out:  imws.NewPongFrame(nil),
		},
		{
			name: "ping",
			in:   imws.NewPingFrame([]byte("catch the ball")),
			out:  imws.NewPongFrame([]byte("catch the ball")),
		},
		{
			name:  "ping",
			state: imws.StateServerSide,
			in:    imws.MaskFrame(imws.NewPingFrame([]byte("catch the ball"))),
			out:   imws.NewPongFrame([]byte("catch the ball")),
		},
		{
			name: "ping",
			in:   imws.NewPingFrame(bytes.Repeat([]byte{0xfe}, 125)),
			out:  imws.NewPongFrame(bytes.Repeat([]byte{0xfe}, 125)),
		},
		{
			name:  "pong",
			in:    imws.NewPongFrame(nil),
			noOut: true,
		},
		{
			name:  "pong",
			in:    imws.NewPongFrame([]byte("catched")),
			noOut: true,
		},
		{
			name: "close",
			in:   imws.NewCloseFrame(nil),
			out:  imws.NewCloseFrame(nil),
			err: ClosedError{
				Code: imws.StatusNoStatusRcvd,
			},
		},
		{
			name: "close",
			in: imws.NewCloseFrame(imws.NewCloseFrameBody(
				imws.StatusGoingAway, "goodbye!",
			)),
			out: imws.NewCloseFrame(imws.NewCloseFrameBody(
				imws.StatusGoingAway, "",
			)),
			err: ClosedError{
				Code:   imws.StatusGoingAway,
				Reason: "goodbye!",
			},
		},
		{
			name: "close",
			in: imws.NewCloseFrame(imws.NewCloseFrameBody(
				imws.StatusGoingAway, "bye",
			)),
			out: imws.NewCloseFrame(imws.NewCloseFrameBody(
				imws.StatusGoingAway, "",
			)),
			err: ClosedError{
				Code:   imws.StatusGoingAway,
				Reason: "bye",
			},
		},
		{
			name:  "close",
			state: imws.StateServerSide,
			in: imws.MaskFrame(imws.NewCloseFrame(imws.NewCloseFrameBody(
				imws.StatusGoingAway, "goodbye!",
			))),
			out: imws.NewCloseFrame(imws.NewCloseFrameBody(
				imws.StatusGoingAway, "",
			)),
			err: ClosedError{
				Code:   imws.StatusGoingAway,
				Reason: "goodbye!",
			},
		},
		{
			name: "close",
			in: imws.NewCloseFrame(imws.NewCloseFrameBody(
				imws.StatusNormalClosure, string([]byte{0, 200}),
			)),
			out: imws.NewCloseFrame(imws.NewCloseFrameBody(
				imws.StatusProtocolError, imws.ErrProtocolInvalidUTF8.Error(),
			)),
			err: imws.ErrProtocolInvalidUTF8,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					stack := make([]byte, 4096)
					n := runtime.Stack(stack, true)
					t.Fatalf(
						"panic recovered: %v\n%s",
						err, stack[:n],
					)
				}
			}()
			var (
				out = bytes.NewBuffer(nil)
				in  = bytes.NewReader(test.in.Payload)
			)
			c := ControlHandler{
				Src:   in,
				Dst:   out,
				State: test.state,
			}

			err := c.Handle(test.in.Header)
			if err != test.err {
				t.Errorf("unexpected error: %v; want %v", err, test.err)
			}

			if in.Len() != 0 {
				t.Errorf("handler did not drained the input")
			}

			act := out.Bytes()
			switch {
			case len(act) == 0 && test.noOut:
				return
			case len(act) == 0 && !test.noOut:
				t.Errorf("unexpected silence")
			case len(act) > 0 && test.noOut:
				t.Errorf("unexpected sent frame")
			default:
				exp := imws.MustCompileFrame(test.out)
				if !bytes.Equal(act, exp) {
					fa := imws.MustReadFrame(bytes.NewReader(act))
					fe := imws.MustReadFrame(bytes.NewReader(exp))
					t.Errorf(
						"unexpected sent frame:\n\tact: %+v\n\texp: %+v\nbytes:\n\tact: %v\n\texp: %v",
						fa, fe, act, exp,
					)
				}
			}
		})
	}
}
