FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main .
EXPOSE 3000 3000
ENTRYPOINT ["/app/main", "-storage=postgres"]