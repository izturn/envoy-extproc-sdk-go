FROM docker.m.daocloud.io/ubuntu:22.04

COPY ./grpc_health_probe /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe

COPY ./extproc /extproc
RUN chmod +x /extproc

ARG EXAMPLE=noop

EXPOSE 50051

CMD [ "/extproc", "noop"]


