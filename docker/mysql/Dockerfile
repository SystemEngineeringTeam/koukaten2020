FROM mysql

RUN apt update&&\
    apt install -y ntp
ADD ./ntp /etc
ADD ./my.cnf /etc/mysql/conf.d/my.cnf

RUN ntpq -p
CMD [ "mysqld" ]