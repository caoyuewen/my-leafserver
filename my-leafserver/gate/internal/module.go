package internal

import (
	"github.com/name5566/leaf/gate"
	"my-leafserver/game"
	"my-leafserver/models"
	"my-leafserver/msg"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      models.Server.MaxConnNum,
		PendingWriteNum: models.PendingWriteNum,
		MaxMsgLen:       models.MaxMsgLen,
		WSAddr:          models.Server.WSAddr,
		HTTPTimeout:     models.HTTPTimeout,
		CertFile:        models.Server.CertFile,
		KeyFile:         models.Server.KeyFile,
		TCPAddr:         models.Server.TCPAddr,
		LenMsgLen:       models.LenMsgLen,
		LittleEndian:    models.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}
