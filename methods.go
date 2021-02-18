package vk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type longpoll struct {
	Response map[string]string
}

func (vk *Session) UpdateCheck(GroupId int) (Updates, error) {
	var upd Updates
	for len(upd.Updates) == 0 {
		URL := fmt.Sprintf("https://api.vk.com/method/groups.getLongPollServer?group_id=%s&v=%s&access_token=%s", strconv.Itoa(GroupId), vk.version, vk.token)
		resp, err := http.Get(URL)
		if err != nil {
			return Updates{}, err
		}
		rawAns, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Updates{}, err
		}
		var jsoned longpoll
		err = json.Unmarshal(rawAns, &jsoned)
		if err != nil {
			return Updates{}, err
		}

		URL = fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=25", jsoned.Response["server"], jsoned.Response["key"], jsoned.Response["ts"])
		resp, err = http.Get(URL)
		if err != nil {
			return Updates{}, fmt.Errorf("err:%e\tdata:%s", err, string(rawAns))
		}
		rawAns, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return Updates{}, err
		}
		err = json.Unmarshal(rawAns, &upd)
		if err != nil {
			return Updates{}, err
		}
	}
	return upd, nil
}

// Отправляет сообщение message в чат ToId
func (vk *Session) SendMessage(ToId int, message string) ([]byte, error) {
	return vk.SendRequest("messages.send", Request{"peer_id": strconv.Itoa(ToId), "message": message, "random_id": rand.Intn(2147483647)})
}

// Редактирует сообщение messageId
func (vk *Session) EditMessage(Peer, MessageId int, NewMessage string) ([]byte, error) {
	return vk.SendRequest("messages.edit", Request{"peer_id": Peer, "message": NewMessage, "message_id": MessageId})
}

func (vk Session) KeyGet(key string, userId int) (string, error) {
	ans := struct {
		Response []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"response"`
	}{}

	resp, err := vk.SendRequest("storage.get", Request{"key": key, "user_id": userId})
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(resp, &ans)
	if err != nil {
		return "", err
	}
	return ans.Response[0].Value, err
}
func (vk Session) KeysGet(keys []string, userId int) ([]struct {
	Key   string
	Value string
}, error) {
	ans := struct {
		Response []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"response"`
	}{}

	resp, err := vk.SendRequest("storage.get", Request{"keys": strings.ReplaceAll(strings.Replace(strings.Replace(fmt.Sprintf("%s", keys), "[", "", 1), "]", "", 1), " ", ","), "user_id": userId})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &ans)
	if err != nil {
		return nil, err
	}
	return []struct {
		Key   string
		Value string
	}(ans.Response), err
}

func (vk Session) KeySet(key string, value string, userID int) ([]byte, error) {
	return vk.SendRequest("storage.set", map[string]interface{}{"key": key, "value": value, "user_id": userID})
}

func (vk *Session) SendKeyboard(ToId int, Keyboard Keyboard, Message string, Attachments ...string) ([]byte, error) {
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
		return nil, err
	}
	var resp []byte

	JSONKey = bytes.Replace(JSONKey, []byte("\"buttons\":null"), []byte("\"buttons\":[]"), 1)

	if len(Attachments) > 0 {
		var StrAtt string
		for i, cur := range Attachments {
			StrAtt += cur
			i++
			if i < len(Attachments) {
				StrAtt += ","
			}
		}
		resp, err = vk.SendRequest("messages.send", Request{"peer_id": ToId, "message": Message, "keyboard": string(JSONKey), "attachments": StrAtt, "random_id": rand.Intn(2147483647)})
	} else {
		resp, err = vk.SendRequest("messages.send", Request{"peer_id": ToId, "message": Message, "keyboard": string(JSONKey), "random_id": rand.Intn(2147483647)})
	}

	return resp, err
}

func (vk *Session) GetConversationInfo(peers []int, GroupId int, fields ...string) ([]Conversation, error) {
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
	resp, err := vk.SendRequest("messages.getConversationsById", Request{"peer_ids": Peers, "fields": Fields})

	err = json.Unmarshal(resp, &output)
	if err != nil {
		return nil, err
	}
	return output.Response.Items, err
}

//Возвращает информацию о пользователях в массиве объектов User
func (vk *Session) GetUsersInfo(Ids []int, fields ...string) ([]User, error) {
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
	resp, err := vk.SendRequest("users.get", Request{"user_ids": Users, "fields": Fields})
	err = json.Unmarshal(resp, &output)
	if err != nil {
		return nil, err
	}

	return output.Response, err
}

//Возвращает информацию о сообществах в массиве объектов Group
func (vk *Session) GroupGetById(GroupIds []int, fields ...string) ([]Group, error) {
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
	resp, err := vk.SendRequest("groups.getById", Request{"group_ids": Groups, "fields": Fields})
	err = json.Unmarshal(resp, &output)
	if err != nil {
		return nil, err
	}
	return output.Response, err
}

func (vk *Session) SendRequest(method string, params Request) (resp []byte, err error) {
	URL := fmt.Sprintf("https://api.vk.com/method/%s?v=%s&access_token=%s", method, vk.version, vk.token)

	ReadyParams := url.Values{}
	for k, v := range params {
		ReadyParams.Add(k, fmt.Sprint(v))
	}
	for {
		resp, err = Post(http.DefaultClient, URL, ReadyParams)
		if err != nil {
			return nil, err
		}
		vkErr := vkError{}
		err := json.Unmarshal(resp, &vkErr)
		if err != nil {
			return nil, err
		}

		if !vk.SkipAutoResend {
			if strings.Contains(string(resp), "Too many requests per second") || strings.Contains(string(resp), "400 Bad Request") {
				time.Sleep(time.Second / 2)
				continue
			}
		}
		if vkErr.Error.ErrorCode != 0 {
			return nil, fmt.Errorf("code:%d\terror:%s\tparams:%v\trawResp:%s", vkErr.Error.ErrorCode, vkErr.Error.ErrorMsg, vkErr.Error.RequestParams, resp)
		}
		break
	}
	return resp, err
}
