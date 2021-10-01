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
$ rm -rf ./repository/moq.go
$ moq -out=repository/moq.go ./repository ItemRepositoryInterface UserRepositoryInterface InviteRepositoryInterface RelationshipRequestInterface RelationshipInterface PushTokenRepositoryInterface CommonRepositoryInterface
```

テストを実行
```
$ go test -race ./...
```

## デプロイ

```
$ gcloud app deploy
```


## CI環境

### レビュー環境

```
$ base64 -i serviceAccount.review.json | pbcopy
```

```
$ base64 -i gcpServiceAccount.review.json | pbcopy
```

```
$ base64 -i app.yaml | pbcopy
```

### 本番環境

```
$ base64 -i serviceAccount.production.json | pbcopy
```

```
$ base64 -i gcpServiceAccount.production.json | pbcopy
```

```
$ base64 -i app.production.yaml | pbcopy
```

