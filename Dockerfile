# buildイメージ
FROM golang:1.13 AS builder

WORKDIR /go/src

## 依存ライブラリをダウンロードする(キャッシュを使いたいので、これを先にやる)
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

ADD . /shuffle-members
WORKDIR /shuffle-members

## main.goをコンパイルし、実行バイナリを保存
RUN CGO_ENABLED=0 GOOS=linux go build -o server ../shuffle-members/cmd/shuffle-app/main.go

# run-timeイメージ
FROM alpine:3.12.12
COPY --from=builder /shuffle-members/server /app
# TODO: テンプレートファイルも実行バイナリに含めてしまいたい
COPY --from=builder /shuffle-members/web/ /web
COPY --from=builder /shuffle-members/public/ /public
EXPOSE 50051
ENTRYPOINT ["/app"]