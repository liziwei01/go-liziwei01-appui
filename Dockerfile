#version 1.0
FROM golang:latest
MAINTAINER liziwei.work@gmail.com
RUN ["mkdir", "-p", "/home/work/go-liziwei01-appui"]
WORKDIR /home/work/go-liziwei01-appui
COPY . /home/work/go-liziwei01-appui
CMD ["/home/work/go-liziwei01-appui/run"] 
