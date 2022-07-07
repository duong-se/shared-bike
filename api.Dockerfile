FROM golang:1.18.3-alpine3.16 as shared-bike

WORKDIR /shared-bike
ADD ./ /shared-bike
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

RUN chmod +rx /usr/local/bin/wait-for ./entrypoint.sh
RUN make install

EXPOSE 8000

ENTRYPOINT [ "sh", "./entrypoint.sh" ]
