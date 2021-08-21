package imdbapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Arif9878/stockbit-test/question-2/models"
)

const baseURL = "http://www.omdbapi.com/?"

type OmdbApi struct {
	ApiKey string
}

func Init(ApiKey string) *OmdbApi {
	return &OmdbApi{ApiKey: ApiKey}
}

//Search for movies given a Title and year, Year is optional you can pass nil
func (api *OmdbApi) Search(query *models.QueryParams) (*models.SearchResponse, error) {
	resp, err := api.RequestAPI("search", query.Search, query.Page)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(models.SearchResponse)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return nil, errors.New(r.Error)
	}

	return r, nil
}

func (api *OmdbApi) GetImdbDetail(ID string) (*models.MovieDetail, error) {
	resp, err := api.RequestAPI("id", ID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(models.MovieDetail)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return r, errors.New(r.Error)
	}

	return r, nil
}

// function to call the API imdb
func (api *OmdbApi) RequestAPI(ApiCategory string, params ...string) (resp *http.Response, err error) {
	var URL *url.URL
	URL, err = url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	URL.Path += "/"
	parameters := url.Values{}
	parameters.Add("apikey", api.ApiKey)

	switch ApiCategory {
	case "search":
		parameters.Add("s", params[0])
		parameters.Add("page", params[1])
	case "id":
		parameters.Add("i", params[0])
	}

	URL.RawQuery = parameters.Encode()
	res, err := http.Get(URL.String())
	err = CheckStatusCode(res.StatusCode)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CheckStatusCode(status int) error {
	if status != 200 {
		return fmt.Errorf("Status Code %d received from IMDB", status)
	}
	return nil
}
