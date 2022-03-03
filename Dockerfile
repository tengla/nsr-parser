FROM golang:alpine AS build

COPY cmd src/cmd
COPY stopplace src/stopplace
COPY go.mod src
COPY go.sum src
COPY Makefile src

RUN apk add make && \
  cd src && \
  make all
COPY data/nsr.current.xml.gz .
RUN pwd
RUN gunzip nsr.current.xml.gz && stat nsr.current.xml

FROM golang:alpine
COPY --from=build /go/nsr.current.xml .
COPY --from=build /go/src/dist/import bin/import
COPY --from=build /go/src/dist/server bin/server
RUN import -jsonout > nsr.current.json
ENV PORT=8080
CMD ["server"]
