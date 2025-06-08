## SSO (Auth, Permissions, User Info)
В данном репозитории представлено построение:  
- gRPC сервера
- авторизации (Auth)
- работы с пермишеннами (Permissions)
- предоставления информации о пользователе (User Info)

На данном этапе реализуется авторизация. 

## Building & Running

- Для генерации protobuf файлов в папке `protos/gen/go/sso` достаточно вызвать команду в корне проекта:
`task generate`

- Запуск приложения:
`go run cmd/sso/main.go --config=./config/local.yaml`
