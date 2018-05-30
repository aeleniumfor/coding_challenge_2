FROM golang:latest

WORKDIR /go/src/coding_challenge_2
COPY ./ ./
RUN go get github.com/lib/pq
RUN go get github.com/gorilla/mux
RUN go install -v ./
CMD ["coding_challenge_2"]