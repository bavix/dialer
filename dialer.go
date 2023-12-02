package dialer

import (
	"context"
	"net"
	"time"
)

const (
	tcp = "tcp"
	utp = "udp"
)

func DialTCP(ctx context.Context, address string) (time.Duration, error) {
	return Dial(ctx, tcp, address)
}

func DialUDP(ctx context.Context, address string) (time.Duration, error) {
	return Dial(ctx, utp, address)
}

func Dial(
	ctx context.Context,
	network string,
	address string,
) (time.Duration, error) {
	var d net.Dialer

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	now := time.Now()

	conn, err := d.DialContext(ctx, network, address)
	if err != nil {
		return 0, err
	}

	if err := conn.Close(); err != nil {
		return 0, err
	}

	return time.Since(now), nil
}
