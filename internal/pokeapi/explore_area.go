package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (ExploreResp, error) {
	url := baseURL + "/location-area/"
	url = url + name

	// Cache hit
	cacheData, cacheExists := c.cache.Get(url)
	if cacheExists {
		var exploreData ExploreResp
		err := json.Unmarshal(cacheData, &exploreData)
		if err != nil {
			return ExploreResp{}, err
		}
		return exploreData, nil
	}

	// Cache Miss
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ExploreResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ExploreResp{}, fmt.Errorf("Explore call not OK: http code: %v, %s", res.StatusCode, res.Status)
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreResp{}, err
	}

	var resData ExploreResp
	err = json.Unmarshal(dat, &resData)
	if err != nil {
		return ExploreResp{}, err
	}
	c.cache.Add(url, dat)

	return resData, nil
}
