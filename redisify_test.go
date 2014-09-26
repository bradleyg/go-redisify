package goredisify

import (
	"testing"
)

func TestConnRedisLocalhost(t *testing.T) {
	client, err := Conn(nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if client.Addr != "" {
		t.Fatalf("Client address should empty when using localhost. Expected %s, Actual %s", "", client.Addr)
	}
}

func TestConnRedisUrl(t *testing.T) {
	client, err := Conn("redis://remotehost.com")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if client.Addr != "remotehost.com" {
		t.Fatalf("Client address should be set when REDIS_URL is set. Expected %s, Actual %s", "remotehost.com", client.Addr)
	}
}

func TestConnRedisUrlPassword(t *testing.T) {
	client, err := Conn("redis://user:pass@remotehost.com")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if client.Addr != "remotehost.com" {
		t.Fatalf("Client address should be set when REDIS_URL is set. Expected %s, Actual %s", "remotehost.com", client.Addr)
	}

	if client.Password != "pass" {
		t.Fatalf("Client password should be set when REDIS_URL is set. Expected %s, Actual %s", "pass", client.Password)
	}
}
