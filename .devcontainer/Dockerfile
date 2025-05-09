FROM ubuntu:24.04
LABEL maintainer="arilsondeveloper@gmail.com"
LABEL version="1.0"

# Evita prompts interativos durante a instalação
ENV DEBIAN_FRONTEND=noninteractive
ENV GO_VERSION=1.24.2
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID


# RUN groupmod --gid $USER_GID $USERNAME \
#     && usermod --uid $USER_UID --gid $USER_GID $USERNAME \
#     && chown -R $USER_UID:$USER_GID /home/$USERNAME

RUN apt-get update \
    && apt-get install -y git wget tar git ca-certificates coreutils build-essential gnupg libpcre3 libpcre3-dev libssl-dev zlib1g-dev redis-server \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*


RUN wget https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz \
    && tar -C /usr/local -zxf go$GO_VERSION.linux-amd64.tar.gz \
    && rm go$GO_VERSION.linux-amd64.tar.gz

# Install Node.js (LTS version) from NodeSource
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get update \
    && apt-get install -y nodejs \
    && node --version \
    && npm --version

# Install pnpm globally
RUN npm install -g pnpm \
    && pnpm --version


RUN mkdir -p /home/$USERNAME/app

USER root

WORKDIR /home/$USERNAME/app

# COPY ./.devcontainer/redis/start.sh /usr/local/bin/start.sh
# RUN chmod +x /usr/local/bin/start.sh

# Instead of using systemctl, run Redis directly
# Create directory for Redis if needed
# RUN mkdir -p /var/run/redis && chown redis:redis /var/run/redis

# Start Redis server in the background and ping it
# RUN redis-server --daemonize no && redis-cli ping

# Start Redis in foreground (Docker best practice)
# Set startup script as entrypoint
# ENTRYPOINT ["/usr/local/bin/start.sh"]
# CMD ["/bin/bash"]# Install pnpm globally
CMD ["node"]