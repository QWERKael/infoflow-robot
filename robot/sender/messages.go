package sender

import "encoding/base64"

type BodyMsgType string

const (
	TEXT    BodyMsgType = "TEXT"
	LINK    BodyMsgType = "LINK"
	AT      BodyMsgType = "AT"
	Command BodyMsgType = "command"
	IMAGE   BodyMsgType = "IMAGE"
	MD      BodyMsgType = "MD"
)

type BodyMsgInterface interface {
	ShowType() string
}

type BodyMsg struct {
	Type BodyMsgType `json:"type"`
}

func (bm *BodyMsg) ShowType() string {
	return string(bm.Type)
}

type TextMsg struct {
	Content string `json:"content"`
	BodyMsg
}

func NewTextMsg(text string) *TextMsg {
	return &TextMsg{
		Content: text,
		BodyMsg: BodyMsg{Type: TEXT},
	}
}

type LinkMsg struct {
	Href string `json:"href"`
	BodyMsg
}

func NewLinkMsg(href string) *LinkMsg {
	return &LinkMsg{
		Href:    href,
		BodyMsg: BodyMsg{Type: LINK},
	}
}

type AtMsg struct {
	AtUserIds []string `json:"atuserids"`
	AtAll     bool     `json:"atall"`
	BodyMsg
}

func NewAtMsg(atUserIds []string, atAll bool) *AtMsg {
	return &AtMsg{
		AtUserIds: atUserIds,
		AtAll:     atAll,
		BodyMsg:   BodyMsg{Type: AT},
	}
}

type CommandMsg struct {
	CommandName string `json:"commandname"`
	BodyMsg
}

func NewCommandMsg(commandName string) *CommandMsg {
	return &CommandMsg{
		CommandName: commandName,
		BodyMsg:     BodyMsg{Type: Command},
	}
}

type ImageMsg TextMsg

func NewImageMsg(imageBytes []byte) *ImageMsg {
	s := base64.StdEncoding.EncodeToString(imageBytes)
	return &ImageMsg{
		Content: s,
		BodyMsg: BodyMsg{Type: IMAGE},
	}
}

type MdMsg TextMsg

func NewMdMsg(md string) *MdMsg {
	return &MdMsg{
		Content: md,
		BodyMsg: BodyMsg{Type: MD},
	}
}

type MessagePackage struct {
	Message Message `json:"message"`
}

type Message struct {
	Header Header             `json:"header"`
	Body   []BodyMsgInterface `json:"body"`
}

type Header struct {
	ToId []uint64 `json:"toid"`
}

