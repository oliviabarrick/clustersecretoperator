FROM alpine:3.6

RUN adduser -D clustersecretoperator
USER clustersecretoperator

ADD tmp/_output/bin/clustersecretoperator /usr/local/bin/clustersecretoperator
