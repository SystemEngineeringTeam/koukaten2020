FROM mysql

RUN apt update&&\
    apt install -y ntp
ADD ./ntp /etc

RUN ntpq -p
