FROM golang:alpine AS build
WORKDIR /usr/src/app

COPY *.go .
COPY *.mod .
COPY *.sum .

RUN go build -o dist/parser .

FROM golang:alpine
COPY --from=build /usr/src/app/dist/parser bin/parser
COPY nsr.current.xml.gz nsr.current.xml.gz
RUN gunzip nsr.current.xml.gz
CMD ["parser"]
