FROM golang:1.20

WORKDIR /app
COPY . /app

RUN go build -o chatter ./cmd

RUN cp /app/chatter /chatter
RUN rm -rf /app

CMD /chatter
