FROM golang:latest

WORKDIR /go/src/WebApp
COPY ./ ./
RUN go get github.com/lib/pq
RUN go get github.com/gorilla/mux
RUN go install -v ./
CMD ["WebApp"]