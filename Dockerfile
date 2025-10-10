# syntax=docker/dockerfile:1
FROM golang:1.25.2-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the provider binary
RUN CGO_ENABLED=0 go build -o terraform-provider-validatefx .

FROM hashicorp/terraform:1.9.8

# Copy provider binary for direct execution
COPY --from=build /app/terraform-provider-validatefx /usr/local/bin/terraform-provider-validatefx

# Install provider into Terraform override path
ENV TF_PLUGIN_DIR="/root/.terraform.d/plugins"
RUN mkdir -p ${TF_PLUGIN_DIR}/registry.terraform.io/The-DevOps-Daily/validatefx/0.0.1/linux_amd64
COPY --from=build /app/terraform-provider-validatefx ${TF_PLUGIN_DIR}/registry.terraform.io/The-DevOps-Daily/validatefx/0.0.1/linux_amd64/terraform-provider-validatefx_v0.0.1

# Provide workspace for Terraform configuration
WORKDIR /workspace

ENTRYPOINT []
CMD ["/bin/sh"]
