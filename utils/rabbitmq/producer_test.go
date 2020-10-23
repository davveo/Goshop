package rabbitmq

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	m, err := New("amqp://zwf:123456@127.0.0.1:5672/").Open()
	if err != nil {
		t.Logf("[ERROR] %s\n", err.Error())
		return
	}
	defer m.Close()

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
	for i := 0; i < 10; i++ {
		go func(idx int) {
			p, err := m.Producer(strconv.Itoa(i))
			if err != nil {
				t.Logf("[ERROR] Create producer failed, %v\n", err)
				return
			}
			if err = p.SetExchangeBinds(exb).Confirm(true).Open(); err != nil {
				t.Logf("[ERROR] Open failed, %v\n", err)
				return
			}
			for j := 0; j < 10; j++ {
				go func(v int) {
					for {
						v++
						msg := NewPublishMsg([]byte(fmt.Sprintf(`{"name":"zwf-%d"}`, v)))
						err = p.Publish("exch.unitest", "route.unitest2", msg)
						if err != nil {
							t.Logf("[ERROR] %s\n", err.Error())
						}
						//log.Info("Producer(%d) state:%d, err:%v\n", i, p.State(), err)
					}
				}(j)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
}
