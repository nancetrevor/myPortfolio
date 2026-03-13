#!/usr/bin/env bash
set -x
set -euo pipefail

cd /srv/portfolio_site

git pull

#VERSION=$(git rev-parse --short HEAD)

#go build -ldflags "-X github.com/lfizzikz/ufc_api/internal/build.Version=$VERSION" -o mma-api .

go build -o portfolio_site .

sudo -n /usr/bin/systemctl daemon-reload
sudo -n /usr/bin/systemctl restart portfolio_site.service
/usr/bin/systemctl status portfolio_site.service --no-pager
