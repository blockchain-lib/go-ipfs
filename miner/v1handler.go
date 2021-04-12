package miner

import (
	"context"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/miner/proto"
	"github.com/libp2p/go-libp2p-core/peer"
)

type V1Handler struct {
	node        *core.IpfsNode
	handlerFunc map[string]HandleFunc
	publisher   MessagePublisher
}

func NewV1Handler(node *core.IpfsNode, publisher MessagePublisher) *V1Handler {
	h := &V1Handler{
		node:        node,
		handlerFunc: make(map[string]HandleFunc),
		publisher:   publisher,
	}
	h.handlerFunc[proto.MsgAddFile] = h.HandleFetchFile
	return h
}

func (h *V1Handler) Handle(ctx context.Context, receivedFrom peer.ID, msg *proto.Message) error {
	if f, ok := h.handlerFunc[msg.Type]; ok {
		return f(ctx, receivedFrom, msg)
	}
	log.Warnf("message type not register: %v", msg.Type)
	return nil
}

func (h *V1Handler) HandleFetchFile(ctx context.Context, receivedFrom peer.ID, msg *proto.Message) error {
	return nil
}
