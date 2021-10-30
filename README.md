# go-go-file

## Описание
Этот проект - обертка над API сервиса gofile - https://gofile.io/
Сервис позволяет загружать и хранить файлы

## Авторизация
Для использования нужен API ключ. Его можно получить при аутентификации в сервисе:<br>
1. Зайдите на страницу - https://gofile.io/myProfile
1. Нажмите - "Login with your email"
1. Получите письмо на почту с ссылкой для входа
1. Перейдите по ссылке
1. Раздел "My Profile"
1. Скопируйте "Api token"

## Возможности
1. **Загрузить файл** <br>
Загрузите файл при помощи метода UploadFile. Метод принимает путь к файлу и при успешной загрузке возвращает ссылку на скачивание файла.
1. **Информация об аккаунте** <br>
Метод GetAccountDetails позволяет проверить информацию об аккаунте. Метод выводит информацию в консоль.

## Пример использования
```go
package main

import (
	"fmt"
	"github.com/mikeyuniverse/go-go-file"
	"log"
)

const TOKEN = "yourToken"

func main() {
	client := gofile.NewClient(TOKEN)

	downloadLink, err := client.UploadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("File uploaded\nDOWNLOAD URL -", downloadLink)

	account, err := client.GetAccountDetails()
	if err != nil {
		log.Fatal(err)
		return
	}
	account.Info()
}
```