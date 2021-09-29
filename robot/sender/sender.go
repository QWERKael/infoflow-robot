package sender

import (
	"github.com/QWERKael/utility-go/codec"
	"github.com/QWERKael/utility-go/net"
)

type Sender struct {
	GroupId  uint64
	RobotUrl string
}

func (s *Sender) SendMsg(msgList []BodyMsgInterface) ([]byte, error) {
	mp := MessagePackage{
		Message: Message{
			Header: Header{ToId: []uint64{s.GroupId}},
			Body:   msgList,
		}}
	b, err := codec.EncodeJson(mp)
	if err != nil {
		return nil, err
	}
	header := map[string]string{"Content-Type": "application/json"}
	var respByte []byte
	respByte, err = net.PostWithHeader(s.RobotUrl, b, header)
	if err != nil {
		return nil, err
	}
	return respByte, nil
}

func (s *Sender) SendTextMsg(text string) ([]byte, error) {
	return s.SendMsg([]BodyMsgInterface{NewTextMsg(text)})
}

func (s *Sender) SendLinkMsg(link string) ([]byte, error) {
	return s.SendMsg([]BodyMsgInterface{NewLinkMsg(link)})
}
