FROM ubuntu:24.04
LABEL maintainer="arilsondeveloper@gmail.com"
LABEL version="1.0"

ENV GO_VERSION=1.24.2
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install -y git wget tar git ca-certificates build-essential libpcre3 libpcre3-dev libssl-dev zlib1g-dev \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN wget https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz \
    && tar -C /usr/local -zxf go$GO_VERSION.linux-amd64.tar.gz \
    && rm go$GO_VERSION.linux-amd64.tar.gz

RUN go version

# RUN groupmod --gid $USER_GID $USERNAME \
#     && usermod --uid $USER_UID --gid $USER_GID $USERNAME \
#     && chown -R $USER_UID:$USER_GID /home/$USERNAME

RUN mkdir -p /home/$USERNAME/app

USER root

WORKDIR /home/$USERNAME/app

CMD [ "/bin/bash" ]