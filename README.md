# memoir-backend

![coverage](docs/coverage.svg)

## ローカル実行

### 導入
```
$ go get -u github.com/cosmtrek/air
```

### 起動

```
$ air
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


### 実行
moqを生成
```
$ make moqgen
```

テストを実行
```
$ go test -race ./...
```

### カバレッジ表示

### インストール
```
$ brew install k1LoW/tap/octocov
```

### 実行
```
$ go test ./... -coverprofile=coverage.out
$ octocov
```

### 各ファイルカバレッジ確認
```
$ octocov ls-files
```

```
$ octocov view graph/invite.go
```

## 手動デプロイ

```
$ gcloud app deploy
```

## 本番デプロイ

```
$ git checkout main
$ git pull --ff-only origin main
$ git tag -a v1.0.0 -m 'リリース内容'
$ git push origin v1.0.0
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

