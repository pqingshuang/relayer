package main

import (
	"context"
	"fmt"
	"github.com/nbd-wtf/go-nostr/nip19"

	"github.com/nbd-wtf/go-nostr"
)

func main() {

	relay, err := nostr.RelayConnect(context.Background(), "ws://localhost:7447")
	//relay, err := nostr.RelayConnect(context.Background(), "wss://nostr.688.org")
	if err != nil {
		panic(err)
	}
	fmt.Println(context.Background(), relay)

	npub := "npub1ejxswthae3nkljavznmv66p9ahp4wmj4adux525htmsrff4qym9sz2t3tv"
	//npub := ""
	var filters nostr.Filters
	if _, v, err := nip19.Decode(npub); err == nil {
		pub := "cc8d072efdcc676fcbac14f6cd6825edc3576e55eb786a2a975ee034a6a026cb"
		fmt.Println(pub, v)
		filters = []nostr.Filter{{
			Kinds: []int{1},
			//Authors: []string{pub},
			Limit: 1,
		}}
	} else {
		panic(err)
	}

	ctx, _ := context.WithCancel(context.Background())
	sub := relay.Subscribe(ctx, filters)
	notice := relay.Notices()
	go func() {
		<-sub.EndOfStoredEvents
		// handle end of stored events (EOSE, see NIP-15)
	}()

	for ev := range sub.Events {
		// handle returned event.
		// channel will stay open until the ctx is cancelled (in this case, by calling cancel())

		fmt.Println(ev.ID)
	}
}
