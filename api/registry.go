package api

import (
	"fmt"
	"reflect"
	"strings"
)

var typeRegistry = make(map[string]reflect.Type)

func init() {
	myTypes := []interface{}{
		pingController{},
	}
	for _, v := range myTypes {
		register(v)
	}
}

func register(t ...interface{}) {
	for _, typ := range t {
		typeRegistry[strings.ToLower(fmt.Sprintf("%T", typ))] = reflect.TypeOf(typ)
	}
}

func makeInstance(name string) (interface{}, error) {
	n, ok := typeRegistry[name]
	if ok {
		return reflect.New(n).Interface(), nil
	}

	return nil, fmt.Errorf("unhandled registry %s", name)
}

func getEntityFromReq(name string, isList bool) (val interface{}, err error) {
	return makeInstance("api." + name)
}
