package main

import (
	"context"
	"log"
	"sync"
)

type hub struct {
	sync.Mutex
	subs map[*subscriber]struct{}
}

func newHub() *hub {
	return &hub{
		subs: map[*subscriber]struct{}{},
	}
}

type message struct {
	data []byte
}

type subscriber struct {
	sync.Mutex
	name    string
	handler chan *message
	quit    chan struct{}
}

func (s *subscriber) run(ctx context.Context) {
	for {
		select {
		case msg := <-s.handler:
			log.Println(s.name, string(msg.data))
		case <-s.quit:
			return
		case <-ctx.Done():
			return
		}
	}
}

func newSubscriber(name string) *subscriber {
	return &subscriber{
		name:    name,
		handler: make(chan *message, 100),
		quit:    make(chan struct{}),
	}
}

func (h *hub) subscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	h.subs[s] = struct{}{}
	h.Unlock()

	go s.run(ctx)

	return nil
}

func main() {

}
