#!/usr/bin/env bash
set -x
set -euo pipefail

cd /srv/portfolio_site

git pull

go build -o portfolio_site .

sudo -n /usr/bin/systemctl daemon-reload
sudo -n /usr/bin/systemctl restart portfolio_site.service
/usr/bin/systemctl status portfolio_site.service --no-pager
