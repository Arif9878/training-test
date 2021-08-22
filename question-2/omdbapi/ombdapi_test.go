package omdbapi_test

import (
	"testing"

	"github.com/Arif9878/stockbit-test/question-2/helper"
	"github.com/Arif9878/stockbit-test/question-2/models"
	"github.com/Arif9878/stockbit-test/question-2/omdbapi"
)

var apiKey = helper.GetEnv("API_KEY_OMDB")

func Test_ApiNoKey(t *testing.T) {
	api := omdbapi.Init("")
	_, err := api.Search(&models.QueryParams{Search: "Her"})
	if err == nil {
		t.Errorf("Expected to fail")
	}
	if err != nil {
		expectedErrorMsg := "Status Code 401 received from IMDB"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected- %s, Got- %s", expectedErrorMsg, err)
		}
	}
}

func Test_Search(t *testing.T) {
	tests := []struct {
		query *models.QueryParams
		title string
		year  string
	}{
		{&models.QueryParams{Search: "Fight Club", Page: "1"},
			"Fight Club",
			"1999",
		},
		{&models.QueryParams{Search: "Her"},
			"Her",
			"2013",
		},
		{&models.QueryParams{Search: "Macbeth", Page: "1"},
			"Macbeth",
			"2015",
		},
	}
	api := omdbapi.Init(apiKey)
	for i, item := range tests {
		resp, err := api.Search(item.query)
		if err != nil {
			t.Errorf("Test[%d]: %s", i, err)
			continue
		}
		if resp.Search[0].Title != item.title {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.title, resp.Search[0].Title)
			continue
		}
		if resp.Search[0].Year != item.year {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.year, resp.Search[0].Year)
			continue
		}
	}
}

func Test_MovieByImdbID(t *testing.T) {
	tests := []struct {
		id    string
		title string
		year  string
	}{
		{
			"tt0137523",
			"Fight Club",
			"1999",
		},
		{
			"tt1798709",
			"Her",
			"2013",
		},
		{
			"tt2884018",
			"Macbeth",
			"2015",
		},
	}

	api := omdbapi.Init(apiKey)

	for i, item := range tests {
		resp, err := api.GetImdbDetail(item.id)
		if err != nil {
			t.Errorf("Test[%d]: %s", i, err)
			continue
		}
		if resp.Title != item.title {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.title, resp.Title)
			continue
		}
		if resp.Year != item.year {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.year, resp.Year)
			continue
		}
	}
}
