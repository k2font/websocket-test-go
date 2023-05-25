# websocket-test-go
## Features
- 簡易チャットアプリケーション。メッセージを入力して送信すると、相手にメッセージがリアルタイムに届きます。
  - GoでWebSocketサーバの手慣らし
- melodyを採用してコーディング
  - `net/http`よりコーディングが簡単でコードがシンプルになるため


## Requirements
- olahol/melody@1.1.3
- rs/zerolog@1.29.1

## Usage
- `$ go run path/to/main.go`
- http:localhost:5001 にアクセス