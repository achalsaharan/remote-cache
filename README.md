## Commands

## String Values

### SET

SET key "value"

| Return Value | Description       |
| ------------ | ----------------- |
| "OK"         | SET was performed |

### GET

GET key

| Return Value | Description     |
| ------------ | --------------- |
| string value | key exists      |
| nil          | key not present |

## Hash Map Values

### HSET

HSET key field value

| Return Value | Description                  |
| ------------ | ---------------------------- |
| int          | no of fields that were added |

### HGET

HGET key field value

| Return Value | Description |
| ------------ | ----------- |
| string       | value       |

## Common

### DEL

DEL key

| Return Value | Description                  |
| ------------ | ---------------------------- |
| integer      | no of keys that were removed |

### EXPIRE

EXPIRE key

| Return Value | Description                                  |
| ------------ | -------------------------------------------- |
| 1            | if timeout was set                           |
| 0            | if timeout was not set, eg key doesn't exist |

### TTL

TTL key

| Return Value | Description                             |
| ------------ | --------------------------------------- |
| -1           | key exists but has no associated expire |
| -2           | key does not exist                      |
| Positive Int | TTL value                               |

---

## Road map

1. Add support for different value types (string and hash map)
2. Concurrency: Try out different type of locking strategies
3. Write code to check if locking is working properly (fire multiple requests with go routines?)
4. Communication: start with HTTP, move to raw TCP. Document the requests size and performance improvements
5. Maintaining memory used by cache
6. Cache eviction mechanism (cron job based, sampling based like Redis)
7. Add atomic commands like INCR DECR for int values

---

## Problems

1. Supporting different value types (int and hash map).

   - [ ] make different maps for different type of values?

   - [ ] is there a way to have a map that supports multiple values

   - [ ] Update the website

2. keeping track of memory consumption

   - what method to use?
