package base

import (
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/module"
	"my-leafserver/models"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              models.GoLen,
		TimerDispatcherLen: models.TimerDispatcherLen,
		AsynCallLen:        models.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(models.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
