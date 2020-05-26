# shuffle-members
登録されたメンバーリストを任意の数でシャッフルするためのアプリケーション

## モデル
### Project
- シャッフルを行うためのメンバーを登録するための箱

### Member
- シャッフル対象の人を表す

### Tag
- Member に登録できるタグ
- シャッフルするときのロジックで使われ、同じタグを持つ人どうしは同じグループになりにくくなる
- 所属チームや世代など、なんでも良い

### ShuffleLogHead
- シャッフル結果のヘッダ
- シャッフルを実行した単位で、そのときのグループ数、グループあたりの人数を保持する

### ShuffleLogDetail
- シャッフル結果の詳細
- シャッフルされた Member の単位

## 環境構築

- Go1.13
- MYSQL5.7

1.  諸々 インストール

    - Go のインストール
    - mysql のインストール
    - docker のインストール

1.  GOPATH の設定

    - Go を使うためには、GOPATH を設定する必要があります
    - GOPATH 配下には import する外部のライブラリが入ります
    - アプリケーションコードを置くことも多いらしい
    - 以下を`.bash_profile`なりに書いておきます
      ```
      export GOPATH=$HOME/.go
      export PATH=$GOPATH/bin:$PATH
      ```
      - (GOPATH はどこでも良いのだけど、これがよくある設定らしい)

1.  リポジトリをクローンし、docker コンテナの起動

    - docker-compose.yml にアプリケーションとDBのコンテナの記述があり、下記コマンドでアプリケーションサーバーが立ち上がります
      ```
      $ git clone https://github.com/RyotaNakaya/shuffle-members.git
      $ cd shuffle-members
      $ docker-compose up
      ```
    - docker 内の 8080 ポートをローカルマシンの 8080 ポートにフォワーディングしているので、localhost でアクセスできます
    - ローカルにマウントする場合は、データベースを事前に作っておきます

## 開発環境のホットリロード

- [realize](https://github.com/oxequa/realize) を使用
- ローカルで go run する代わりに、`realize start` で実行すると自動でホットリロードされます

## test

- テストは以下のコマンドで実行できます。
- ローカルに `shuffle_members_test` のデータベースを作っておいてください
```sh
go test ./... -timeout 30s -cover
```
