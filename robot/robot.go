package robot

import (
	"Baidu/infoflow-robot/robot/receiver"
	"Baidu/infoflow-robot/robot/sender"
)

type IRobot struct {
	Sender   sender.Sender
	Receiver receiver.Receiver
}
