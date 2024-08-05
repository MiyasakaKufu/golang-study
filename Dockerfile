FROM golang:1.22.5-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Go Modules を使用するための環境変数を設定
ENV GO111MODULE=on

# 必要なパッケージをインストール
RUN apk add --no-cache git

# go.mod と go.sum をコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . /app/

# アプリケーションをビルド
RUN go build -a -installsuffix cgo -o main ./main.go
