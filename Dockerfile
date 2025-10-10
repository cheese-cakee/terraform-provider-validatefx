# syntax=docker/dockerfile:1
FROM golang:1.25.2-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the provider binary
RUN CGO_ENABLED=0 go build -o terraform-provider-validatefx .

FROM hashicorp/terraform:1.9.8

COPY --from=build /app/terraform-provider-validatefx /bin/terraform-provider-validatefx

# Place provider in a directory structure Terraform expects for development overrides
# terraform.d/plugins/<namespace>/<type>/<name>/<version>/<target>
RUN mkdir -p /root/.terraform.d/plugins/The-DevOps-Daily/validatefx/0.0.0/linux_amd64
COPY --from=build /app/terraform-provider-validatefx /root/.terraform.d/plugins/The-DevOps-Daily/validatefx/0.0.0/linux_amd64/terraform-provider-validatefx

WORKDIR /workspace

ENTRYPOINT ["/bin/sh"]
