FROM golang:1.13.1

ADD . /root/krakenlab/plataform
WORKDIR /root/krakenlab/plataform

RUN go build -o plataform -mod vendor cmd/web/main.go

WORKDIR /root
CMD ./krakenlab/plataform/plataform
