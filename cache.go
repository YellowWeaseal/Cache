package cache

import "errors"

type InMemoryCache interface {
	Set(key string,value interface{})
	Get(key string)
	Delete(key string)
}

type Cache struct {
	key string
	value interface{}
	Items map[string]interface{}
}

func (s *Cache) Set(key string,value interface{}){
	s.Items[key]=value
	return
}
func (g *Cache) Get(key string) interface{}{
	 _, ok:=g.Items[key]
	if !ok{
		return nil
	}
	return g.Items[key]
}
func (d *Cache) Delete(key string)error{
	_ , ok := d.Items[key]
	if !ok{
		return errors.New("element not found")
	}
	delete(d.Items,key)
	return nil
}
func (n *Cache) New() map[string]interface{} {
	items:=make(map[string]interface{})
	return items
}

