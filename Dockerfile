FROM alpine:latest

WORKDIR /app

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY ./vue-project-backend /app
COPY ./config/config.yaml /app

LABEL app-name=vue-project-backend

EXPOSE 12080

CMD ["./vue-project-backend","start"]