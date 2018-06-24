# Build
FROM golang:1.10.3 AS build-env
ADD . $GOPATH/src/app
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN cd $GOPATH/src/app && dep ensure && go build -o entry
 
# Run
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/app/entry /app/
ENTRYPOINT ./entry