# Mini Blog API

A basic API for a blog page.

## Used dependencies & Credits

- [MongoDB Driver](https://go.mongodb.org/mongo-driver/mongo)
- [Viper](https://github.com/spf13/viper) for .env files.
- [Chi-router](https://github.com/go-chi/chi)
- [Validation](https://github.com/go-playground/validator)

## Usage

### .env file

All variables in the `.env` file are optional.

| variable         | description                                                       | default                     | default-.env            |
| ---------------- | ----------------------------------------------------------------- | --------------------------- | ----------------------- |
| PORT             | the port number this server will listen on                        | `4000`                      | `4000`                  |
| ENVIRONMENT      | has no impact whatsoever yet but can be seen in the server status | `development`               | `development`           |
| DSN              | mongodb connection string                                         | `mongodb://localhost:27017` | `mongodb://db:27017`    |
| JWT_TOKEN_SECRET | secret phrase for encrypting the passwords                        | `wonderfulsecretphrase`     | `wonderfulsecretphrase` |

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

## Contributing

Even though this project is made for private learning purposes I would never decline recommendations for improvements.

## TODO

- validation (done for now)
- login/logout done
- pagination (might want to revisit later for a "better" response)
- graphql?! why is this a thing?! D:
  - still cba'd

## License

This project is licensed under the [MIT](LICENSE) License.
