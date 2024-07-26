package pokeapi

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"time"
)

const mapLimit = 20

var mapState = MapState{20, 0, 0, true}
var cache = pokecache.NewCache(1 * time.Hour)

type MapState struct {
	pageLength uint
	count      uint
	index      uint
	increasing bool
}

func MapRequest() (PokeResponseList, error) {
	url := fmt.Sprintf(
		"%s?limit=%d&offset=%d",
		POKEAPI_URL,
		mapState.pageLength,
		mapState.index*mapState.pageLength,
	)
	data, exists := cache.Get(url)
	var err error
	if !exists {
		data, err = GetRequest(url)
		if err != nil {
			return PokeResponseList{}, err
		}
		cache.Add(url, data)
	}

	res := PokeResponseList{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return PokeResponseList{}, err
	}
	return res, nil
}

func HandleMap() error {
	list, err := MapRequest()
	if err != nil {
		return err
	}
	for _, location := range list.Results {
		fmt.Println(location.Name)
	}
	mapState.count = uint(list.Count)
	return nil
}

func Map() error {
	if mapState.count == 0 && mapState.index > 0 {
		return fmt.Errorf("count should never be 0 when the index is not 0")
	} else if mapState.index*mapState.pageLength > mapState.count {
		return fmt.Errorf("you've reached the end of the location area list")
	}
	HandleMap()
	if mapState.increasing {
		mapState.index++
	} else {
		mapState.index += 2
	}
	mapState.increasing = true
	return nil
}

func MapB() error {
	if mapState.index == 0 || mapState.count < 1 {
		return fmt.Errorf("you're at the beginning of the location area list")
	}
	if !mapState.increasing || mapState.index == 0 {
		mapState.index--
	} else {
		mapState.index -= 2
	}
	mapState.increasing = false
	HandleMap()
	return nil
}
