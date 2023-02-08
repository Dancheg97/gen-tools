FROM archlinux/archlinux:base-devel

LABEL maintainer="Dancheg97 <dangdancheg@gmail.com>"

ENV IN_DOCKER=true

RUN pacman -Syu --needed --noconfirm git go
ENV PATH="${PATH}:/home/makepkg/go/bin"

RUN chmod a+rwx -R /var/cache/pacman/pkg
ARG user=makepkg
RUN useradd --system --create-home $user \
    && echo "$user ALL=(ALL:ALL) NOPASSWD:ALL" > /etc/sudoers.d/$user

USER $user
WORKDIR /home/$user
COPY . .
RUN git clone https://aur.archlinux.org/yay.git \
    && cd yay \
    && makepkg -sri --needed --noconfirm \
    && cd \
    && rm -rf .cache yay

RUN yay -Sy --noconfirm golangci-lint
RUN yay -Sy --noconfirm protoc-gen-go
RUN yay -Sy --noconfirm protoc-gen-go-grpc
RUN yay -Sy --noconfirm buf
RUN yay -Sy --noconfirm sqlc

RUN go install mvdan.cc/gofumpt@latest
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.7
RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest
RUN go install github.com/go-acme/lego/v4/cmd/lego@latest
RUN go install dancheg97.ru/templates/gen-tools@latest

USER root
