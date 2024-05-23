# social-media-api

Система для добавления и чтения постов и комментариев с
использованием GraphQL на Go.

### Сборка и запуск

#### Требования к системе

- Docker 19.03.0+
- Docker Compose 1.25.0+
- Make

#### Установка

```shell
git clone https://github.com/bitbox228/social-media-api
```

#### Конфигурация

Перед запуском приложения, нужно прописать конфиг и .env файл. Примеры, которые можно использовать:

`.env`
```dotenv
PG_USER=postgres
PG_PASSWORD=qwerty
PG_DB=postgres
PG_URL=postgres://postgres:qwerty@postgres:5432/postgres?sslmode=disable
```

`config/config.yaml`
```yml
http:
  port: '8000'

postgres:
  username: 'postgres'
  host: 'postgres'
  port: '5432'
  dbname: 'postgres'
  sslmode: 'disable'
  poolmax: 5

repository:
  type: 'postgres'
```

При конфигурации можно выбрать тип хранения данных. Тип `postgres` для хранения данных в БД PostgreSQL.
```yml
repository:
  type: 'postgres'
```

Для In-Memory хранилища:
```yml
repository:
  type: 'inmemory'
```

#### Запуск

1. Удостоверьтесь, что файл `.env` лежит в корне проекта и содержит нужные переменные среды (можно использовать пример выше).
    ```shell
    touch .env
    echo "PG_USER=postgres" >> .env
    echo "PG_PASSWORD=qwerty" >> .env
    echo "PG_DB=postgres" >> .env
    echo "PG_URL=postgres://postgres:qwerty@postgres:5432/postgres?sslmode=disable" >> .env
    ```
2. Соберите приложение:
    ```shell
    make build
    ```
   
3. Запустите приложение:
    ```shell
    make up
    ```
4. Для остановки приложения напишите:
    ```shell
    make down
    ```

### Использование

Приложение предоставляет GraphQL ручку, а так же GraphQL Playground для удобного тестирования.

- GraphQL ручка: http://localhost:8000/graphql
- GraphQL Playground: http://localhost:8000/

Спецификация API есть на Graphql Playground.

#### Примеры запросов

##### Query

Доступны для получения списком или по одному посты и комментарии. Для списков доступны параметры для пагинации а так же фильтры, например фильтр по UserID для постов.

```graphql
query Comments {
    comments(filter: { postID: "5" }, limit: 10, offset: 0) {
        id
        userID
        postID
        parentID
        body
    }
}
```

В запросе выше выводится список комментариев к посту с ID = 5, при этом выводятся первые 10 элементов, начиная с 0-го.

##### Mutations

Доступны мутации по созданию постов, комментариев к постам и реплаев к комментариям, а также отключение/включение комментариев к постам.

```graphql
mutation CreateComment {
    createComment(
        input: { userID: "5", postID: "10", body: "Hello, I'm new here >.<" }
    ) {
        id
        body
    }
}
```

В запросе выше создается коментарий от пользователя с ID = 5 к посту с ID = 10, возвращает мутация ID и тело комментария.

##### Subscriptions

Есть возможность подписаться на обновления поста для получения новых комментариев под этим постом. 

```graphql
subscription {
  newComment(userID: 1 postID: 1) {
    id
    body
  }
}
```

Пользователь может подписаться на конкретный пост только один раз, поэтому нужно указывать ID пользователя и поста. 