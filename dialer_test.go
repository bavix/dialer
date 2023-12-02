package dialer_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/bavix/dialer"
)

type DialTestSuite struct {
	suite.Suite

	svr  *httptest.Server
	addr string
}

func TestDialTestSuite(t *testing.T) {
	suite.Run(t, new(DialTestSuite))
}

func (s *DialTestSuite) SetupSuite() {
	s.svr = httptest.NewServer(nil)
	s.addr = s.svr.URL[7:]
}

func (s *DialTestSuite) TearDownSuite() {
	s.svr.Close()
}

func (s *DialTestSuite) TestSuccess() {
	t := s.T()

	t.Run("success [tcp]", func(t *testing.T) {
		elapsed, err := dialer.DialTCP(context.Background(), s.addr)

		require.NoError(t, err)
		require.NotZero(t, elapsed)

		t.Logf("elapsed: %v", elapsed)
	})

	t.Run("success [udp]", func(t *testing.T) {
		elapsed, err := dialer.DialUDP(context.Background(), s.addr)

		require.NoError(t, err)
		require.NotZero(t, elapsed)

		t.Logf("elapsed: %v", elapsed)
	})
}

func (s *DialTestSuite) TestHost() {
	t := s.T()

	t.Run("host [tcp]", func(t *testing.T) {
		elapsed, err := dialer.DialTCP(context.Background(), "github.com:80")

		require.NoError(t, err)
		require.NotZero(t, elapsed)

		t.Logf("elapsed: %v", elapsed)
	})

	t.Run("host [udp]", func(t *testing.T) {
		elapsed, err := dialer.DialUDP(context.Background(), "github.com:80")

		require.NoError(t, err)
		require.NotZero(t, elapsed)

		t.Logf("elapsed: %v", elapsed)
	})

	t.Run("host ssl [tcp]", func(t *testing.T) {
		elapsed, err := dialer.DialTCP(context.Background(), "github.com:443")

		require.NoError(t, err)
		require.NotZero(t, elapsed)

		t.Logf("elapsed: %v", elapsed)
	})

	t.Run("host ssl [udp]", func(t *testing.T) {
		elapsed, err := dialer.DialUDP(context.Background(), "github.com:443")

		require.NoError(t, err)
		require.NotZero(t, elapsed)

		t.Logf("elapsed: %v", elapsed)
	})
}

func (s *DialTestSuite) TestDeadline() {
	t := s.T()

	t.Run("deadline [tcp]", func(t *testing.T) {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now())
		defer cancel()

		elapsed, err := dialer.DialUDP(ctx, s.addr)

		require.Error(t, err)
		require.Error(t, ctx.Err())
		require.Zero(t, elapsed)

		t.Logf("elapsed: %v", elapsed)
	})

	t.Run("deadline [udp]", func(t *testing.T) {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now())
		defer cancel()

		elapsed, err := dialer.DialUDP(ctx, s.addr)

		require.Error(t, err)
		require.Error(t, ctx.Err())
		require.Zero(t, elapsed)

		t.Logf("elapsed: %v", elapsed)
	})
}
