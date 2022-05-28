package main

type CacheEntryString struct {
	valueType string
	value     string
	ttl       int64
	valueSize int
}

var Cache = make(map[string]interface{})
