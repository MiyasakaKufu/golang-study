FROM golang:1.22.5

# アプリケーションのソースコードをコピー

# Set the Current Working Directory inside the container
WORKDIR /app

# Go Modules を使用するための環境変数を設定
ENV GO111MODULE=on

# 必要なパッケージをインストール
# RUN apk add --no-cache git

# go.mod と go.sum をコピー ./app ごと
# COPY go.mod go.sum ./

# 依存関係をダウンロード, いらないらしい. 学びだ〜.
# RUN go mod download

# アプリケーションを実行
# CMD ["ls"]
CMD ["go", "run", "main.go"]
