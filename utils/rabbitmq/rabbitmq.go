package rabbitmq

import "log"

func GetRabbitmq() *MQ {
	var (
		mq  *MQ
		err error
	)
	for i := 0; i < 3; i++ {
		mq, err = New("amqp://zwf:123456@127.0.0.1:5672/").Open()
		if err == nil {
			break
		}
	}

	if mq == nil {
		log.Panicf("[ERROR] %s\n", err.Error())
	}

	return mq
}
