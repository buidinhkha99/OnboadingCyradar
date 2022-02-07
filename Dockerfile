FROM golang:1.15-alpine AS build
RUN apk add --no-cache bash

WORKDIR /src/
COPY go.mod ./
COPY go.sum ./
COPY . ./
RUN CGO_ENABLED=0 
RUN go build -o /bin/demo


FROM alpine:3.15.0
COPY --from=build /bin/demo /bin/demo
CMD ["/bin/demo"]
