FROM golang:1.8
WORKDIR /go/src/Birthdays
EXPOSE 9090

ENTRYPOINT /go/bin/Birthdays
#CMD ["/go/src/Birthdays/Birthdays"]


RUN go get -d -v github.com/google/go-github/github
RUN go get -d -v github.com/justinas/alice
RUN go get -d -v golang.org/x/oauth2
RUN go get -d -v golang.org/x/oauth2/github
RUN go get -d -v golang.org/x/oauth2/google
RUN go get -d -v google.golang.org/api/calendar/v3
RUN go get -d -v gopkg.in/mgo.v2
RUN go get -d -v github.com/sirupsen/logrus

COPY  . /go/src/Birthdays
RUN go install -v Birthdays




# docker run --name birthdays --rm -p 9090:9090 nniikkoollaaii/birthdays:0.4