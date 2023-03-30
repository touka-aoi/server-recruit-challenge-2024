# サーバーエンジニア向け 2024新卒採用事前課題

あなたは歌手とアルバムを管理するAPIの機能開発にたずさわることになりました。

次の課題に順に取り組んでください。

できない課題があっても構いません。

面接中に課題に関して質問をしますので、分かる範囲で説明してください。

## 課題1
プログラムのコードを読み、中身を把握しましょう。

## 課題2
go をインストールし(各自で調べてください)、歌手を管理するAPIの動作を確認しましょう。

```
# (ターミナルを開いて)
# サーバーを起動する
go run main.go
```

```
# (別のターミナルを開いて)
# 歌手の一覧を取得する
curl http://localhost:8888/singers

# レスポンス
[{"id":1,"name":"Alice"},{"id":2,"name":"Bella"},{"id":3,"name":"Chris"},{"id":4,"name":"Daisy"},{"id":5,"name":"Ellen"}]

# 指定したIDの歌手を取得する
curl http://localhost:8888/singers/1

# レスポンス
{"id":1,"name":"Alice"}

# 歌手を追加する
curl -X POST -d '{"id":10,"name":"John"}' http://localhost:8888/singers

# レスポンス
{"id":10,"name":"John"}

# 歌手を削除する
curl -X DELETE http://localhost:8888/singers/1
```

## 課題3
アルバムを管理するAPIを新規作成しましょう。

### 3-1
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer_id":1},{"id":2,"title":"Alice's 2nd Album","singer_id":1},{"id":3,"title":"Bella's 1st Album","singer_id":2}]
```

### 3-2
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer_id":1}
```

### 3-3
アルバムを追加するAPI
```
curl -X POST -d '{"id":10,"title":"Chris 1st","singer_id":3}' http://localhost:8888/albums

# このようなレスポンスを期待しています
{"id":10,"title":"Chris 1st","singer_id":3}

# そして、アルバムを取得するAPIでは、追加したものが存在するように
curl http://localhost:8888/albums/10
```

### 3-4
アルバムを削除するAPI
```
curl -X DELETE http://localhost:8888/albums/1
```

## 課題4
アルバムを取得するAPIでは、歌手の情報も付加するように改修しましょう。

### 4-1
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}}
```

### 4-2
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}},{"id":2,"title":"Alice's 2nd Album","singer":{"id":1,"name":"Alice"}},{"id":3,"title":"Bella's 1st Album","singer":{"id":2,"name":"Bella"}}]
```

## 課題5
歌手とそのアルバムを管理するという点で、現状のAPIの改善点を検討し思いつく限り書き出してください。

実装をする必要はありません。
    

1. [Gorilla Web Toolkit (github.com)](https://github.com/gorilla#gorilla-toolkit)によるとgorilla/muxはメンテナンスされなくなってしまったので、別のwebフレームワークに移行したほうが良い。
2. GetALL関数が大量のデータを取得してしまい、アプリケーションのパフォーマンスが悪くなる可能性がある。NからN+Mまでのような制限を設けたほうが良い。実際にパフォーマンスが悪くなったら対処すればよい問題ではあるが。
3. 同姓同名の歌手の判別か人間にはつかないので、歌手のクエリにアルバム情報などを添付したい。
4. POSTによるアーティスト情報の上書きが容易にできてしまうことが明示的ではない。関数名などを工夫して上書きをするPOSTだということを明示的にしたい。
5. データのクエリにdeffer_recoverを設けていないので、不明なエラーが起きた場合サーバーが終了してしまう可能性があるので、例外処理を追加することで対処したい。
6. リクエストのタイプエラーが起きたときに、model.AlbumIDなどのgoの自作型情報が返ってくる。これではどの型が正しい方なのかがわからないので、intやstringなどのプリミティブな型情報を提供するようにしたい。
7. Getで存在しないIDを指定した際、500 InternalServiceErrorが返ってくるが、404 Not Foundを返したい。
8. 7で404 Not Foundを返したいといったが、GetALLで一つでも存在しないsingerIDがあるとすべてのクエリが死んでしまうので、404が出た場合はnullパターンを使用して、空データを返すようにして、エラーで止まることを防ぎたい。
9. しかしながら、singerIDが存在しないクエリを弾くようにすることによって、8のパターンを除去できるのでPOSTの時に対処を行ってもよい。
10. GetALLの時、以前クエリした歌手IDだったとしてももう一度クエリを行うため、GetALL中は歌手IDの結果をキャッシュして、クエリの回数を減らしたい。メモリと時間との兼ね合いになると思うが。
11. 誰でも管理情報にアクセスできてしまうため、認証を追加して管理情報を保護した方が良い。

