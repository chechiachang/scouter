FROM python:3.6-alpine3.7


RUN apk update \
    && apk add --no-cache git \
       cmake \
       libstdc++ libgcc g++ \
       make \
#       jpeg jpeg-dev \
#       libpng libpng-dev \
#       giflib giflib-dev \
#       openblas \
#       openblas-dev \
       ca-certificates curl wget \
    && rm -rf /var/cache/apk/*

RUN pip install dlib flask face_recognition pymongo bson

# Data Directory
RUN mkdir -p /face_recognition

COPY face_recognition /face_recognition

WORKDIR /

#ARG ARG=VAL

ENTRYPOINT "python ./face_recognition/apiserver.py"
EXPOSE 5000
