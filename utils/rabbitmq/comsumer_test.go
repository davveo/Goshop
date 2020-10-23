package rabbitmq

import (
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	m, err := New("amqp://zwf:123456@127.0.0.1:5672/").Open()
	if err != nil {
		t.Logf("[ERROR] %s\n", err.Error())
		return
	}
	defer m.Close()

	c, err := m.Consumer("test-consume")
	if err != nil {
		t.Logf("[ERROR] Create consumer failed, %v\n", err)
		return
	}
	defer c.Close()

	exb := []*ExchangeBinds{
		&ExchangeBinds{
			Exch: DefaultExchange("exch.unitest", ExchangeDirect),
			Bindings: []*Binding{
				&Binding{
					RouteKey: "route.unitest1",
					Queues: []*Queue{
						DefaultQueue("queue.unitest1"),
					},
				},
				&Binding{
					RouteKey: "route.unitest2",
					Queues: []*Queue{
						DefaultQueue("queue.unitest2"),
					},
				},
			},
		},
	}
	msgC := make(chan Delivery, 1)
	defer close(msgC)

	c.SetExchangeBinds(exb)
	c.SetMsgCallback(msgC)
	c.SetQos(10)
	if err = c.Open(); err != nil {
		t.Logf("[ERROR] Open failed, %v\n", err)
		return
	}

	for msg := range msgC {
		t.Logf("Tag(%d) Body: %s\n", msg.DeliveryTag, string(msg.Body))
		_ = msg.Ack(true)
		//if i%5 == 0 {
		//	c.CloseChan()
		//}
		//log.Info("Consumer receive msg `%s`", string(msg))
		time.Sleep(time.Second)
	}
}
