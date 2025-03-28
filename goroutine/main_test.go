package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func connectToService() interface{} {
	time.Sleep(time.Second)
	return struct{}{}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		connpool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		defer server.Close()

		for {

			conn, err := server.Accept()
			if err != nil {
				log.Printf("failed to accept: %v", err)
				continue
			}

			//connectToService()
			svcConn := connpool.Get()
			fmt.Fprintln(conn, "")
			connpool.Put(svcConn)
			conn.Close()
		}
	}()

	return &wg
}

func init() {
	daemon := startNetworkDaemon()
	daemon.Wait()
}

func BenchmarkNetworkRewuest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("failed to dial host: %v", err)
		}
		if _, err := io.ReadAll(conn); err != nil {
			b.Fatalf("failed to read from connection: %v", err)
		}
		conn.Close()
	}
}
