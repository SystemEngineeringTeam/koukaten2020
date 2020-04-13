FROM mysql

RUN apt update&&\
    apt install -y ntp
ADD ./ntp /etc
ADD ./my.cnf /etc/mysql/conf.d/my.cnf
ADD ./hosts /etc/hosts

RUN ntpq -p
EXPOSE 33060
EXPOSE 3306
CMD [ "mysqld" ]