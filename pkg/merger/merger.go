package merger

import "encoding/json"

func Merge(from interface{}, to interface{}) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, to)
	return err
}

func Combine(from interface{}, to interface{}, target interface{}) error {
	tempFrom := make(map[string]interface{})
	tempTo := make(map[string]interface{})

	if to != nil {
		data, _ := json.Marshal(to)
		json.Unmarshal(data, &tempTo)
	}

	if from != nil {
		data, _ := json.Marshal(from)
		json.Unmarshal(data, &tempFrom)
	}

	for key, value := range tempFrom {
		tempTo[key] = value
	}

	data, _ := json.Marshal(tempTo)
	err := json.Unmarshal(data, target)
	return err
}
