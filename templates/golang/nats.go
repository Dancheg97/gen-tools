package golang

const NatsWrapperGo = `package nats

import (
	pb "%s/gen/pb/proto/v1"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Params struct {
	Url       string
	Subject   string
	OnMessage func(mes *pb.AddResponse)
	OnErr     func(err error)
}

type Wrapper struct {
	Conn      *nats.Conn
	sub       *nats.Subscription
	subject   string
	OnMessage func(mes *pb.AddResponse)
	OnErr     func(err error)
}

func Get(params *Params) (*Wrapper, error) {
	nc, err := nats.Connect(params.Url)
	if err != nil {
		return nil, err
	}
	return &Wrapper{
		Conn:      nc,
		subject:   params.Subject,
		OnMessage: params.OnMessage,
		OnErr:     params.OnErr,
	}, nil
}

func (p *Wrapper) Produce(msg *pb.AddResponse, handleErr ...func(err error)) {
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

func (w *Wrapper) Consume() error {
	sub, err := w.Conn.Subscribe(w.subject, func(msg *nats.Msg) {
		var pbMes pb.AddResponse
		err := proto.Unmarshal(msg.Data, &pbMes)
		if err != nil {
			w.OnErr(err)
			return
		}
		w.OnMessage(&pbMes)
	})
	w.sub = sub
	return err
}

func (w *Wrapper) Stop() error {
	err := w.sub.Unsubscribe()
	if err != nil {
		return err
	}
	err = w.sub.Drain()
	if err != nil {
		return err
	}
	return nil
}
`
