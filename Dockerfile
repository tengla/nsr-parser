FROM golang:latest AS build
COPY cmd src/cmd
COPY db src/db
COPY stopplace src/stopplace
COPY go.mod src
COPY go.sum src
COPY Makefile src
RUN cd src && make all

FROM golang:alpine AS content
COPY --from=build /go/src/dist bin
COPY nsr.current.xml.gz nsr.current.xml.gz
RUN gunzip nsr.current.xml.gz
RUN ["/bin/sh", "-c", "import -jsonout >nsr.current.json"]

FROM golang:alpine
COPY --from=content /go/nsr.current.json nsr.current.json
COPY --from=build /go/src/dist bin
CMD ["find", "-key", "uicCode", "-value", "7601320", "-jsonout"]
