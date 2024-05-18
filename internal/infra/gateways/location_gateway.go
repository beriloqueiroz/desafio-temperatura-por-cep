package gateways

import "context"

type GetLocationGatewayImpl struct {
	Ctx context.Context
}

func (gt *GetLocationGatewayImpl) GetLocationByZipCode(ctx context.Context, zipCode string) (string, error) {
	gt.Ctx.Done()
	return "12354", nil
}
