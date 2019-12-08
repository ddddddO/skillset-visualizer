# ref: https://jp.vuejs.org/v2/cookbook/dockerize-vuejs-app.html
FROM node:lts-alpine

# 静的コンテンツを配信するシンプルな http サーバをインストールする
RUN npm install -g http-server

# カレントワーキングディレクトリとして 'app' フォルダを指定する
WORKDIR /app

# `package.json` と `package-lock.json` （あれば）を両方コピーする
COPY app/package*.json ./

# プロジェクトの依存ライブラリをインストールする
RUN npm install

# カレントワーキングディレクトリ(つまり 'app' フォルダ)にプロジェクトのファイルやフォルダをコピーする
COPY ./app .

# 本番向けに圧縮しながらアプリケーションをビルドする
RUN npm run build

EXPOSE 8080
CMD [ "http-server", "dist" ]
