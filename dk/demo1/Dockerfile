FROM golang
ADD main.go /main.go
RUN go build -o http-server main.go
CMD ["http-server"]