FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ARG IMAGE_NAME
ARG IMAGE_VENDOR
ARG IMAGE_OWNER
ARG IMAGE_VERSION
ARG IMAGE_RELEASE
ARG IMAGE_ARCH


LABEL   name="$IMAGE_NAME" \
        vendor="$IMAGE_VENDOR" \
        version="$IMAGE_VERSION" \
        release="$IMAGE_RELEASE" \
        summary="Operator to install and manage IBM Infrastructure Management" \ 
        description="IMInstall is the Schema for the iminstalls API" \
        architecture="$IMAGE_ARCH" \
        maintainer="$IMAGE_OWNER"

ENV OPERATOR=/usr/local/bin/ibm-infra-management-install-operator \
    USER_UID=1001 \
    USER_NAME=ibm-infra-management-install-operator

# install operator binary
COPY build/_output/bin/ibm-infra-management-install-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
