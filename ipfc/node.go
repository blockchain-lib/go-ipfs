package ipfc

import (
	"context"
	"github.com/ipfs/go-ipfs/core"
	logging "github.com/ipfs/go-log/v2"
	"time"
)

const (
	version = "v1"
)

var log = logging.Logger("ipfc")

func Run(ctx context.Context, node *core.IpfsNode) {
	ln := LittleNode{
		node: node,
	}
	go ln.Run(ctx)
}

type LittleNode struct {
	node *core.IpfsNode
}

func (n *LittleNode) Run(ctx context.Context) {
	n.sub()

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (n *LittleNode) sub() error {
	topic, err := n.node.PubSub.Join(version + "/" + n.node.Identity.String())
	if err != nil {
		log.Errorf("failed to create sub topic: %v", err)
		return err
	}
	sub, err := topic.Subscribe()

	log.Infof("subscribe: %v", n.node.Identity.String())
	if err != nil {
		log.Errorf("failed to subscribe: %v", err)
		return err
	}
	go func() {
		for {
			msg, err := sub.Next(context.Background())
			if err != nil {
				log.Errorf("failed get message: %v", err)
				time.Sleep(time.Second)
				continue
			}
			log.Infof("received message from %v: %v", msg.ReceivedFrom.String(), msg.String())
			receiverTopic, err :=  n.node.PubSub.Join(version + "/" + msg.ReceivedFrom.String())
			if err != nil{
				continue
			}
			receiverTopic.Publish(context.Background(), msg.Data)
			receiverTopic.Close()
		}
	}()
	return nil
}
