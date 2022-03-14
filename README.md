# Mini Blog API

A basic API for a blog page.

## Used dependencies & Credits

- [MongoDB Driver](https://go.mongodb.org/mongo-driver/mongo)
- [Viper](https://github.com/spf13/viper) for .env files.
- [Chi-router](https://github.com/go-chi/chi)
- [Validation](https://github.com/go-playground/validator)

## Installation

### .env file

All variables in the `.env` file are optional.

| variable             | description                                                       | default                     | default-.env            |
| -------------------- | ----------------------------------------------------------------- | --------------------------- | ----------------------- |
| PORT                 | the port number this server will listen on                        | `4000`                      | `4000`                  |
| ENVIRONMENT          | has no impact whatsoever yet but can be seen in the server status | `development`               | `development`           |
| DSN                  | mongodb connection string                                         | `mongodb://localhost:27017` | `mongodb://db:27017`    |
| JWT_SECRET           | secret phrase for encrypting the passwords                        | `wonderfulsecretphrase`     | `wonderfulsecretphrase` |
| CORS_ALLOWED_ORIGINS | allowed domains for CORS requests separated by spaces             | `http://* https://*`        | `http://* https://*`    |
| COOKIE_NAME          | cookie name which gets set in the browser                         | `uwu-blog-cookie`           | `uwu-blog-cookie`       |

### docker-compose

For testing you can use the provided docker-compose. This will spin up the API along side a mongodb and uses the .env file as defaults.
The api will then run on port 4000. The db will be run on port 27017.

### With existing mongodb

For this you just need to run the Dockerfile and adjust the `.env`-file in the root directory.

### From docker hub

Run:

> docker run schattenbrot/mini-blog-api

If you want to change the defaults of the .env file:

> docker run -v ./path/to/.env:./run/.env

## Usage

### Routes

`apiUrl` is the url to the api.

#### Get App status

Get request on:

> apiURL/

The app status gets returned as a json object:

```json
{
  "status": "Available",
  "up_since": "2022-02-17T18:35:17.894497612Z",
  "current_uptime": 5916280315,
  "environment": "development",
  "version": "1.0.0"
}
```

#### Posts

Base URL:

> apiURL/v1/posts[/option]

| REQUEST                   | option                     | middlewares                 | description                             |
| ------------------------- | -------------------------- | --------------------------- | --------------------------------------- |
| `GET`                     | `/`                        | -                           | Gets a list of all posts in a jsonarray |
| [`GET`](#get-paging)      | `/paging?limit=%X&page=%Y` | -                           | Gets a list of all posts by paging.     |
| [`GET`](#get-single-post) | `/{id}`                    | -                           | Gets a single post by its ID.           |
| `POST`                    | `/`                        | Auth                        | Adds a new POST                         |
| `PATCH`                   | `/{id}`                    | Auth & IsPostCreatorOrAdmin | Patches a single post by its ID.        |
| `DELETE`                  | `/{id}`                    | Auth & IsPostCreatorOrAdmin | Deletes a single post by its ID.        |

##### GET base

- Doesn't take arguments

Example Response:

> Status: 200 OK

> Header: Content-Type application/json

```json
[
  {
    "id": "62019c31ef131e8cd42847ab",
    "title": "title",
    "text": "this is the text",
    "created_at": "2022-02-07T22:24:49.869Z",
    "updated_at": "2022-02-07T22:24:49.869Z"
  }
]
```

If no document is found it will return status 200 and an empty array.

##### GET paging

- `%X` needs to be an integer
- `%Y` needs to be an integer

Example Response:

> Status: 200 OK

> Header: Content-Type application/json

```json
[
  {
    "id": "62019c31ef131e8cd42847ab",
    "title": "title",
    "text": "this is the text",
    "created_at": "2022-02-07T22:24:49.869Z",
    "updated_at": "2022-02-07T22:24:49.869Z"
  }
]
```

If no document is found it will return status 200 and an empty array.

##### GET single post

- Doesn't take any arguments

Example Response:

> Status: 200 OK

> Header: Content-Type application/json

```json
{
  "id": "62019c31ef131e8cd42847ab",
  "title": "title",
  "text": "this is the text",
  "created_at": "2022-02-07T22:24:49.869Z",
  "updated_at": "2022-02-07T22:24:49.869Z"
}
```

If no document is found it will return Status 404 Not Found.

##### POST a post

Example Request Body:

```json
{
  "name": "Username",
  "email": "email@email.com",
  "password": "12345aA;"
}
```

Example Response:

> Status: 201 Created

> Header: Content-Type application/json

```json
{
  "ok": true
}
```

#### Users

Base URL:

> apiURL/v1/users[/option]

| REQUEST  | option    | middlewares          | description               |
| -------- | --------- | -------------------- | ------------------------- |
| `GET`    | `/logout` | Auth                 | Logs a user out           |
| `GET`    | `/{id}`   | Auth                 | Gets a user by its ID.    |
| `POST`   | `/`       | -                    | Adds a new user           |
| `POST`   | `/login`  | -                    | Logs a user in            |
| `PATCH`  | `/{id}`   | Auth & IsUserOrAdmin | Patches a user by its ID. |
| `DELETE` | `/{id}`   | Auth & IsUserOrAdmin | Deletes a user by its ID. |

### Middlewares

#### Auth

Allows authenticated people with a valid jwt token to access the specificied path.

#### IsPostCreatorOrAdmin

Allows only the creator of the post or admin to modify and delete the post.

#### IsUserOrAdmin

Allows only the user himself or admin to modify and delete the user.

## Contributing

Even though this project is made for private learning purposes I would never decline recommendations for improvements.

## TODO

- pagination (might want to revisit later for a "better" response)
- graphql?! why is this a thing?! D:
  - still cba'd
- cba completing the readme for now D:

## License

This project is licensed under the [MIT](LICENSE) License.
