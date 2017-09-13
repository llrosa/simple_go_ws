# How to build it
# docker build -t simple-go-ws -f Dockerfile .
FROM scratch

ENV PORT=443

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
CMD ["/app/main"]
