package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateGoodType(t *testing.T) {
	successes := []interface{}{
		[]interface{}{
			map[interface{}]interface{}{
				"name": "foo",
				"good": true,
			},
		},
		[]interface{}{},
		[]interface{}{
			map[interface{}]interface{}{
				"name": "foo",
				"good": true,
			},
			map[interface{}]interface{}{
				"name": "bar",
				"good": false,
			},
		},
	}

	var err error

	for _, v := range successes {
		err = validateGoodType(v)
		assert.NoError(t, err)
	}

	type Failure struct {
		Data  interface{}
		Error string
	}

	failures := []Failure{
		{
			Data:  1234,
			Error: "expected input to be an array, not: int",
		},
		{
			Data: []interface{}{
				1234,
			},
			Error: "expected entry 0 to be an object, not: int",
		},
		{
			Data: []interface{}{
				map[interface{}]interface{}{
					1234: 5678,
				},
			},
			Error: "expected entry 0 to have string keys, not: int",
		},
		{
			Data: []interface{}{
				map[interface{}]interface{}{
					"name": 1234,
				},
			},
			Error: "expected entry 0's name to be a string, not: int",
		},
		{
			Data: []interface{}{
				map[interface{}]interface{}{
					"name": "asdf",
					"good": "badval",
				},
			},
			Error: "expected entry 0's 'good' to be a bool, not: string",
		},
		{
			Data: []interface{}{
				map[interface{}]interface{}{
					"good": true,
				},
			},
			Error: "did not find 'name' in entry 0",
		},
		{
			Data: []interface{}{
				map[interface{}]interface{}{
					"name": "foo",
				},
			},
			Error: "did not find 'good' in entry 0",
		},
		{
			Data: []interface{}{
				map[interface{}]interface{}{
					"name":  "foo",
					"good":  true,
					"other": "value",
				},
			},
			Error: "unknown key in entry 0: other",
		},
	}

	for _, failure := range failures {
		err = validateGoodType(failure.Data)
		assert.EqualError(t, err, failure.Error)
	}
}
