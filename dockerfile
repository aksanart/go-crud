FROM golang:1.20-alpine

WORKDIR /app
ENV config="k8s"
COPY . .

RUN go mod vendor

RUN CGO_ENABLE=0 go build -o aksan-crud

ENTRYPOINT [ "./aksan-crud", "serve" ]