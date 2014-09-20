package main

import (
	"fmt"
)

var validTypes = map[string]struct{}{
	"good":   struct{}{},
	"status": struct{}{},
}

func IsValidType(t string) bool {
	_, ok := validTypes[t]
	return ok
}

func ValidateTypeData(ty string, data interface{}) error {
	switch ty {
	case "good":
		return validateGoodType(data)

	case "status":
		return validateStatusType(data)

	default:
		return fmt.Errorf("unknown type: %s", ty)
	}
}

func validateGoodType(data interface{}) error {
	arr, ok := data.([]interface{})
	if !ok {
		return fmt.Errorf("expected input to be an array, not: %T", data)
	}

	for i, v := range arr {
		item, ok := v.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected entry %d to be an object, not: %T", i, v)
		}

		var name string
		var good bool
		var foundName, foundGood bool

		for k, v := range item {
			if k == "name" {
				vs, ok := v.(string)
				if !ok {
					return fmt.Errorf("expected entry %d's name to be a string, not: %T", i, v)
				}

				name = vs
				foundName = true
			} else if k == "good" {
				vb, ok := v.(bool)
				if !ok {
					return fmt.Errorf("expected entry %d's 'good' to be a bool, not: %T", i, v)
				}

				good = vb
				foundGood = true
			} else {
				return fmt.Errorf("unknown key in entry %d: %s", i, k)
			}
		}

		if !foundName {
			return fmt.Errorf("did not find 'name' in entry %d", i)
		}
		if !foundGood {
			return fmt.Errorf("did not find 'good' in entry %d", i)
		}

		// TODO: do we need to use these?
		_ = name
		_ = good
	}

	return nil
}

func validateStatusType(data interface{}) error {
	arr, ok := data.([]interface{})
	if !ok {
		return fmt.Errorf("expected input to be an array, not: %T", data)
	}

	for i, v := range arr {
		item, ok := v.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected entry %d to be an object, not: %T", i, v)
		}

		var name, status string
		var foundName, foundStatus bool

		for k, v := range item {
			vs, ok := v.(string)
			if !ok {
				return fmt.Errorf("expected entry %d to have string values, not: %T", i, v)
			}

			if k == "name" {
				name = vs
				foundName = true
			} else if k == "status" {
				status = vs
				foundStatus = true
			} else {
				return fmt.Errorf("unknown key in entry %d: %s", i, k)
			}
		}

		if !foundName {
			return fmt.Errorf("did not find 'name' in entry %d", i)
		}
		if !foundStatus {
			return fmt.Errorf("did not find 'status' in entry %d", i)
		}

		// TODO: do we need to use these?
		_ = name
		_ = status
	}

	return nil
}
