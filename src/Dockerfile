FROM alpine:latest
MAINTAINER "Haris Shafiq(harisshafiq08@gmail.com)"
#by default alpine is not compatible with go application due to libc difference
#so creating a hardlink
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN mkdir /root/app
COPY ./main /root/app/
#as application listens on 8083 so expose it to container
EXPOSE 8083
WORKDIR /root/app
ENTRYPOINT [./main]