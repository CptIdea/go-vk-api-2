# vk_api Verison 2
--------
 Golang api for vk
 
 Включает в себя все объекты, описанные в vk.com/dev. Добавлять все функции смысла не вижу, они без проблем вызываются `Session.SendRequest`

Зачем V2?
---------------
В прошлой версии библиотеки полностью отсутствовала проверка ошибок, из-за чего могли происходить внезапные необработанные краши.
# Использование
---------------
 Все работает через объект `Session`. 
 
 `Session.UpdateCheck` Проверяет longpoll сервер группы, переданной в параметрах функции
 
 `Session.SendRequest` Отправляет запрос `method` с параметрами `params`. По умолчанию автоматически повторяет запрос если превышен лимит запросов в секунду. Отключается через параметр `Session.SkipAutoResend`.
 
 # Клавиатуры
 -----------
 Эта часть может вызвать вопросы. Вот пример кода
 
 ```golang
YNKeyboard := vk.Keyboard{
		OneTime: true,
		Inline:  false,
		Buttons: [][]vk.Button{make([]vk.Button, 2)},
	}
	YNKeyboard.Buttons[0][0].Action.Type = "text"
	YNKeyboard.Buttons[0][0].Action.Label = "Да"
	YNKeyboard.Buttons[0][0].Color = "secondary"
	YNKeyboard.Buttons[0][1].Action.Type = "text"
	YNKeyboard.Buttons[0][1].Action.Label = "Нет"
	YNKeyboard.Buttons[0][1].Color = "secondary"
	botSession.SendKeyboard(ID, YNKeyboard, "Да или нет?")
 ```
 
