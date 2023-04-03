FROM golang:1.20 as builder

ARG PRVT_KEY
ARG MAIN_PATH

WORKDIR /root

COPY . .

RUN mkdir -p /root/.ssh
COPY $PRVT_KEY /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

RUN CGO_ENABLED=0 GOOS=linux go build -o app $MAIN_PATH

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ENV APP_CONTAINERIZATION=true

COPY --from=builder /root/app /root/app
CMD ["/root/app"]
