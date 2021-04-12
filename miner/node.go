package miner

import (
	"context"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/miner/proto"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"runtime/debug"
	"time"
)

var log = logging.Logger("miner")

func Run(ctx context.Context, node *core.IpfsNode) {
	smallNode := &SmallNode{
		node: node,
	}
	api, err := coreapi.NewCoreAPI(node, options.Api.FetchBlocks(true))
	if err != nil {
		log.Errorf("")
		return
	}
	smallNode.handler = NewV1Handler(api, smallNode)

	go smallNode.Run(ctx)
}

type SmallNode struct {
	node    *core.IpfsNode
	handler MessageHandler
}

func (n *SmallNode) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (n *SmallNode) Subscribe() error {
	topic, err := n.node.PubSub.Join(proto.V1Topic(n.node.Identity.String()))
	if err != nil {
		log.Errorf("failed to create sub topic: %v", err)
		return err
	}
	sub, err := topic.Subscribe()
	if err != nil {
		log.Errorf("failed to subscribe: %v", err)
		return err
	}
	log.Infof("subscribe: %v", n.node.Identity.String())

	go func() {
		for {
			pmsg, err := sub.Next(context.Background())
			if err != nil {
				log.Errorf("failed get message: %v", err)
				time.Sleep(time.Second)
				continue
			}
			log.Infof("received message from %v: %v", pmsg.ReceivedFrom.String(), pmsg.String())
			go func() {
				defer func() {
					if err := recover(); err != nil {
						log.Errorf("%v", string(debug.Stack()))
					}
				}()
				msg, err := proto.DecodeMessage(pmsg.Data)
				if err != nil {
					log.Errorf("failed to decode message: %v", err)
					return
				}
				err = n.handler.Handle(context.TODO(), pmsg.ReceivedFrom, &msg)
				if err != nil {
					log.Errorf("failed to handler message: %v", err)
				}
			}()
		}
	}()
	return nil
}

func (n *SmallNode) PublishMessage(ctx context.Context, topic string, msg *proto.Message) error {
	data, err := msg.EncodeMessage()
	if err != nil {
		log.Errorf("failed to encode message: %v", err)
		return err
	}
	receiverTopic, err := n.node.PubSub.Join(topic)
	if err != nil {
		log.Errorf("failed to create pub message: %v", err)
		return err
	}
	defer receiverTopic.Close()
	err = receiverTopic.Publish(ctx, data)
	if err != nil {
		log.Errorf("failed publish message: %v", err)
	}
	return nil
}