FROM golang
WORKDIR /app

EXPOSE 80

RUN apt update &&\
    apt install -y ntp
ADD ./ntp /etc
RUN ntpq -p&&\
    go get -u github.com/go-sql-driver/mysql

# waitライブラリの追加
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.6.0/wait /wait 
# ライブラリの権限変更
RUN chmod +x /wait


ADD ./app /app
CMD [ "go","run","main.go" ]