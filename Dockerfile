# would've use golang:alpine if it weren't for issues with running apt-get
FROM golang

RUN mkdir /app
ADD . /app/
WORKDIR /app

# get import dependencies
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/dgrijalva/jwt-go

# build bin
RUN go build -o /app_bin

# expose port for app
EXPOSE 8080

# execute app
CMD ["/app_bin"]