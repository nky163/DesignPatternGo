FROM golang:latest
WORKDIR ~/workspace
RUN apt-get update \
    && go get golang.org/x/tools/cmd/... \
    && go get golang.org/x/tools/gopls \
    && go get github.com/go-delve/delve/cmd/dlv 
RUN echo "alias ll='ls -alF'" >> ~/.bashrc
RUN echo "alias ls='ls --color=auto'" >> ~/.bashrc

ADD . ~/workspace