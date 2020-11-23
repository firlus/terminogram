FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go build ./cmd/nononet/nononet.go
EXPOSE 42002 
CMD ["./nononet"]
