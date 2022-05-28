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

	switch v := value.(type) {
	case CacheEntryString:

		if v.ttl == int64(-1) {
			return
		}

		if v.ttl <= time.Now().UnixMilli() {
			delete(Cache, key)
		}
	}
}

func Set(key string, value string) {
	entry := CacheEntryString{"string", value, -1, len(value)}
	Cache[key] = entry
}

func Get(key string) (string, error) {
	// check ttl
	CheckExpire(key)

	value, ok := Cache[key]

	if !ok {
		return "", errors.New("key not found")
	}

	v, ok := value.(CacheEntryString)

	if !ok {
		return "", errors.New("key of wrong type")
	}

	return v.value, nil
}

func Del(key string) {
	delete(Cache, key)
}

func Expire(key string, ttl int64) int {
	value, ok := Cache[key]
	if !ok {
		return 0
	}

	switch v := value.(type) {
	case CacheEntryString:
		Cache[key] = CacheEntryString{
			valueType: "string",
			value:     v.value,
			ttl:       time.Now().UnixMilli() + ttl*1000,
		}
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

	switch v := value.(type) {
	case CacheEntryString:
		// ttl not set
		if v.ttl == -1 {
			return -1
		}

		return (v.ttl - time.Now().UnixMilli()) / int64(1000)

	default:
		return -2
	}

}
