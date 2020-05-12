FROM golang
WORKDIR /app

EXPOSE 80

# waitライブラリの追加
ADD /wait /wait 

RUN go get -u github.com/go-sql-driver/mysql&&\
    chmod +x /wait&&\
    ln -sf  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime


ADD ./app /app
CMD [ "go","run","main.go" ]