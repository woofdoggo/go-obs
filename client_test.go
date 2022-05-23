package go_obs_test

import (
	"testing"

	obs "github.com/woofdoggo/go-obs"
)

func connect(t *testing.T) *obs.Client {
	c := obs.Client{}
	needsAuth, _, err := c.Connect("localhost:4440")
	if err != nil {
		t.Error(err)
	}
	if !needsAuth {
		return &c
	}
	if needsAuth {
		err = c.Authenticate("password")
		if err != nil {
			t.Error(err)
		}
	}
	return &c
}

func TestConnect(t *testing.T) {
	connect(t)
}
