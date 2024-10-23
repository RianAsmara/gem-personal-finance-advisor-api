package impl

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/client"
	"github.com/RianAsmara/personal-finance-advisor-api/common"
	"github.com/RianAsmara/personal-finance-advisor-api/model"
	services "github.com/RianAsmara/personal-finance-advisor-api/service"
)

func NewHttpBinServiceImpl(httpBinClient *client.HttpBinClient) services.HttpBinService {
	return &httpBinServiceImpl{HttpBinClient: *httpBinClient}
}

type httpBinServiceImpl struct {
	client.HttpBinClient
}

func (h *httpBinServiceImpl) PostMethod(ctx context.Context) {
	httpBin := model.HttpBin{
		Name: "rizki",
	}
	var response map[string]interface{}
	h.HttpBinClient.PostMethod(ctx, &httpBin, &response)
	common.NewLogger().Info("log response service ", response)
}
