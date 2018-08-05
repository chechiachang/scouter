FROM chechiachang/scouter-apiserver-base

# Data Directory
RUN mkdir -p /face_recognition

COPY face_recognition /face_recognition

WORKDIR /

#ARG ARG=VAL

ENTRYPOINT "python ./face_recognition/apiserver.py"
EXPOSE 5000
