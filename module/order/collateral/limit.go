package collateral

import (
	"github.com/whitebit-exchange/go-sdk"
)

const limitEndpointUrl = "/api/v4/order/collateral/limit"

type LimitOrder struct {
	MarketOrder
	Price string `json:"price"`
}

type LimitOrderParams struct {
	Market        string `json:"market"`
	Amount        string `json:"amount"`
	Side          string `json:"side"`
	Price         string `json:"price"`
	ClientOrderId string `json:"clientOrderId,omitempty"`
}

type limitEndpoint struct {
	whitebit.AuthParams
	LimitOrderParams
}

func newLimitEndpoint(params LimitOrderParams) *limitEndpoint {
	return &limitEndpoint{
		AuthParams:       whitebit.NewAuthParams(limitEndpointUrl),
		LimitOrderParams: params,
	}
}
