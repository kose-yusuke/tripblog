# gocrud
Goを用いたWebアプリ作成

## 立ち上げ
1. docker desktopを起動する
2. ターミナルでターミナルでプロジェクトディレクトリに移動
3. docker compose up -dで起動
4. docker compose exec front shでフロント起動
5. yarn startでport3000にてlocalサイト起動
6. docker compose exec api shでバックエンド起動
7. cd api/cmdの後、go run main.goでAPIサーバー起動
8. http://localhost:8000/browser/ にてpgadimn起動

APIの起動は逐次Postmanで確認可能(エンドポイント:http://localhost:8080)

##サイトのRoute
calendarサイト:http://localhost:3000/home
todoサイト:http://localhost:3000/todo
api起動テストサイト:http://localhost:3000/user
login:http://localhost:3000/login

## ディレクトリ構成

## 技術
- React
- Go/Gin/GORM
- Postgresql
- Docker
- AWS ECS(未遂)

## 構成
<img width="623" alt="スクリーンショット 2022-09-02 11 03 53" src="https://user-images.githubusercontent.com/61424757/188043149-b22ad2e0-d676-494f-bebd-3af69d885c17.png">
