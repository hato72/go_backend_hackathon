# CookMeet(サーバー)
![cookmeet](https://github.com/hato72/go_backend_hackathon/assets/139688965/54235b01-2da0-491e-857c-18581b70b518)

## フロントエンド
https://github.com/hato72/CookMeet

## DB設計
https://free-casquette-dee.notion.site/d558148d80f742a4ac77c0bf76b4a2c9?pvs=4


## 実行方法(テスト環境)

```sh
.env.dev:

PORT=8080
POSTGRES_USER=
POSTGRES_PW=
POSTGRES_DB=
POSTGRES_PORT=5432
POSTGRES_HOST=db
SECRET=
GO_ENV=dev
API_DOMAIN=localhost
FE_URL=http://localhost:3000
```

.env.devをbackendディレクトリ直下に配置した後に以下を実行

<!-- ```sh
docker compose build

docker compose up

docker compose run --rm backend sh

go run src/migrate/migrate.go

go run src/main.go

``` -->

```sh
docker compose build

docker compose up
```

## 実行方法(本番環境)
GCPはうまくできずに断念した

今後は(https://render.com/)を使用する予定

## メモ
dbイメージ　postgres latest 

バックエンドイメージ　hackathon-backend latest

Docker Composeで作ったコンテナ、イメージ、ボリューム、ネットワークを一括削除：
docker-compose down -v --rmi local

postmanで確認する際はrouter.goの26行目をコメントアウトして27行目のコメントアウトを外してから実行

missing csrf ~~ というエラーが出たらX-CSRF-TOKENを設定する

cuisines/ではフォームから入力する
