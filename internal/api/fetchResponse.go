package api

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"io"
	"net/http"
)

func (c *Client) FetchFeed(contx context.Context, destURL string) (*XMLResponse, error) {
	// Create request
	req, err := http.NewRequestWithContext(contx, "GET", destURL, http.NoBody)
	if err != nil {
		return &XMLResponse{}, err
	}
	req.Header.Set("User-Agent", "gator")

	// Make request, get response
	res, err := c.httpClient.Do(req)
	if err != nil {
		return &XMLResponse{}, err
	}

	// Close body after we are done
	defer res.Body.Close()

	// Check for problematic status codes
	if res.StatusCode > 299 {
		return &XMLResponse{}, errors.New(res.Status)
	}

	// Read response
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return &XMLResponse{}, err
	}

	// Unmarshal response
	xmlResponse := XMLResponse{}
	err = xml.Unmarshal(dat, &xmlResponse)
	if err != nil {
		return &XMLResponse{}, err
	}

	// Unscaping strings
	corrTitle := html.UnescapeString(xmlResponse.Channel.Title)
	xmlResponse.Channel.Title = corrTitle
	corrDesc := html.UnescapeString(xmlResponse.Channel.Description)
	xmlResponse.Channel.Title = corrDesc
	for _, item := range xmlResponse.Channel.Item {
		corrTitle := html.UnescapeString(item.Title)
		item.Title = corrTitle
		corrDesc := html.UnescapeString(item.Description)
		item.Title = corrDesc
	}

	return &xmlResponse, nil
}
