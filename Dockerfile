FROM golang:latest

WORKDIR $GOPATH/src/github.com/kbabuadze/httpheader

COPY . . 

RUN go mod tidy 

RUN go install -v . 

EXPOSE 8080

CMD ["httpheader"]
