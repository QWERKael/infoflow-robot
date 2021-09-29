package receiver

type MessagePackage struct {
	EventType string  `json:"eventtype,omitempty"`
	AgentId   uint64  `json:"agentid,omitempty"`
	GroupId   uint64  `json:"groupid,omitempty"`
	CorpId    string  `json:"corpid,omitempty"`
	Message   Message `json:"message"`
	Time      int64   `json:"time,omitempty"`
}

type Message struct {
	Header Header    `json:"header"`
	Body   []BodyMsg `json:"body"`
}

type Header struct {
	FromUserId    string `json:"fromuserid,omitempty"`
	ToId          uint64 `json:"toid"`
	ToType        string `json:"totype,omitempty"`
	MsgType       string `json:"msgtype,omitempty"`
	ClientMsgId   uint64 `json:"clientmsgid,omitempty"`
	MessageId     uint64 `json:"messageid,omitempty"`
	MsgSeqId      string `json:"msgseqid,omitempty"`
	At            At     `json:"at,omitempty"`
	Compatible    string `json:"compatible,omitempty"`
	OfflineNotify string `json:"offlinenotify,omitempty"`
	Extra         string `json:"extra,omitempty"`
	ServerTime    uint64 `json:"servertime,omitempty"`
	ClientTime    uint64 `json:"clienttime,omitempty"`
	UpdateTime    uint64 `json:"updatetime,omitempty"`
}

type At struct {
	AtRobotIds []uint64 `json:"atrobotids,omitempty"`
	AtUserIds  []string `json:"atuserids,omitempty"`
}

type BodyMsg struct {
	Type        string `json:"type"`
	CommandName string `json:"commandname,omitempty"`
	Content     string `json:"content,omitempty"`
	RobotId     uint64 `json:"robotid,omitempty"`
	Name        string `json:"name,omitempty"`
	Label       string `json:"label,omitempty"`
	DownloadUrl string `json:"downloadurl,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
