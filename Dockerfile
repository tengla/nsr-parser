FROM golang:alpine AS build

COPY cmd src/cmd
COPY stopplace src/stopplace
COPY trainroute src/trainroute
COPY go.mod src
COPY go.sum src
COPY Makefile src

RUN apk add make && \
  cd src && \
  make all

FROM golang:alpine
COPY nsr.current.xml nsr.current.xml
COPY --from=build /go/src/dist/import bin/import
COPY --from=build /go/src/dist/server bin/server
RUN import -jsonout > nsr.current.json
CMD ["server"]
