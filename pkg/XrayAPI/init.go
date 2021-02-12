package XrayAPI

import (
	"fmt"
	"github.com/crossfw/Air-Universe/pkg/structures"
	"github.com/xtls/xray-core/app/proxyman/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
)

func InitApi(cfg *structures.BaseConfig, xrayCtl *structures.XrayController) (err error) {
	xrayCtl.CmdConn, err = grpc.Dial(fmt.Sprintf("%s:%d", cfg.Proxy.APIAddress, cfg.Proxy.APIPort), grpc.WithInsecure())
	if err != nil {
		return err
	}
	hsClient := command.NewHandlerServiceClient(xrayCtl.CmdConn)
	ssClient := statsService.NewStatsServiceClient(xrayCtl.CmdConn)

	xrayCtl.HsClient = &hsClient
	xrayCtl.SsClient = &ssClient

	return
}