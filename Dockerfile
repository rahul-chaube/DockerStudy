FROM golang:latest 

WORKDIR /app

COPY FirstProject/go.mod FirstProject/go.sum ./

RUN go mod download

COPY FirstProject/*.go ./

RUN CGO_ENABLE=0 GOOS=linux go build -o /firstProject

EXPOSE 8080

CMD [ "/firstProject" ]



