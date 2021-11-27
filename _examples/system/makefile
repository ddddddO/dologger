# Fluent Bitをtcpモードで起動
fbd:
	docker run -ti -p 5170:5170 fluent/fluent-bit:1.8 /fluent-bit/bin/fluent-bit -i tcp -o stdout

# Fluent Bitに一定間隔でログを送出するためだけのappを起動
app:
	go run app/main.go

.PHONY: app
