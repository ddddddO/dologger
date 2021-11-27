ログをTCP通信で別プロセスへ送るサンプル

# Usage

1. ログを受信するプロセスを起動
```console
go run receive_log/main.go
```

2. 別ターミナルでログを一定間隔で出力するプロセスを起動
```console
go run send_log/main.go
```

# References

- https://zenn.dev/hsaki/books/golang-io-package/viewer/netconn
