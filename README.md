```bash
nohup /arts/arthas/arthas --web.name-space=prod --web.env=prod --web.version=latest --callback.server.address=207.148.68.250:30092 > /arts/arthas/arthas.log &

nohup /arts/arthproxy/arthproxy --web.name-space=prod --web.env=prod --web.version=latest > /arts/arthproxy/arthproxy.log &
nohup /arts/grata > /arts/grata.log &
```