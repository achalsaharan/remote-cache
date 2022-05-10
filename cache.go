package main

type CacheEntry struct {
	valueType string
	value     string
	ttl       int64
	// valueSize int
}

var Cache = make(map[string]CacheEntry)
