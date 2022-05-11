package grpc

import (
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/grpc/proto"
	"github.com/jau1jz/cornus"
	slog "github.com/jau1jz/cornus/commons/log"
	"google.golang.org/grpc"
)

var grpcConfig struct {
	GRPC struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"grpc"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&grpcConfig)
	SignPrepare()
}

var SignGrpc *SignStatus

type SignStatus struct {
	SignClient proto.DataControllerClient
	Status     int
}

func SignPrepare() {
	SignGrpc = &SignStatus{
		Status: common.SignGrpcConnectBefore,
	}
	SignGrpc.Connect()
}

func (ds *SignStatus) Connect() {
	if ds.Status == common.SignGrpcConnectBefore {
		ds.Status = common.SignGrpcConnecting
		//for {
		conn, err := grpc.Dial(grpcConfig.GRPC.Host+":"+grpcConfig.GRPC.Port, grpc.WithInsecure())
		if err == nil {

			ds.SignClient = proto.NewDataControllerClient(conn)
			ds.Status = common.SignGrpcConnected
		} else {
			slog.Slog.ErrorF(nil, "grpcClient connect failed")
		}

		//time.Sleep(3 * time.Second)
		//}
	}
}
