FROM golang:1.8
WORKDIR /go/src/Birthdays
COPY  . /go/src/Birthdays

#theoretisch benötigten packages downlaoden und im Container compilen
#RUN go get -d -v ...
#RUN go get install -v 

CMD ["/go/src/Birthdays/Birthdays"]