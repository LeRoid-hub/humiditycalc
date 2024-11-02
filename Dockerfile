FROM golang:1.22

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY . .

RUN go build -o /humiditycalc

EXPOSE 8080

CMD [ "/humiditycalc" ]