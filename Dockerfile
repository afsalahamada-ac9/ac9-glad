# Copyright 2024 AboveCloud9.AI Products and Services Private Limited
# All rights reserved.
# This code may not be used, copied, modified, or distributed without explicit permission.

# Build stage
FROM golang:1.23-alpine AS builder

# Install git and build dependencies
RUN apk add --no-cache git build-base

ENV USER=appuser
ENV UID=10001 
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Invoke 'make' command to build the Go binaries
RUN make linux-binaries

# Final stage
FROM scratch
LABEL org.opencontainers.image.authors="devops@abovecloud9.ai"

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy ca-certificates for HTTPS requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary
COPY --from=builder /app/bin/gcd /gcd
COPY --from=builder /app/bin/coursed /coursed
COPY --from=builder /app/bin/ldsd /ldsd
COPY --from=builder /app/bin/mediad /mediad
COPY --from=builder /app/bin/pushd /pushd
COPY --from=builder /app/bin/sfsyncd /sfsyncd

# Copy any additional config files if needed
# COPY --from=builder /app/config /config

# Use an unprivileged user.
USER appuser:appuser

# Expose port (adjust as needed)
EXPOSE 8080

# TODO: uncomment once healthcheck is supported
# Add healthcheck
# HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
#     # non-HTTP check:
#     # CMD ["/gcd", "health"]
#     # HTTP check: (use one of these)
#     CMD ["/gcd", "-http-get=http://localhost:8090/health"]

# Command to run
#ENTRYPOINT ["/gcd"] #commented as the command will be passed on container start in k8s