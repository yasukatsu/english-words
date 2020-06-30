# english-words
英単語学習アプリケーション  
ソース元：http://www.newgeneralservicelist.org.

## 初期設定
`server/config`配下に、`.env.local.example`を参考にして`.env.local`ファイルを作成する。（値が明記されてない項目は任意の値を入れる）

## 起動方法
```
$ make up
```

## 停止方法
```
^(control) C
$ make down
```

## xorm reverse の使い方

[Github](https://gitea.com/xorm/reverse)  

DBのスキーマが変更された際に実行する必要がある。

`server/internal/pkg/postgres`配下に、`xorm-reverse.yml.example`を参考にして`xorm-reverse.yml`を作成する。（conn_strの値は`.env.local`で設定した値を入れる）

作成した後に、下記コマンドを実行
```
$ reverse -f internal/pkg/db/postgres/xorm-reverse.yml
```
`server/internal/pkg/postgres`配下に`models.go`が作成される。
