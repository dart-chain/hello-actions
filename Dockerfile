FROM golang:1.25-alpine

ARG VERSION=
ENV VERSION=${VERSION}

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -ldflags="-X 'main.version=$VERSION'" -o /usr/bin/hello-actions cmd/main.go

ENTRYPOINT [ "/usr/bin/hello-actions" ]
CMD [ "" ]