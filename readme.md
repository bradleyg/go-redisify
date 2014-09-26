# goredisify
--
    import "github.com/bradleyg/go-redisify"

goredisify returns a redis connection via https://github.com/hoisie/redis for a
given url. If no url if given localhost is used.

## Usage

#### func  Conn

```go
func Conn(redisUrl interface{}) (*redis.Client, error)
```
Conn takes a string containing a URL for a redis server and returns a redis
connection. Passing nil as the redisUrl will look for a server locally.
