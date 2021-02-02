FROM golang:1.15-buster
WORKDIR $GOPATH/src/go-app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["go-app"]

#docker build --rm --tag=gowebhook-example . && docker run -d -p 8080:8080 gowebhook-example