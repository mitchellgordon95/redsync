package redsync

import (
	"os"
	"testing"

	"github.com/stvp/tempredis"
)

var servers []*tempredis.Server

func TestMain(m *testing.M) {
	for i := 0; i < 8; i++ {
		server, err := tempredis.Start(tempredis.Config{})
		if err != nil {
			panic(err)
		}
		servers = append(servers, server)
	}
	result := m.Run()
	for _, server := range servers {
		server.Term()
	}
	os.Exit(result)
}

func TestRedsync(t *testing.T) {
	pools := newMockPools(8)
	rs := New(pools)

	m := rs.NewMutex("test-redsync")
	_ = m.Lock()

	assertAcquired(t, pools, m.(*mutex))
}
