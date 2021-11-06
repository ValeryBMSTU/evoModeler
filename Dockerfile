FROM golang:latest 
RUN mkdir /app 
COPY . /app/ 
WORKDIR /app 
RUN go build -o ./bin/server ./cmd/main.go
EXPOSE 8080
CMD ["/app/bin/server"]