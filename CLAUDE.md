# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## コードベースの概要

memoir-backendは、Go言語で作成されたGraphQLバックエンドアプリケーションです。

- GraphQLサーバー (gqlgen使用)
- Firebase/Firestore認証・データベース
- Google Cloud Platform上でデプロイ
- 個人の日記・思い出管理アプリのバックエンド

## アーキテクチャ

### ディレクトリ構成
- `graph/` - GraphQL関連 (schema, resolvers, generated code)
- `repository/` - データアクセス層 (Firestore操作)
- `usecase/` - ビジネスロジック層
- `client/` - 外部クライアント用インターフェース・モック
- `e2e/` - E2Eテスト (scenarigo使用)

### レイヤー構成
1. GraphQL Resolver (`graph/`)
2. Use Case (`usecase/`)  
3. Repository (`repository/`)
4. Firestore Database

## 開発コマンド

### セットアップ
```bash
make install
```

### 起動
```bash
air  # ホットリロード付きでサーバー起動
```

### テスト
```bash
go test ./...  # 全テスト実行
```

### モック生成 (テスト前に必要)
```bash
make moqgen
```

### GraphQLスキーマ更新
```bash
go run github.com/99designs/gqlgen generate
```

### リント・静的解析
```bash
make precommit  # goimports, fmt, vet, errcheck, staticcheck実行
make golangci   # golangci-lint実行
```

### カバレッジ
```bash
go test ./... -coverprofile=coverage.out
octocov
```

### E2Eテスト
```bash
# Firebaseエミュレータ起動
FIRESTORE_EMULATOR_HOST=localhost:3600 air

# E2Eテスト実行
cd e2e
make create_login_yaml
make local_scenarigo
```

## 重要な設定ファイル

- `gqlgen.yml` - GraphQL生成設定
- `Makefile` - 開発用コマンド定義
- `graph/schema.graphqls` - GraphQLスキーマ定義
- `e2e/scenarigo.yaml` - E2Eテスト設定

## データベース

Firestoreを使用。認証はFirebase Authenticationで実装。

## デプロイ

Google Cloud Runにデプロイ。タグをプッシュすることで自動デプロイ:
```bash
git tag -a v1.0.0 -m 'リリース内容'
git push origin v1.0.0
```