// Go-redisify returns a redis connection for a given url.
package goredisify

import (
	"net/url"

	"github.com/hoisie/redis"
)

// Conn takes a string containing a URL for a redis server and returns a redis connection.
// Passing nil as the redisUrl will look for a server locally. Returns a client
// with the type *redis.Client from the package https://github.com/hoisie/redis.
func Conn(redisUrl interface{}) (*redis.Client, error) {
	var value string

	switch redisUrl.(type) {
	case string:
		value = redisUrl.(string)
	default:
		value = ""
	}

	parsed, err := url.Parse(value)
	if err != nil {
		return nil, err
	}

	c := redis.Client{
		Addr: parsed.Host,
	}

	if parsed.User != nil {
		password, ok := parsed.User.Password()
		if ok {
			c.Password = password
		}
	}

	return &c, nil
}
