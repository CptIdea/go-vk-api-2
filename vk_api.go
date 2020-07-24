//Golang VK api by Captain Idea
package vk

//Объект, описывающий текущую сессию, все методы вызываются через него
type Session struct {
	Token          string
	Version        string
	SkipAutoResend bool
}
type Request map[string]interface{}
