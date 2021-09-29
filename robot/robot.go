package robot

import (
	"github.com/QWERKael/infoflow-robot/robot/receiver"
	"github.com/QWERKael/infoflow-robot/robot/sender"
)

type IRobot struct {
	Sender   sender.Sender
	Receiver receiver.Receiver
}
