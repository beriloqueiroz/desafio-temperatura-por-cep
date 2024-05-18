package gateways

import "context"

type GetTemperatureGatewayImpl struct {
	Ctx context.Context
}

func (gt *GetTemperatureGatewayImpl) GetTemperatureByLocation(ctx context.Context, location string) (float64, error) {
	gt.Ctx.Done()
	return 10.5, nil
}
