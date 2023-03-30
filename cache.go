package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]Item
}
type Item struct {
	value interface{}
	Created    time.Time
	Expiration int64
}
func (s *Cache) Set(key string,value interface{},duration time.Duration){

	var expiration int64
	if duration==0{
		duration = s.defaultExpiration
	}else if duration>0{
		expiration = time.Now().Add(duration).UnixNano()
	}

	s.Lock()

	defer s.Unlock()

	s.items[key]=Item{
		value: value,
		Expiration: expiration,
		Created: time.Now(),
	}

	return
}
func (g *Cache) Get(key string) interface{}{
	g.RLock()
	defer g.RUnlock()

	 item, ok:=g.items[key]
	if !ok{
		return nil
	}
	if   item.Expiration > 0{

		if time.Now().UnixNano() > item.Expiration{
			return nil
		}

	}

	return g.items[key]
}
func (d *Cache) Delete(key string)error{

	d.Lock()
	defer d.Unlock()

	_, ok := d.items[key]
	if !ok{
		return errors.New("element not found")
	}

	delete(d.items,key)
	return nil
}
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items:=make(map[string]Item)
	cache:=Cache{
		items: items,
		defaultExpiration: defaultExpiration,
		cleanupInterval: cleanupInterval,
	}
	return &cache
}

