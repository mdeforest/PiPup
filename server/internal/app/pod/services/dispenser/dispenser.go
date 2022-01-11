package dispenser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const treatMultiplier int = 2

type Response struct {
}

func DispenseTreats() error {
	client := &http.Client{}

	// request treatMultipler num treats
	req, err := http.NewRequest("POST", "?num="+strconv.Itoa(treatMultiplier), nil)

	if err != nil {
		return errors.New("failed to dispense treats")
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return errors.New("failed to dispense treats")
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("failed to dispense treats")
	}

	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)

	return nil
}
