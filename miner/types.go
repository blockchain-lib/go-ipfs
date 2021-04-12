package miner

import (
	"context"
	"github.com/ipfs/go-ipfs/miner/proto"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type HandleFunc func(ctx context.Context, receivedFrom peer.ID, msg *proto.Message) error

type MessageHandler interface {
	Handle(ctx context.Context, receivedFrom peer.ID, msg *proto.Message) error
}

type MessagePublisher interface {
	PublishMessage(topic string, msg *pubsub.Message) error
}
