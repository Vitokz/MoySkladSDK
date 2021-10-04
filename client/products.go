package client

import (
	"encoding/json"
	"github.com/Vitokz/MoySkladSDK/models/entity"
	"github.com/Vitokz/MoySkladSDK/pkg"
	"github.com/pkg/errors"
)

type Product interface {
	Update(idProduct string, prd entity.Product) (*entity.Product, error)
	Create(product *entity.Product) (*entity.Product, error)
	CreateAndUpdateList(products []entity.Product) (*[]entity.Product, error)
	Delete(idProduct string) error
	DeleteList(products []entity.Product) error
	Take(idProduct string) (*entity.Product, error)
	TakeList(params interface{}) (*[]entity.Product, error)
}

type product struct {
	ms *moySklad
}

func productsClient(moySklad *moySklad) Product {
	var prd Product

	prd = &product{
		ms: moySklad,
	}

	return prd
}

func (p *product) Update(idProduct string, prd entity.Product) (*entity.Product, error) {
	method, endpoint := pkg.UpdateProduct(idProduct)
	body, err := json.Marshal(prd)
	if err != nil {
		return nil, err
	}
	response := new(entity.Product)
	err = p.ms.MakeRequest(method, endpoint, body, nil, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *product) Create(product *entity.Product) (*entity.Product, error) {
	method, endpoint := pkg.CreateProduct()
	body, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	response := new(entity.Product)
	err = p.ms.MakeRequest(method, endpoint, body, nil, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (p *product) CreateAndUpdateList(products []entity.Product) (*[]entity.Product, error) {
	if len(products) > 1000 {
		return nil, errors.New("max products slice len: 1000")
	} else if len(products) == 0 {
		return nil, nil
	}
	method, endpoint := pkg.CreateProduct()
	body, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	response := make([]entity.Product, 0)
	err = p.ms.MakeRequest(method, endpoint, body, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *product) Delete(idProduct string) error {
	method, endpoint := pkg.DeleteProduct(idProduct)

	err := p.ms.MakeRequest(method, endpoint, nil, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *product) DeleteList(products []entity.Product) error {
	if len(products) > 1000 {
		return errors.New("max products slice len: 1000")
	} else if len(products) == 0 {
		return nil
	}
	method, endpoint := pkg.DeleteProductList()
	body, err := json.Marshal(products)
	if err != nil {
		return err
	}

	err = p.ms.MakeRequest(method, endpoint, body, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *product) Take(idProduct string) (*entity.Product, error) {
	method, endpoint := pkg.TakeProduct(idProduct)
	response := new(entity.Product)

	err := p.ms.MakeRequest(method, endpoint, nil, nil, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (p *product) TakeList(params interface{}) (*[]entity.Product, error) {
	method, endpoint := pkg.TakeProductList()
	resp := new(entity.ListResponse)
	prds := make([]entity.Product, 0)
	err := p.ms.MakeRequest(method, endpoint, nil, params, &resp)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(resp.Rows)
	err = json.Unmarshal(data, &prds)
	if err != nil {
		return nil, err
	}
	return &prds, nil
}
