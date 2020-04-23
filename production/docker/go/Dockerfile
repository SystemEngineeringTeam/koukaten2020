FROM golang
WORKDIR /app

EXPOSE 80

RUN apt update &&\
    apt install -y ntp
ADD ./ntp /etc
# waitライブラリの追加
ADD /wait /wait 

RUN ntpq -p&&\
    go get -u github.com/go-sql-driver/mysql&&\
    chmod +x /wait


ADD ./app /app
CMD [ "go","run","main.go" ]