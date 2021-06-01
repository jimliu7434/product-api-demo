# multi-stage build 

# stage 1 :build
FROM golang:1.16 AS build-ver

# Copy app files into app folder
ADD . /src
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# stage 2 : minimize image
FROM alpine

# setting TimeZone for log time
RUN apk add --update tzdata && cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime && echo "Asia/Taipei" >  /etc/timezone && date && apk del tzdata

# Copy file from image0 to image1
COPY --from=build-ver /src/app /exec/ 
ADD ./_config /exec/_config 
#ADD ./public /exec/public 
WORKDIR /exec
RUN pwd && ls && ls /exec/_config

ENTRYPOINT ["/exec/app"]

LABEL maintainer="jimliu<jimliu7434@gmail.com>" 
LABEL service="product-api-demo"