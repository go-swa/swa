package service_swa

import (
	"context"
	"go.uber.org/zap"
)

type TSystem interface {
	GetServerInfo(ctx context.Context) (server Server, err error)


}



func (imp *impl) GetServerInfo(ctx context.Context) (server Server, err error) {
	var s Server

	s.Os = InitOS()

	if s.Cpu, err = InitCPU(); err != nil {
		imp.Logger(ctx).Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return s, err
	}

	if s.Ram, err = InitRAM(); err != nil {
		imp.Logger(ctx).Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return s, err
	}

	if s.Disk, err = InitDisk(); err != nil {
		imp.Logger(ctx).Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return s, err
	}

	return s, nil
}
