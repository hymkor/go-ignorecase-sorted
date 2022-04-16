package ignoreCaseSorted

import (
	"sort"
	"strings"
)

type KeyValue[T any] struct {
	Key   string
	Value T
}

type Dictionary[T any] struct {
	maps  map[string]KeyValue[T]
	order []string
}

func (d *Dictionary[T]) Store(key string, val T) {
	lowerKey := strings.ToLower(key)
	if d.maps == nil {
		d.maps = make(map[string]KeyValue[T])
	}
	d.maps[lowerKey] = KeyValue[T]{Key: key, Value: val}
	d.order = nil
}

func (d *Dictionary[T]) Delete(key string) {
	if d.maps == nil {
		return
	}
	lowerKey := strings.ToLower(key)
	delete(d.maps, lowerKey)
	d.order = nil
}

func (d *Dictionary[T]) Load(key string) (val T, ok bool) {
	if d.maps == nil {
		return
	}
	lowerKey := strings.ToLower(key)
	var v KeyValue[T]
	v, ok = d.maps[lowerKey]
	if ok {
		val = v.Value
	}
	return
}

func (d *Dictionary[t]) makeOrder() {
	if d.order == nil {
		d.order = make([]string, 0, len(d.maps))
		for key := range d.maps {
			d.order = append(d.order, key)
		}
		sort.Strings(d.order)
	}
}

func (d *Dictionary[T]) Range(f func(key string, val T)) {
	d.makeOrder()
	for _, key := range d.order {
		pair := d.maps[strings.ToLower(key)]
		f(pair.Key, pair.Value)
	}
}
