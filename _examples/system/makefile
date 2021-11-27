# Elasticsearchを起動。 少し時間かかる
## ref: https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html#docker-cli-run-dev-mode
es:
	docker run -p 127.0.0.1:9200:9200 -p 127.0.0.1:9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.15.2

# Fluent Bitを in:tcp / out:stdout,Elasticsearch モードで起動
ES_IP_ON_DOCKER = $(shell docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' `docker ps | grep elasticsearch | cut -d' ' -f1`)
fbd:
	docker run -ti -p 5170:5170 fluent/fluent-bit:1.8 /fluent-bit/bin/fluent-bit -i tcp -o es://$(ES_IP_ON_DOCKER):9200/fluent_bit_out/_doc -m '*' -o stdout -m '*'

# Fluent Bitに一定間隔でログを送出するためだけのappを起動
app:
	go run app/main.go

.PHONY: app
