FROM golang:latest
ARG token
ENV token $token
ADD . /go/src/github.com/Mi7teR/mamoeb3000
WORKDIR /go/src/github.com/Mi7teR/mamoeb3000
RUN go build -v
CMD /go/src/github.com/Mi7teR/mamoeb3000/mamoeb3000 -token=${token}