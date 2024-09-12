package times

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl string = "https://public.opendatasoft.com/"

type Result struct {
	City string `json:"gem_name"`
}

type CityResponse struct {
	Count   int      `json:"total_count"`
	Results []Result `json:"results"`
}

func getCity(pc4Code int) (string, error) {
	formattedURL := fmt.Sprintf("%s/api/explore/v2.1/catalog/datasets/georef-netherlands-postcode-pc4/records?select=gem_name&where=pc4_code%%3D%d&limit=20",
		baseUrl, pc4Code,
	)
	response, err := http.Get(formattedURL)
	if err != nil {
		return "", err
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var cityResponse CityResponse
	err = json.Unmarshal(responseData, &cityResponse)
	if err != nil {
		return "", err
	}
	if len(cityResponse.Results) > 1 {
		return "", fmt.Errorf("somehow there was more than 1 city returned")
	}
	return string(cityResponse.Results[0].City), nil
}

func CallAPI(pc4Code int) (string, error) {
	city, err := getCity(pc4Code)
	if err != nil {
		return "", err
	}
	return city, nil
}
