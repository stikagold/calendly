FROM golang:1.18-stretch as builder
EXPOSE 9997:9997
RUN mkdir /app
WORKDIR /app
COPY ./ ./
COPY local.env.yml /app/local.env.yml
RUN go build -o /app/app-exec ./main.go

ENTRYPOINT [ "./app-exec" ]
CMD ["--mode local"]
