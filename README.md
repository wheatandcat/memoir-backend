# memoir-backend

![coverage](docs/coverage.svg)

## ローカル実行

### 導入

```
$ go install github.com/go-delve/delve/cmd/dlv@latest
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

## GraphQL スキーマ更新

```zsh
$ go run github.com/99designs/gqlgen generate
```

## ユニットテスト

### 実行

moq を生成

```zsh
$ make moqgen
```

テストを実行

```zsh
$ go test -race ./...
```

### カバレッジ表示

### インストール

```zsh
$ brew install k1LoW/tap/octocov
```

### 実行

```zsh
$ go test ./... -coverprofile=coverage.out
$ octocov
```

### 各ファイルカバレッジ確認

```zsh
$ octocov ls-files
```

```zsh
$ octocov view graph/invite.go
```

## E2E テスト

```zsh
$ FIRESTORE_EMULATOR_HOST=localhost:3600 air
```

```zsh
$ cd e2e
$ make create_login_yaml
$ make local_scenarigo
```

## 環境変数を更新

```zsh
$ gcloud run services describe SERVICE --format export > service.yaml
# 環境変数を追加して以下を実行
$ gcloud run services replace service.yaml
```

## 本番デプロイ

```zsh
$ git checkout main
$ git pull --ff-only origin main
$ git tag -a v1.0.0 -m 'リリース内容'
$ git push origin v1.0.0
```

## CI 環境

### Firebase トークン

```zsh
$ firebase login:ci
```

### レビュー環境

```zsh
$ base64 -i serviceAccount.review.json | pbcopy
```

```zsh
$ base64 -i gcpServiceAccount.review.json | pbcopy
```

### E2E

```zsh
$ base64 -i .env | pbcopy
```

```zsh
$ base64 -i envenb.go | pbcopy
```

```zsh
$ base64 -i e2e/.env | pbcopy
```

### 本番環境

```zsh
$ base64 -i serviceAccount.production.json | pbcopy
```

```zsh
$ base64 -i gcpServiceAccount.production.json | pbcopy
```
