package main

import (
	"errors"
	"fmt"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func newCacheImpl() *cacheImpl {
	return &cacheImpl{}
}

type cacheImpl struct {
	items map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
	item, ok := c.items[k]

	if !ok {
		return "", false
	}

	return item, true

}

func (c *cacheImpl) Set(k, v string) {
	c.items[k] = v
}

// Функция удаления элемента
func (c *cacheImpl) Delete(key string) error {
	if _, ok := c.items[key]; !ok {
		return errors.New("Key not found!")
	}

	delete(c.items, key)
	return nil
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)
	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
}
