# vk_api
 Golang api for vk
# Использование
 Все работает через объект Session. 
 Session.UpdateCheck проверяет longpoll сервер группы, переданной в параметрах функции
 Session.SendRequest Отправляет запрос method с параметрами params. По умолчанию автоматически повторяет запрос если превышен лимит запросов в секунду. Отключается через параметр Session.SkipAutoResend.
 
