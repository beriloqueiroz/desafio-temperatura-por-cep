package gateways

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GetLocationGatewayImpl struct {
	Ctx context.Context
}

type viaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	HasError    bool   `json:"erro"`
}

func (gt *GetLocationGatewayImpl) GetLocationByZipCode(ctx context.Context, zipCode string) (string, error) {
	location, err := buscaCep(zipCode)
	if err != nil {
		return "", err
	}
	gt.Ctx.Done()
	return location.Localidade, nil
}

func buscaCep(cep string) (*viaCEP, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var c viaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	if c.HasError {
		return nil, errors.New("can not find zipcode")
	}
	return &c, nil
}
