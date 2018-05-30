## Description
課題1のアプリケーションを拡張し、リレーショナルデータベースを用いて、 RESTful な Web Application を作成

##URL
http://localhost:8080/
##### 使用したdocker image
```docker
golang:latest
```
```docker
postgres:latest
```
## Usage

### build and run
```docker:
# docker-compose
$ git clone https://github.com/aeleniumfor/coding_challenge_2.git
$ cd coding_challenge_2
$ docker-compose build
$ docker-compose up
```
バックグラウンドで実行する場合
```
$ docker-compose up -d
```
### confirmation

##### ステータスコードの確認
```
$ curl -LI http://localhost:8080/ -o /dev/null -w '%{http_code}' -s
```

##### レスポンスの確認
```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```
##### user を作成
```
curl -XPOST -H 'Content-Type:application/json' http://localhost:8080/users -d '{"name": "test", "email": "hoge@example.com" }'
```

##### user を更新
```
$ curl -XPUT -H 'Content-Type:application/json' http://localhost:8080/users/1 -d '{"name": "koudaiii", "email": "hoge@example.com" }'

```
##### user を確認
```
curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users/1
```
##### user を一覧
```
curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users
```
##### user を削除
```
$ curl -XDELETE -H 'Content-Type:application/json' http://localhost:8080/users/1
```
