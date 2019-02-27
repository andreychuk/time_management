FROM golang:1.12-alpine3.9

WORKDIR /root

COPY . .

RUN apk add --no-cache gcc git libc-dev
RUN go mod tidy
RUN go build -a -ldflags '-extldflags "-static"' .

FROM alpine:3.9

LABEL mainteiner="Vanya Andreychuk <vandreychyk@gmail.com>"
RUN apk add --no-cache \
		ca-certificates
WORKDIR /root/
COPY --from=0 /root/time_management .

CMD ["./time_management"]
