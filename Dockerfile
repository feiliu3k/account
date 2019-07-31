FROM golang:1.12.5
COPY docker/account /data/docker/account
EXPOSE 6060
WORKDIR /data/docker/account
CMD [ "bin/account", "-c", "configs/account.json" ]
