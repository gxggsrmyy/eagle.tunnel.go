package comm

import (
	"errors"

	"github.com/eaglexiang/go-tunnel"
)

// Handler ET子协议的handler接口
type Handler interface {
	Handle(req string, tunnel *tunnel.Tunnel) error // 处理业务
	Type() int                                      // ET子协议的类型
	Name() string                                   // ET子协议的名字
}

// Sender ET子协议的sender
type Sender interface {
	Send(e *NetArg) error //发送流程
	Type() int            // ET子协议的类型
	Name() string         // ET子协议的名字
}

// SubHandlers 子协议处理器
var SubHandlers map[string]Handler

// SubSenders 子协议发射器
var SubSenders map[int]Sender

// AddSubHandler 添加ET子协议handler
func AddSubHandler(h Handler) {
	if SubHandlers == nil {
		SubHandlers = make(map[string]Handler)
	}
	SubHandlers[h.Name()] = h
}

// AddSubSender 添加子协议Sender
func AddSubSender(s Sender) {
	if SubSenders == nil {
		SubSenders = make(map[int]Sender)
	}
	SubSenders[s.Type()] = s
}

// GetHandler 根据特征头获取相应Handler
func GetHandler(subReq string) (Handler, error) {
	h, ok := SubHandlers[subReq]
	if !ok {
		return nil, errors.New("handler not found for: " + subReq)
	}
	return h, nil
}
