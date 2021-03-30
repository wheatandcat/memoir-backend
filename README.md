# memoir-backend

## ローカル実行

```
$ dev_appserver.py app.local.yaml
```

## GraphQL スキーマドキュメント

```
$ yarn graphql-markdown http://localhost:8080/query > schema.md
```

## GraphQLスキーマ更新

```
$ go run github.com/99designs/gqlgen generate
```

## テスト

```
$ go test -race ./...
```


## デプロイ

```
$ gcloud app deploy
```