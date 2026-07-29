FROM ghcr.io/oracle/oraclelinux:9-slim
RUN microdnf -y install ca-certificates bash \
  && microdnf clean all

ARG TARGETOS
ARG TARGETARCH
ARG USER_AGENT_PROVIDER_NAME

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532
ENV USER_AGENT_PROVIDER_NAME="${USER_AGENT_PROVIDER_NAME}"

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
