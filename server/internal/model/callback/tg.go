package callback

type ImCallback struct {
	Type int    //类型
	Data []byte // data
}

type ReceiverCallback struct {
	Msg       string
	MsgId     int
	MsgFromId int64
	Out       bool
	PeerId    int64
}
