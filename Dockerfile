FROM golang:1.18

WORKDIR /go/
EXPOSE 80
COPY . /go/src

RUN ["apt", "update"]
RUN ["mkdir", "data"]
RUN ["apt", "install","-y", "curl", "autoconf", "automake", "libtool", "pkg-config"]
RUN ["git", "clone", "https://github.com/openvenues/libpostal"]
WORKDIR /go/libpostal
RUN ["./bootstrap.sh"]
RUN ["./configure", "--datadir=/go/data"]
RUN ["make"]
RUN ["make", "install"]
RUN ["ldconfig"]

WORKDIR /go/src
RUN ["go", "build", "-o", "/go/bin/main", "main.go"]

ENTRYPOINT ["/go/bin/main"]