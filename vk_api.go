//Golang VK api by Captain Idea
package vk

import (
	"math/rand"
	"time"
)

//Объект, описывающий текущую сессию, все методы вызываются через него
type Session struct {
	token          string
	version        string
	SkipAutoResend bool
}
type Request map[string]interface{}

func NewSession(token string, version string) Session {
	rand.Seed(time.Now().UnixNano())
	return Session{
		token:   token,
		version: version,
	}
}
