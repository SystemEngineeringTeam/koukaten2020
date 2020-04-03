FROM golang
WORKDIR /app
ADD ./app /app
ADD ../settings/ntp /etc

RUN apt update &&\
    apt install -y ntp&&\
    ntpq -p

CMD [ "go","run","main.go" ]

# EXPOSE 9000
EXPOSE 80