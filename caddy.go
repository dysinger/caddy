package main

import (
    "github.com/caddyserver/caddy/caddy/caddymain"

    _ "github.com/caddyserver/dnsproviders/dnsimple"
)

func main() {
    caddymain.EnableTelemetry = false
    caddymain.Run()
}
