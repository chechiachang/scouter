FROM chechiachang/scouter-apiserver-base


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
