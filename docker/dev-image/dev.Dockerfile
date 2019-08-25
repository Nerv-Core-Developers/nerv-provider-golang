FROM augustoroman/v8-lib:6.7.290 as lib
FROM golang:1.12-stretch
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
COPY --from=lib /v8 /v8
RUN apt-get update && \
    apt-get install -y build-essential \
    libssl-dev uuid-dev libgpgme11-dev libseccomp-dev pkg-config squashfs-tools && \
    rm -rf /var/lib/apt/lists/*
RUN mkdir -p ${GOPATH}/src/github.com/sylabs && \
    cd ${GOPATH}/src/github.com/sylabs && \
    git clone https://github.com/sylabs/singularity.git && \
    cd singularity && git checkout v3.1.1
RUN cd ${GOPATH}/src/github.com/sylabs/singularity && \
    ./mconfig && \
    cd ./builddir && \
    make && \
    make install