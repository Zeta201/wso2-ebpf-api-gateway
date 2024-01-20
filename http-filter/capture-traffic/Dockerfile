FROM golang:1.19.0-bullseye as builder

RUN apt-get update
RUN apt-get install -y python3-setuptools
RUN apt-get install -y zip
RUN echo deb http://cloudfront.debian.net/debian sid main >> /etc/apt/sources.list
RUN apt-get install -y bpfcc-tools libbpfcc libbpfcc-dev
RUN apt-get install -y python3-setuptools
RUN apt-get install -y kmod
RUN apt-get install -y tar
RUN apt-get install -y xz-utils 
COPY ./lib /lib
# COPY ./lib /lib
WORKDIR /app
COPY main.go .
COPY sourcecode.c .
COPY go.mod .
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=1 

CMD ["go","run","main.go","sourcecode.c"]

# RUN apt update

# RUN apt install -y linux-generic 
# RUN depmod
# RUN modprobe cordic
# CMD ["go", "run" ,"main.go", "./sourcecode.c" ]

