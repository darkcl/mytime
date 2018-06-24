FROM golang:1.10.3
ADD . $GOPATH/src/github.com/darkcl/mytime
RUN go get github.com/dgrijalva/jwt-go && \
  go get github.com/gin-gonic/gin && \
  go get github.com/spf13/viper && \
  go get github.com/jinzhu/gorm && \
  go get golang.org/x/crypto/bcrypt && \
  go get github.com/mattn/go-sqlite3 && \
  go get github.com/pilu/fresh
RUN cd $GOPATH/src/github.com/darkcl/mytime && go build -o entry
WORKDIR $GOPATH/src/github.com/darkcl/mytime
ENTRYPOINT ./entry