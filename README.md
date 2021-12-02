# messenger
_Запустить:_
```sh
docker-compose up -d
```
_Добавить нового пользователя_
```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username": "user_1"}' \
  http://localhost:9000/users/add
# где "user_1" - имя нового пользователя
```
_Создать новый чат между пользователями_
```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": [1, 2, 3]}' \
  http://localhost:9000/chats/add
# где 1, 2 и 3 - это id(integer) пользователей
# "chat_1" - название чата
```
_Отправить сообщение в чат от лица пользователя_
```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1, "author": 2, "text": "hi"}' \
  http://localhost:9000/messages/add
# где 1 b 2 - это id(integer) чата и пользователя соответственно
```
_Получить список чатов конкретного пользователя_
```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user": 3}' \
  http://localhost:9000/chats/get
# где 3 - это id(integer) пользователя
```
_Получить список сообщений в конкретном чате_
```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1}' \
  http://localhost:9000/messages/get
# где 1 - это id(integer) чата
```
