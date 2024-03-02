FROM ubuntu:12.04 AS beanstankd-builder
ENV DEBIAN_FRONTEND noninteractive
ENV PATH /usr/local/rvm/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list
RUN apt-get update
RUN apt-get -y dist-upgrade
RUN apt-get install -y beanstalkd

RUN sed -i 's/\#START=yes/START=yes/g' /etc/default/beanstalkd

EXPOSE 11300
ENTRYPOINT /usr/local/bin/beanstalkd -l 0.0.0.0 -p 11300