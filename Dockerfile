FROM golang:1.21 as builder 
ENV CGO_ENABLED 0

COPY . /ptolemaios
WORKDIR /ptolemaios/app/api

# Build the ptolemaios, passing in VERSION from the Makefile 
RUN go build -o ptolemaios

WORKDIR /ptolemaios/app/tooling

# Build Admin Migration Tool which will be used by the InitContainer
RUN go build -o migrations

FROM alpine:3.19
# Keep these ARGS in the final image
ARG BUILD_DATE
ARG APIPORT
ENV APIPORT=${APIPORT}
ARG VERSION
ENV VERSION=${VERSION}
ARG BUILD_REF
ENV BUILD_REF=${BUILD_REF}
ARG DB_DSN
ENV DB_DSN=${DB_DSN}
# Perhaps move this to k8s
ARG CLIENT_ID 
ARG CLIENT_SECRET
ENV CLIENT_ID=${CLIENT_ID}
ENV CLIENT_SECRET=${CLIENT_SECRET}
# Set initial ptolemaios admin password for db seeding.
# This is used by migrations binary.
ARG ADMIN_PASSWD
ENV ADMIN_PASSWD=${ADMIN_PASSWD}

# Ensure we have a valid user and group
RUN addgroup -g 1000 -S ptolemy && \
    adduser -u 1000 -G ptolemy -S ptolemy

# Copy application binary from builder image
COPY --from=builder --chown=ptolemy:ptolemy /ptolemaios/app/api/ptolemaios /helx/ptolemaios
# COPY --from=builder --chown=api-user:api-user /ptolemaios/cmd/ui /helx/ui
COPY --from=builder --chown=ptolemy:ptolemy /ptolemaios/app/tooling/migrations /helx/migrations

USER ptolemy
WORKDIR /helx
EXPOSE ${APIPORT}
CMD ["./ptolemaios"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="ptolemaios" \
      org.opencontainers.image.authors="Joshua Seals" \
      org.opencontainers.image.source="https://github.com/helxplatform/ptolemaios" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Renci"