FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/fotosidla/vlaky_test/
RUN cd /build && git clone https://github.com/fotosidla/vlaky_test.git

RUN cd /build/vlaky_test/ && go build

EXPOSE 8080

ENTRYPOINT [ "/build/vlaky_test/main" ]