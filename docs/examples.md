
# Examples

## CIDR Validation

```terraform
module "network_policy" {
  source = "./modules/network-policy"

  networks = [
    "10.0.0.0/24",
    "2001:db8::/48",
  ]
}
```
