package cloudflare

import (
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

type Pool struct {
	PoolParams
	ID                string       `json:"id"`
	MinimumOrigins    int          `json:"minimum_origins"`
	ModifiedOn        time.Time    `json:"modified_on"`
}

type PoolParams struct {
	Description       string       `json:"description"`
	Enabled           bool         `json:"enabled"`
	MinimumOrigins    int          `json:"minimum_origins"`
	Monitor           string       `json:"monitor"`
	Name              string       `json:"name"`
	NotificationEmail string       `json:"notification_email"`
	Origins           []PoolOrigin `json:"origins"`
}

type PoolOrigin struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Enabled bool   `json:"enabled"`
}

type PoolResponse struct {
	Response
	Result Pool `json:"result"`
}

func (api *API) GetPoolByID(id string) (Pool, error) {
	res, err := api.makeRequest("GET", "/user/load_balancers/pools/"+id, nil)
	if err != nil {
		return Pool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r PoolResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Pool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

func (api *API) UpdatePool(id string, p PoolParams) error {
	res, err := api.makeRequest("PUT", "/user/load_balancers/pools/"+id, p)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	var r PoolResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}
