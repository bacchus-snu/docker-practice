FROM golang:1.19

WORKDIR /code
COPY ./ ./
RUN go build -o exe

CMD ["/code/exe"]
