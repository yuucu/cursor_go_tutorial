# Hello API

ユーザーが送信したhelloメッセージをそのまま返すシンプルなAPIです。

## 機能

- `POST /hello` - 送信されたメッセージをそのまま返す

## 技術スタック

- Go 1.24.3
- [ogen](https://github.com/ogen-go/ogen) - OpenAPI 3.0からのコード生成
- OpenAPI 3.0.3 仕様

## セットアップ

1. 依存関係のインストール
```bash
go mod tidy
```

2. サーバーの起動
```bash
go run main.go
```

または

```bash
go build -o hello-api && ./hello-api
```

サーバーは `http://localhost:1323` で起動します。

## API使用例

### リクエスト
```bash
curl -X POST http://localhost:1323/hello \
  -H "Content-Type: application/json" \
  -d '{"message": "hello"}'
```

### レスポンス
```json
{"message": "hello"}
```

## テスト

テストスクリプトを実行してAPIの動作を確認できます：

```bash
./test_api.sh
```

## ファイル構成

- `openapi.yaml` - API仕様書
- `main.go` - サーバーのエントリーポイント
- `handlers/` - ogenによって生成されたハンドラー
  - `hello_handler.go` - 実際のビジネスロジック
- `test_api.sh` - テストスクリプト

## 開発の流れ

1. OpenAPI仕様書 (`openapi.yaml`) を作成
2. ogenを使用してhandlerを自動生成
3. 生成されたhandlerインターフェースを実装
4. main.goでサーバーを起動

これにより、API仕様とコードの一貫性が保たれ、メンテナンスが容易になります。
