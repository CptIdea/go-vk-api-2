package vk

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"github.com/CptIdea/go-vk-api/httputil"
)

type longpoll struct {
	Response map[string]string
}

func (vk *Session) UpdateCheck(GroupId int) Updates {
	var upd Updates
	for len(upd.Updates) == 0 {
		URL := "https://api.vk.com/method/groups.getLongPollServer?group_id=" + strconv.Itoa(GroupId) + "&v=" + vk.Version + "&access_token=" + vk.Token
		resp, err := http.Get(URL)
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
		URL = jsoned.Response["server"] + "?act=a_check&key=" + jsoned.Response["key"] + "&ts=" + jsoned.Response["ts"] + "&wait=25"
		resp, err = http.Get(URL)
		if err != nil {
			fmt.Println("error:", err)
			fmt.Println(string(bs[:n]))

		}
		bs = make([]byte, 8112)
		n, err = resp.Body.Read(bs)
		if err != nil {
			fmt.Println("error:", err)
		}
		err = json.Unmarshal(bs[:n], &upd)
		if err != nil {
			fmt.Println("error:", err)
		}
	}
	return upd
}

// Отправляет сообщение message в чат ToId
func (vk *Session) SendMessage(ToId int, message string) []byte {
	return vk.SendRequest("messages.send", Request{"peer_id": strconv.Itoa(ToId), "message": message,"random_id":rand.Intn(2147483647)})
}
func (vk *Session) EditMessage(Peer, MessageId int, NewMessage string) []byte {
	return vk.SendRequest("messages.edit", Request{"peer_id": Peer, "message": NewMessage, "message_id": MessageId})
}
func (vk *Session) SendKeyboard(ToId int, Keyboard Keyboard, Message string, Attachments ...string) []byte {
	ReadyKeyboard := map[string]interface{}{
		"one_time": Keyboard.OneTime,
		"inline":   Keyboard.Inline,
	}

	var ButtonsArray [][]map[string]interface{}

	for _, k := range Keyboard.Buttons {
		var ButtonString []map[string]interface{}
		for _, c := range k {
			CurrentButton := make(map[string]interface{})
			if c.Color != "" {
				CurrentButton["color"] = c.Color
			}
			switch c.Action.Type {
			case "text":
				CurrentButton["action"] = map[string]interface{}{
					"type":    "text",
					"label":   c.Action.Label,
					"payload": c.Action.Payload,
				}
			case "open_link":
				CurrentButton["action"] = map[string]interface{}{
					"type":    "open_link",
					"link":    c.Action.Link,
					"label":   c.Action.Label,
					"payload": c.Action.Payload,
				}
			case "location":
				CurrentButton["action"] = map[string]interface{}{
					"type":    "location",
					"payload": c.Action.Payload,
				}
			case "vkpay":
				CurrentButton["action"] = map[string]interface{}{
					"type":    "open_link",
					"payload": c.Action.Payload,
					"hash":    c.Action.Hash,
				}
			case "open_app":
				CurrentButton["action"] = map[string]interface{}{
					"type":     "open_app",
					"app_id":   c.Action.AppId,
					"owner_id": c.Action.OwnerId,
					"payload":  c.Action.Payload,
					"label":    c.Action.Label,
					"hash":     c.Action.Hash,
				}
			default:
				continue
			}

			ButtonString = append(ButtonString, CurrentButton)
		}
		if ButtonString != nil {
			ButtonsArray = append(ButtonsArray, ButtonString)
		}

	}
	ReadyKeyboard["buttons"] = ButtonsArray
	JSONKey, err := json.Marshal(ReadyKeyboard)
	if err != nil {
		fmt.Println(err)
	}
	var resp []byte
	if len(Attachments) > 0 {
		var StrAtt string
		for i, cur := range Attachments {
			StrAtt += cur
			i++
			if i < len(Attachments) {
				StrAtt += ","
			}
		}
		resp = vk.SendRequest("messages.send", Request{"peer_id": ToId, "message": Message, "keyboard": string(JSONKey), "attachments": StrAtt,"random_id":rand.Intn(2147483647)})
	} else {
		resp = vk.SendRequest("messages.send", Request{"peer_id": ToId, "message": Message, "keyboard": string(JSONKey),"random_id":rand.Intn(2147483647)})
	}

	return resp
}

func (vk *Session) GetConversationInfo(peers []int, GroupId int, fields ...string) []Conversation {
	var params string
	var output struct {
		Response struct{ Items []Conversation }
	}
	var Peers, Fields string
	params = "peer_ids="
	for i := 0; i < len(peers); {
		Peers += strconv.Itoa(peers[i])
		i++
		if i < len(peers) {
			Peers += ","
		}
	}
	if len(fields) != 0 {
		for i := 0; i < len(fields); {
			Fields += fields[i]
			i++
			if i < len(fields) {
				Fields += ","
			}
		}
	}
	params = params + "&GroupId=" + strconv.Itoa(GroupId)
	resp := vk.SendRequest("messages.getConversationsById", Request{"peer_ids": Peers, "fields": Fields})
	fmt.Println(string(resp))
	err := json.Unmarshal(resp, &output)
	if err != nil {
		fmt.Println("error:", err)
	}
	return output.Response.Items
}

//Возвращает информацию о пользователях в массиве объектов User
func (vk *Session) GetUsersInfo(Ids []int, fields ...string) []User {
	var output struct{ Response []User }
	var Users, Fields string
	for i := 0; i < len(Ids); {
		Users += strconv.Itoa(Ids[i])
		i++
		if i < len(Ids) {
			Users += ","
		}
	}
	for i := 0; i < len(fields); {
		Fields += fields[i]
		i++
		if i < len(fields) {
			Fields += ","
		}
	}
	resp := vk.SendRequest("users.get", Request{"user_ids": Users, "fields": Fields})
	err := json.Unmarshal(resp, &output)
	if err != nil {
		fmt.Println("error:", err)
	}
	return output.Response
}

//Возвращает информацию о сообществах в массиве объектов Group
func (vk *Session) GroupGetById(GroupIds []int, fields ...string) []Group {
	var output struct{ Response []Group }
	var Groups string
	var Fields string
	for i := 0; i < len(GroupIds); i++ {
		GroupIds[i] = GroupIds[i] * -1
	}
	for i := 0; i < len(GroupIds); {
		Groups += strconv.Itoa(GroupIds[i])
		i++
		if i < len(GroupIds) {
			Groups += ","
		}
	}
	for i := 0; i < len(fields); {
		Fields += fields[i]
		i++
		if i < len(fields) {
			Fields += ","
		}
	}
	resp := vk.SendRequest("groups.getById", Request{"group_ids": Groups, "fields": Fields})
	err := json.Unmarshal(resp, &output)
	if err != nil {
		fmt.Println("error:", err)
	}
	return output.Response
}

func (vk *Session) SendRequest(method string, params Request) []byte {
	Url := "https://api.vk.com/method/" + method + "?" + "v=" + vk.Version + "&access_token=" + vk.Token
	ReadyParams := url.Values{}
	for k, v := range params {
		ReadyParams.Add(k, fmt.Sprint(v))
	}
	var resp []byte
	var err interface{}
	for {
		resp, err = httputil.Post(http.DefaultClient, Url, ReadyParams)
		if err != nil {
			log.Fatalln(err)
		}
		if !vk.SkipAutoResend {
			if strings.Contains(string(resp), "Too many requests per second") || strings.Contains(string(resp), "400 Bad Request") {
				time.Sleep(time.Second)
				continue
			}
		}
		break
	}
	return resp
}
