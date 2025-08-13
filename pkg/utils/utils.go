package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func FetchGender(name string) (string, error) {
	encodedName := url.QueryEscape(name)
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", encodedName)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("an error occured", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
		return "", err
	}
	var genderResp GenderResponse
	if err := json.Unmarshal(body, &genderResp); err != nil {
		fmt.Println("Error parsing json", err)
		return "", err
	}
	if genderResp.Gender == "" {
		fmt.Printf("Gender for %s could not be determined\n", name)
	} else {
		fmt.Printf("Name: %s, Gender: %s, Probability: %.2f%%\n",
			genderResp.Name, genderResp.Gender, genderResp.Probability*100)
	}

	return genderResp.Gender, nil
}
