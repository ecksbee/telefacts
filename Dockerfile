FROM golang:1.16.7-alpine as builder
RUN mkdir /mybuild
ADD . /mybuild/
WORKDIR /mybuild/cmd/telefacts
RUN apk update && apk add --no-cache git
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o /mybuild/main /mybuild/cmd/telefacts/main.go

FROM scratch
COPY --from=builder /mybuild/main /
COPY --from=builder /mybuild/data/ /data/
WORKDIR /
USER 1000
EXPOSE 8080
ENTRYPOINT ["./main"]