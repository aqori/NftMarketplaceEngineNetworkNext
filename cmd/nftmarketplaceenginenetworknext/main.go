// cmd/nftmarketplaceenginenetworknext/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplaceenginenetworknext/internal/nftmarketplaceenginenetworknext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplaceenginenetworknext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
