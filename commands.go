package main

import (
	"errors"
	"time"
)

func CheckExpire(key string) {
	value, ok := Cache[key]

	// key not present
	if !ok {
		return
	}

	// no expiry set on key
	if value.ttl == int64(-1) {
		return
	}

	// delte key
	if value.ttl <= time.Now().UnixMilli() {
		delete(Cache, key)
	}

}

func Set(key string, value string) {
	entry := CacheEntry{"string", value, -1}
	Cache[key] = entry
}

func Get(key string) (string, error) {
	// check ttl
	CheckExpire(key)

	value, ok := Cache[key]

	if !ok {
		return "", errors.New("key not found")
	}

	return value.value, nil
}

func Del(key string) {
	delete(Cache, key)
}

func Expire(key string, ttl int64) int {
	value, ok := Cache[key]
	if !ok {
		return 0
	}

	Cache[key] = CacheEntry{
		valueType: "string",
		value:     value.value,
		ttl:       time.Now().UnixMilli() + ttl,
	}

	return 1
}

func Ttl(key string) int64 {
	// get the key from cache
	value, ok := Cache[key]
	if !ok {
		// return -2 if key not found
		return -2
	}

	return (time.Now().UnixMilli() - value.ttl) / int64(1000)
}
