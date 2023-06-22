package admin

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stephenmusangeya/go-shopify/admin/access"
)

// AccessScope represents the available Shopify API scopes.
func (c Client) AccessScope(id int) (*access.AccessScopeResponse, error) {
	var response access.AccessScopeResponse
	url := fmt.Sprintf("%s/products.json", c.baseUrl)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// req.SetBasicAuth(c.APIKey, c.Password)
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("User-Agent", c.UserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Handle the response
	if resp.StatusCode != http.StatusOK {
		return nil, NewError(string(responseBody), resp.StatusCode)
	}

	return &response, nil
}
