// TODO: Collect metrics.

package npp

import (
	"context"
	"net"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"gopkg.in/oleiade/lane.v1"
)

// NATPuncher describes an interface of NAT Punching Protocol.
//
// It should be used to penetrate a NAT while connecting two peers located
// either under the same or different firewalls with network address
// translation enabled.
type NATPuncher interface {
	// Dial dials the given address.
	// Should be used only on client side.
	Dial(addr common.Address) (net.Conn, error)
	// DialContext connects to the address using the provided context.
	// Should be used only on client side.
	DialContext(ctx context.Context, addr common.Address) (net.Conn, error)
	// Accept blocks the current execution context until a new connection
	// arrives.
	//
	// Indented to be used on server side.
	Accept() (net.Conn, error)
	// @antmat said that this method is clearly self-descriptive and much obvious. Wow.
	AcceptContext(ctx context.Context) (net.Conn, error)
	// RendezvousAddr returns rendezvous remote address.
	RemoteAddr() net.Addr
	// Close closes the puncher.
	// Any blocked operations will be unblocked and return errors.
	Close() error
}

type serverConnectionWatcherDeprecated struct {
	Queue *lane.Queue
	Log   *zap.Logger
}

func (m *serverConnectionWatcherDeprecated) OnMoreConnections(conn net.Conn) {
	m.Log.Debug("enqueueing pending connection")
	m.Queue.Enqueue(conn)
}

type clientConnectionWatcherDeprecated struct{}

func (clientConnectionWatcherDeprecated) OnMoreConnections(conn net.Conn) {
	conn.Close()
}
