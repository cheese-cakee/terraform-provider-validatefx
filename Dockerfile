# syntax=docker/dockerfile:1
# Base image
FROM golang:1.25.2-alpine AS base
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go mod download

# Linter
FROM golangci/golangci-lint:v2.5-alpine AS lint
ENV CGO_ENABLED=0
WORKDIR /src
COPY --from=base /app .
COPY .golangci.yaml .
RUN go mod download && golangci-lint run --timeout 5m

# Binary build
FROM base AS build
WORKDIR /app
COPY --from=lint /src/lint_report.json .
RUN ls .
RUN go test ./... && go build -o terraform-provider-validatefx .


# Final image
FROM hashicorp/terraform:1.9.8

# Copy provider binary for direct execution
COPY --from=build /app/terraform-provider-validatefx /usr/local/bin/terraform-provider-validatefx

# Install provider into Terraform override path
ENV TF_PLUGIN_DIR="/root/.terraform.d/plugins"
RUN mkdir -p ${TF_PLUGIN_DIR}/registry.terraform.io/the-devops-daily/validatefx/0.0.1/linux_amd64
COPY --from=build /app/terraform-provider-validatefx ${TF_PLUGIN_DIR}/registry.terraform.io/the-devops-daily/validatefx/0.0.1/linux_amd64/terraform-provider-validatefx_v0.0.1

# Provide workspace for Terraform configuration
WORKDIR /workspace

ENTRYPOINT []
CMD ["/bin/sh"]
