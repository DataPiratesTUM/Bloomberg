FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /out

ENV PORT=8080
EXPOSE 8080

CMD [ "/out" ]