package golang

const NatsProducerGo = `package nats

import (
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Params struct {
	Url     string
	Subject string
}

type Producer struct {
	Conn    *nats.Conn
	subject string
}

func Get(params *Params) (*Producer, error) {
	nc, err := nats.Connect(params.Url)
	if err != nil {
		return nil, err
	}
	return &Producer{
		Conn:    nc,
		subject: params.Subject,
	}, nil
}

func (p *Producer) Produce(msg *pb.Message, handleErr ...func(err error)) {
	bytes, err := proto.Marshal(msg)
	processErr(err)
	err = p.Conn.Publish(p.subject, bytes)
	processErr(err)
}

func processErr(err error, onErr ...func(err error)) {
	if err != nil {
		for _, errFunc := range onErr {
			errFunc(err)
		}
	}
}
`

const NatsConsumerGo = `package nats

import (
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Params struct {
	Url       string
	Subject   string
	OnMessage func(mes *pb.Message)
	OnErr     func(err error)
}

type Consumer struct {
	conn *nats.Conn
	sub  *nats.Subscription
}

func Get(params *Params) (*Consumer, error) {
	nc, err := nats.Connect(params.Url)
	if err != nil {
		return nil, err
	}
	sub, err := nc.Subscribe(params.Subject, func(msg *nats.Msg) {
		var pbMes pb.Message
		err := proto.Unmarshal(msg.Data, &pbMes)
		if err != nil {
			params.OnErr(err)
			return
		}
		params.OnMessage(&pbMes)
	})
	if err != nil {
		return nil, err
	}
	return &Consumer{
		conn: nc,
		sub:  sub,
	}, err
}

func (c *Consumer) Stop() error {
	err := c.sub.Unsubscribe()
	if err != nil {
		return err
	}
	err = c.sub.Drain()
	if err != nil {
		return err
	}
	return nil
}
`
