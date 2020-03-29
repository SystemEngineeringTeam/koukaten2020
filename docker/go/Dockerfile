FROM golang
WORKDIR /app
ADD ./app /app
RUN apt update &&\
    apt install -y sudo&&\
    apt install -y ssh vim network-manager git iputils-ping net-tools

CMD [ "go","run","main.go" ]

# EXPOSE 9000
EXPOSE 80