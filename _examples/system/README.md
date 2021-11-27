dologger (on Application) -> Fluent Bit -> Elasticsearch <- Redash のサンプルコード


# Usage

1. Elasticsearchを起動
```console
make es
```

2. Fluent Bitを in:tcp / out:stdout,Elasticsearch モードで起動

```console
make fbd
```

3. 別ターミナルでFluent Bitへログを送出するappを起動

```console
make app
```

4. ブラウザからElasticsearchにアクセスし、カウントが上がっていればログがElasticsearchに貯められている

```console
open http://localhost:9200/_cat/count/fluent_bit_out
```


# References

- https://github.com/ddddddO/dologger/issues/3#issue-1064890682
