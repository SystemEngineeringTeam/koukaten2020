FROM golang
WORKDIR /app

EXPOSE 80

RUN apt update &&\
    apt install -y ntp
ADD ./ntp /etc
RUN ntpq -p&&\
    go get -u github.com/go-sql-driver/mysql


ADD ./app /app
CMD [ "go","run","main.go" ]