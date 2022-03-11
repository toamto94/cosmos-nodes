FROM golang:1.16 as builder

RUN git clone https://github.com/terra-money/core
RUN cd core && make install

CMD ["systemctl", "start", "terrad"]