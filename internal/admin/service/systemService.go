/**
*@author:yangzhou
*@date: 2023/2/13
*@email: yangzhou2224@shengtian.com
*@description:
 */
package service

import (
	"yzgin/global"
	"yzgin/pkg/utils"

	"go.uber.org/zap"
)

func (s *Service) GetServerInfo() (srv *utils.Server, err error) {
	var server utils.Server
	server.Os = utils.InitOS()
	if server.Cpu, err = utils.InitCPU(); err != nil {
		global.Log.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &server, err
	}
	if server.Ram, err = utils.InitRAM(); err != nil {
		global.Log.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &server, err
	}
	if server.Disk, err = utils.InitDisk(); err != nil {
		global.Log.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &server, err
	}

	return &server, nil
}
