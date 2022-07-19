package models

import (
	"reflect"
	"time"
)

type Product struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rating      int8      `json:"rating", omitempty`
	CreateAt    time.Time `json:"createAt"`
}

func (p *Product) GetMapString() map[string]interface{} {
	pMap := make(map[string]interface{}, 1)

	v := reflect.Indirect(reflect.ValueOf(p))
	for j := 0; j < v.NumField(); j++ {
		f := v.Type().Field(j)
		tagName := f.Tag.Get("json")
		pMap[tagName] = v.Field(j)
	}

	return pMap
}
