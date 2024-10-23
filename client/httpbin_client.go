package client

import (
	"context"

	"github.com/RianAsmara/personal-finance-advisor-api/model"
)

type HttpBinClient interface {
	PostMethod(ctx context.Context, requestBody *model.HttpBin, response *map[string]interface{})
}
