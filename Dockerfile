# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /fb-email-service

FROM scratch

WORKDIR /

COPY --from=build /fb-email-service /fb-email-service

EXPOSE 50053

CMD [ "/fb-email-service" ]