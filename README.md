## Description
Docker を用いて、 "Hello world" というメッセージ を JSON 形式で返す API 
- http://153.122.97.129/

##URL
http://153.122.97.129/
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
$ git clone https://github.com/aeleniumfor/coding_challenge_1.git
$ cd coding_challenge_1
$ docker build -t golang_app .
$ docker run -d -p 8080:8080 --rm golang_app
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


