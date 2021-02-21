package vk

import (
	"strconv"
	"strings"
)

// Генерирует пустую клавиатуру
// pattern устанавливает количество кнопок в каждой строке.
// Пример: "122" установит 1 кнопку в первой строке, 2 кнопки во второй и 2 в третьей.
func GenerateEmptyKeyBoard(pattern string) Keyboard {
	KB := Keyboard{}
	i := 0
	for _, len := range pattern {
		curInt, err := strconv.Atoi(string(len))
		if err != nil {
			continue
		}
		KB.Buttons = append(KB.Buttons, []Button{})
		for i := 0; i < curInt; i++ {
			KB.Buttons[i] = append(KB.Buttons[i], Button{
				Action: struct {
					Type    string
					Label   string
					Payload string
					Link    string
					Hash    string
					AppId   int `json:"app_id"`
					OwnerId int `json:"owner_id"`
				}{Type: "text", Label: "auto"},
				Color: "secondary",
			})
		}
		i++
	}
	return KB
}

//Format Raw:
//Cell 1,Cell2;
//Cell3(new line);
//!pPrimary,!sSecondary;
//!rRed(negative),!gGreen(positive)
func GenerateKeyBoard(rawKeyboard string, inline bool, oneTime bool) Keyboard {
	lines := strings.Split(rawKeyboard, ";")
	KB := Keyboard{}
	KB.Buttons = make([][]Button, len(lines))

	for i, line := range lines {
		line = strings.ReplaceAll(line, ", ", ",")
		cells := strings.Split(line, ",")
		KB.Buttons[i] = make([]Button, len(cells))
		for i2, cell := range cells {
			newButton := Button{
				Color: "secondary",
			}
			newButton.Action.Type = "text"
			if cell[0] == '!' {
				switch cell[1] {
				case 'r':
					newButton.Color = "negative"
				case 'g':
					newButton.Color = "positive"
				case 'p':
					newButton.Color = "primary"
				}
				cell = cell[2:]
			}
			newButton.Action.Label = cell
			KB.Buttons[i][i2] = newButton
		}
	}
	KB.Inline = inline
	KB.OneTime = oneTime
	return KB
}

func NewTextButton(color string, label string, payload string) Button {
	return Button{
		Action: ButtonAction{
			Type:    "text",
			Label:   label,
			Payload: payload,
		},
		Color: color,
	}
}

func NewOpenLinkButton(color string, link string, label string, payload string) Button {
	return Button{
		Action: ButtonAction{
			Type:    "open_link",
			Link:    link,
			Label:   label,
			Payload: payload,
		},
		Color: color,
	}
}
func NewLocationButton(color string, payload string) Button {
	return Button{
		Action: ButtonAction{
			Type:    "location",
			Payload: payload,
		},
		Color: color,
	}
}
func NewVKPayButton(color string, hash string, payload string) Button {
	return Button{
		Action: ButtonAction{
			Type:    "vkpay",
			Payload: payload,
			Hash:    hash,
		},
		Color: color,
	}
}

func NewOpenAppButton(color string, label string, appID int, ownerID int, hash string, payload string) Button {
	return Button{
		Action: ButtonAction{
			Type:    "open_app",
			Payload: payload,
			Label:   label,
			AppId:   appID,
			OwnerId: ownerID,
			Hash:    hash,
		},
		Color: color,
	}
}

func NewCallbackButton(color string, label string, payload string) Button {
	return Button{
		Action: ButtonAction{
			Type:    "callback",
			Label:   label,
			Payload: payload,
		},
		Color: color,
	}
}
