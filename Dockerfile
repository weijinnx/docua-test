FROM golang:1.13.12-alpine as builder

# maintainer info
LABEL maintainer="Vladyslav Prykhodko <weijinnx@gmail.com>"

# install git.
# git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# set the current working directory inside the container
WORKDIR /app

# copy go mod and sum files 
COPY go.mod go.sum ./

# download all dependencies. they will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download

COPY . .

# build app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# start a new stage from scratch
# it's needed to make app container smaller size
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

# copy the pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# expose app on port 8080
EXPOSE 8080

# run app (executable binary)
CMD ["./main"]