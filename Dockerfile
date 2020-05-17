FROM golang:alpine AS build

ENV GOOS=linux
ENV GOARCH=amd64

COPY ./main.go /app/main.go

ENV CGO_ENABLED=0
RUN go build -o /app/helloworld /app/main.go

FROM scratch

COPY --from=build /app/helloworld /helloworld

EXPOSE 8080

CMD ["/helloworld"]
