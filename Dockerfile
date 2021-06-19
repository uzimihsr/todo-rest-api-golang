# ビルド用イメージ
FROM golang:1.16

# mainパッケージがあるディレクトリ(.)をまるごとコピー
COPY . /todo-api
WORKDIR /todo-api

# goapp内のgo.mod, go.sumで依存関係を管理している場合に使用
RUN ls -a ./
RUN go mod download

# クロスコンパイル
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

# バイナリを載せるイメージ
FROM scratch

# ビルド済みのバイナリをコピー
COPY --from=0 /app ./

# httpsで通信を行う場合に使用
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["./app"]