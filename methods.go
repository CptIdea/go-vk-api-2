package vk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type longpoll struct {
	Response map[string]string
}

func (vk *Session) UpdateCheck(GroupId int) Updates {
	var upd Updates
	for len(upd.Updates) == 0 {
		url := "https://api.vk.com/method/groups.getLongPollServer?GroupId=" + strconv.Itoa(GroupId) + "&v=" + vk.Version + "&access_token=" + vk.Token
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		var jsoned longpoll
		err = json.Unmarshal(bs[:n], &jsoned)
		if err != nil {
			fmt.Println("error:", err)
		}
		url = jsoned.Response["server"] + "?act=a_check&key=" + jsoned.Response["key"] + "&ts=" + jsoned.Response["ts"] + "&wait=25"
		resp, err = http.Get(url)
		bs = make([]byte, 8112)
		n, err = resp.Body.Read(bs)
		err = json.Unmarshal(bs[:n], &upd)
		if err != nil {
			fmt.Println("error:", err)
		}
	}
	return upd
}

// Отправляет сообщение message в чат ToId
func (vk *Session) SendMessage(ToId int, message string) string {
	message = strings.Replace(message, " ", "+", -1)
	newrequest := "peer_id=" + strconv.Itoa(ToId) + "&" + "message=" + message
	return string(vk.SendRequest("messages.send", newrequest))
}

func (vk *Session) GetConversationInfo(peers []int, ext int, fields []string, GroupId int) []Conversation {
	var params string
	var output struct {
		Response struct{ Items []Conversation }
	}
	params = "peer_ids="
	for i := 0; i < len(peers); {
		params = params + strconv.Itoa(peers[i])
		i++
		if i < len(peers) {
			params = params + ","
		}
	}
	if ext == 1 {
		params = params + "&extended=1&fields="
		for i := 0; i < len(fields); {
			params = params + fields[i]
			i++
			if i < len(fields) {
				params = params + ","
			}
		}
	}
	params = params + "&GroupId=" + strconv.Itoa(GroupId)
	resp := vk.SendRequest("messages.getConversationsById", params)
	fmt.Println(string(resp))
	err := json.Unmarshal(resp, &output)
	if err != nil {
		fmt.Println("error:", err)
	}
	return output.Response.Items
}

//Возвращает информацию о пользователях в массиве объектов User
func (vk *Session) GetUsersInfo(Ids []int, fields []string) []User {
	var params string
	var output struct{ Response []User }
	params = "user_ids="
	for i := 0; i < len(Ids); {
		params = params + strconv.Itoa(Ids[i])
		i++
		if i < len(Ids) {
			params = params + ","
		}
	}
	params = params + "&fields="
	for i := 0; i < len(fields); {
		params = params + fields[i]
		i++
		if i < len(fields) {
			params = params + ","
		}
	}
	params = params + "&name_case=nom,gen,dat,acc,ins,abl"
	resp := vk.SendRequest("users.get", params)
	err := json.Unmarshal(resp, &output)
	if err != nil {
		fmt.Println("error:", err)
	}
	return output.Response
}

//Возвращает информацию о сообществах в массиве объектов Group
func (vk *Session) GroupGetById(GroupIds []int, fields []string) []Group {
	var params string
	var output struct{ Response []Group }
	params = "GroupIds="
	for i := 0; i < len(GroupIds); i++ {
		GroupIds[i] = GroupIds[i] * -1
	}
	for i := 0; i < len(GroupIds); {
		params = params + strconv.Itoa(GroupIds[i])
		i++
		if i < len(GroupIds) {
			params = params + ","
		}
	}
	if len(fields) != 0 {
		params = params + "&fields="
		for i := 0; i < len(fields); {
			params = params + fields[i]
			i++
			if i < len(fields) {
				params = params + ","
			}
		}
	}
	resp := vk.SendRequest("groups.getById", params)
	err := json.Unmarshal(resp, &output)
	if err != nil {
		fmt.Println("error:", err)
	}
	return output.Response
}

func (vk *Session) SendRequest(method, params string) []byte {
	url := "https://api.vk.com/method/" + method + "?" + params + "&v=" + vk.Version + "&access_token=" + vk.Token
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	bs := make([]byte, 1014)
	n, err := resp.Body.Read(bs)
	return bs[:n]
}
