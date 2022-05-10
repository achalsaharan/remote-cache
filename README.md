## Commands

### SET

**SET key "value"**

| Return Value | Description        |
| ------------ | ------------------ |
| "OK"         | SET was performend |

### GET

GET key

| Return Value | Description     |
| ------------ | --------------- |
| string value | key exists      |
| nil          | key not present |

## DEL

DEL key

| Return Value | Description                  |
| ------------ | ---------------------------- |
| integer      | no of keys that were removed |

## EXPIRE

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

## Roadmap

1. Add support for different value types (string and int)
2. Concurrency: Try out different type of locking stratergies
3. Write code to check if locking is working properly (fire multiple requests with go routines?)
4. Communication: start with HTTP, move to raw TCP. Document the requests size and performance improvements
5. Maintaining memory used by cache
6. Cache eviction mechanism (cron job based, sampling based like Redis)
7. Add atomic commands like INCR DECR for int values

---

## Problems

1. Supporting different value types (int and string).

   - make different maps for different type of values?
   - is there a way to have a map that supports multiple values

2. keeping track of memory consumpltion

   - what method to use?
