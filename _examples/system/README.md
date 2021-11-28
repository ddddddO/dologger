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

5. Redashを起動

https://redash.io/help/open-source/dev-guide/docker を１度、上から「Run webpack Dev Server」まで実施する。この手順だと、ブラウザからのアクセス時に以下エラーがでる。

```
FileNotFoundError
FileNotFoundError: [Errno 2] No such file or directory: '/app/redash/settings/../../client/dist/index.html'
```

なので、/redash/clientディレクトリ下で、以下を実行しdistディレクトリを生成する。

```sh
npm install --no-optional
npm run build
```

上の手順を実施した後、再度redashを動かす場合は、以下を実行する。

```sh
cd ~/work/redash
docker-compose up -d
open http://localhost:5000/
```

そして、redashをブラウザで開いた後、Elastic Search用のData Sourceを作成する。
この時、Base URLには、`http://host.docker.internal:9200` を設定する。

redashから以下クエリを実行すると、送出したログの一覧がテーブル形式で確認できる。
```json
{
    "index": "fluent_bit_out"
}
```

ElasticsearchもSQLチックにクエリ出来ると思ってたけど違うっぽい。

終わったらdown
```sh
cd ~/work/redash
docker-compose down
```

# References

- https://github.com/ddddddO/dologger/issues/3#issue-1064890682
