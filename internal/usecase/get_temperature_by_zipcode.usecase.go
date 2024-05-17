package usecase

import (
	"context"
	"errors"

	"github.com/beriloqueiroz/desafio-temperatura-por-cep/internal/entity"
)

type GetTemperByZipCodeUseCase struct {
	LocationGateway    LocationGateway
	TemperatureGateway TemperatureGateway
}

type GetTemperByZipCodeUseCaseOutput struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func (uc *GetTemperByZipCodeUseCase) Execute(ctx context.Context, zipCode string) (GetTemperByZipCodeUseCaseOutput, error) {
	output := GetTemperByZipCodeUseCaseOutput{}
	zipCodeObj, err := entity.NewZipCode(zipCode)
	if err != nil {
		return output, err
	}
	location, err := uc.LocationGateway.GetLocationByZipCode(ctx, zipCode)
	if err != nil {
		return output, errors.New("can not find zipcode")
	}
	temperature, err := uc.TemperatureGateway.GetTemperatureByLocation(ctx, location)
	if err != nil {
		return output, err
	}
	tempLocation, err := entity.NewTemperatureLocation(zipCodeObj, temperature)
	if err != nil {
		return output, err
	}
	output.TempC = tempLocation.TempC
	output.TempF = tempLocation.TempF
	output.TempK = tempLocation.TempK
	return output, nil
}
