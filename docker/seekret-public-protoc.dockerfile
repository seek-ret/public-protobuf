FROM golang:1.17

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -yq curl unzip gnupg

RUN curl -sL https://deb.nodesource.com/setup_14.x  | bash -
RUN apt -y install python3-pip nodejs

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protoc-3.14.0-linux-x86_64.zip && \
	unzip protoc-3.14.0-linux-x86_64.zip -d /usr/local

ADD ./ /build
WORKDIR /build

RUN pip3 install -r protopy/build-requirements.txt
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 && \
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
RUN npm install

CMD ["/bin/bash"]
