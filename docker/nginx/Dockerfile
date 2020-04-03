FROM nginx

WORKDIR /etc/nginx/conf.d


RUN apt update&&\
    apt install -y ntp
ADD ./ntp /etc
RUN cp default.conf ./default.conf.bak &&\
    rm default.conf&&\
    mkdir ../html

ADD ./settings .

EXPOSE 80
EXPOSE 443