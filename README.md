# REST API Server wit Go

## Initialize project

This project uses [`dep`](https://github.com/golang/dep). After install `dep`, run the following command:

```sh
dep ensure
```

## Run

```sh
go run cmd/sample-server/main.go --port=8080
```

## Development

### Generate template

```sh
swagger generate server -f ./spec/user.yaml -A Sample
```

This project uses [`go-swagger`](https://github.com/go-swagger/go-swagger). To install it :

```sh
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

### Reference

JP

- [go-swagger で API 定義(swagger.yml)からサーバーのコードを生成する](https://qiita.com/croquette0212/items/f7a21608b1eec1446c1c)
- [Swagger の記法まとめ](https://qiita.com/rllllho/items/53a0023b32f4c0f8eabb)
