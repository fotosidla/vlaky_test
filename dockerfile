FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN export GO111MODULE=on
RUN go get github.com/fotosidla/vlaky_test/
RUN cd /app && git clone https://github.com/fotosidla/vlaky_test.git

RUN cd /app/vlaky_test/ && go build main.go

EXPOSE 8000

CMD [ "vlaky_test", "run" ]