FROM nginx

WORKDIR /etc/nginx/conf.d

RUN apt update &&\
    apt install -y sudo&&\
    apt install -y ssh vim network-manager git iputils-ping net-tools
RUN cp default.conf ./default.conf.bak &&\
    rm default.conf&&\
    mkdir ../html

ADD ./settings .


EXPOSE 80
EXPOSE 443