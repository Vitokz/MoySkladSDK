package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Vitokz/MoySkladSDK/models/entity"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Params структура параметров
type Params struct {
	Filters []Filter
	Query   map[string]string
}

// Filter структура фильтра
type Filter struct {
	Key      string
	Value    string
	Operator string
}

// AddFilter Допустимые операторы: ['=', '>', '<', '>=', '<=', '!=', '~', '~=', '=~']
func (params *Params) AddFilter(key, operator, value string) {
	params.Filters = append(params.Filters, Filter{key, value, operator})
}

// AddQuery добавляет параметр к запросу
func (params *Params) AddQuery(key string, value string) {
	if params.Query == nil {
		params.Query = make(map[string]string)
	}

	params.Query[key] = value
}

// getFilterString склеивает фильтры в строку и возвращает
func (params Params) getFilterString() (filters string) {
	for _, filter := range params.Filters {
		filters = fmt.Sprintf("%s%s%s%s;", filters, filter.Key, filter.Operator, filter.Value)
	}

	return
}

//MakeRequest Создает и запускает одиночный запрос, если же очередь забита, ожидает пока не освободиться место
func (ms *moySklad) MakeRequest(method, endpoint string, body []byte, params interface{}, resp interface{}) error {
	req, err := createRequest(method, endpoint, body, params, ms.credentials)
	if err != nil {
		return errors.Errorf("failed to create request : %s", err)
	}

	ms.requestLimit <- 1
	response, err := ms.api.Do(req)
	if err != nil {
		return errors.Errorf("%s request to %s failed with status: %d , error text : %s", method, endpoint, response.StatusCode, err.Error())
	}
	defer func() {
		<-ms.requestLimit
	}()

	if response.StatusCode >= 301 {
		msErr := new(entity.Errors)
		bod, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(bod, &msErr)
		if err != nil {
			msErrors := make([]entity.Errors, 0)
			err = json.Unmarshal(bod, &msErrors)
			if err != nil {
				return err
			}
			msErr = &msErrors[0]
		}
		return errors.Errorf("create failed with status: %v and body: %s", response.StatusCode, msErr.Errors[0].Error)
	} else if err != nil {
		return err
	}

	if resp != nil {
		return json.NewDecoder(response.Body).Decode(resp)
	}
	return nil
}

//MakeListRequest Функция которая поможет вывести список всех сущностей будь то товар или что-либо еще
func (ms *moySklad) MakeListRequest(method, endpoint string, body []byte, params interface{}) (interface{}, error) {
	req, err := createRequest(method, endpoint, body, params, ms.credentials)
	if err != nil {
		return nil, err
	}

	ms.requestLimit <- 1
	response, err := ms.api.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		<-ms.requestLimit
	}()

	list := new(entity.ListResponse)
	resp, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(resp, list)
	if err != nil {
		return nil, err
	}

	//// Если же количество сущностей больше лимита(1000), то создаем ссылки на вывод всех остальных товаров и запускаем request
	//if list.Meta.Size >= LimitRows {
	//	urls := listUrls(list, method, endpoint, ms.credentials) // Создает список ссылок для вывода сущностей
	//	for i := range *urls {
	//		ms.requestLimit <- 1
	//		response, err = ms.api.Do(&(*urls)[i])
	//		if err != nil {
	//			return nil, err
	//		}
	//		<-ms.requestLimit
	//		vremList := new(entity.ListResponse)
	//		resp, err = ioutil.ReadAll(response.Body)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		err = json.Unmarshal(resp, vremList)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		list.Rows = append(list.Rows.([]entity.Product), vremList.Rows.([]entity.Product))
	//	}
	//}
	return nil, nil
}

// setHeaders Устанавливает стандартные header для request
func setHeaders(req *http.Request, creds string) {
	req.Header.Set(`Authorization`, fmt.Sprintf(`Basic %s`, creds))
	req.Header.Set(`Accept`, `application/json;charset=utf-8`)
	req.Header.Set(`Content-Type`, `application/json`)
}

// addQuery Устанавливает параметры запроса, также смотрит фильтры для МойСклад
func addQuery(req *http.Request, params Params) {
	query := req.URL.Query()
	if filter := params.getFilterString(); filter != "" {
		query.Add("filter", filter)
	}

	for key, value := range params.Query {
		query.Add(key, value)
	}

	req.URL.RawQuery = query.Encode()
}

// listUrls Создает список ссылок для того, чтобы достать все сущности
func listUrls(list *entity.ListResponse, method, endpoint, credentials string) *[]http.Request {
	size := list.Meta.Size
	offset := LimitRows

	requests := make([]http.Request, 0)
	for ; offset < size; offset += LimitRows {
		params := new(Params)
		params.AddQuery("offset", strconv.Itoa(offset))
		params.AddQuery("limit", strconv.Itoa(LimitRows))

		req, _ := createRequest(method, endpoint, nil, params, credentials)
		requests = append(requests, *req)
	}
	return &requests
}

// createRequest Создает request
func createRequest(method, endpoint string, body []byte, params interface{}, credentials string) (*http.Request, error) {
	URL := fmt.Sprintf("%s%s", APIPath, endpoint)
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	setHeaders(req, credentials)
	if params != nil {
		p := params.(Params)
		addQuery(req, p)
	}

	return req, nil
}
