package top

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
)

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

type kvPairList struct {
	list []*kvPair
}

func (l kvPairList) Len() int {
	return len(l.list)
}

func (l kvPairList) Less(i, j int) bool {
	return l.list[i].key < l.list[j].key
}

func (l kvPairList) Swap(i, j int) {
	l.list[i], l.list[j] = l.list[j], l.list[i]
}

type kvPair struct {
	key   string
	value string
}

func newKVPairList() *kvPairList {
	return &kvPairList{
		list: make([]*kvPair, 0),
	}
}

func (l *kvPairList) load(data interface{}) error {
	if l.list == nil {
		l.list = make([]*kvPair, 0)
	}

	typ := reflect.TypeOf(data)
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return errors.New("data is nil")
		}
		value = value.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < value.NumField(); i++ {
		tag := strings.Split(typ.Field(i).Tag.Get("json"), ",")[0]
		v := value.Field(i)
		log.Printf("key:%v value:%v  \n", tag, v)

		l.list = append(l.list, &kvPair{
			key:   tag,
			value: fmt.Sprintf("%v", value.Field(i)),
		})
	}

	sort.Sort(l)
	return nil
}
