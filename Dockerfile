FROM golang:1.16.7-alpine as builder
RUN mkdir /mybuild
ADD . /mybuild/
RUN apk update && apk add --no-cache git
WORKDIR /mybuild/cmd/telefacts
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go
FROM scratch
COPY --from=builder /mybuild /app/
WORKDIR /app/cmd/telefacts
ENTRYPOINT ["./main"]