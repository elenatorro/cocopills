################################################################################
# Cocopills Docker environment for local development                           #
# Elena Torró                                                                  #
################################################################################

FROM ubuntu:20.04 as build

################################################################################
# General                                                                      #
################################################################################
ENV DEBIAN_FRONTEND=noninteractive 

RUN apt-get update && \
    apt-get install --yes --no-install-recommends \
        sudo \
        build-essential \
        wget \
        tar \
        git \
        software-properties-common

################################################################################
# Setup sudo                                                                   #
################################################################################
RUN adduser --disabled-password --gecos '' docker
RUN adduser docker sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
USER docker


################################################################################
# Install Go                                                                   #
################################################################################
RUN sudo wget https://golang.org/dl/go1.16.3.linux-amd64.tar.gz
RUN sudo tar -xzvf go1.16.3.linux-amd64.tar.gz
RUN sudo mv go /usr/local
ENV PATH=$PATH:/usr/local/go/bin


################################################################################
# Set up Go                                                                    #
################################################################################

RUN go env -w GO111MODULE=auto
