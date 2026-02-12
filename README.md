# cloudflare-cli-updatedns
App para atualizar DNS magicamente


## Pre-requisito
cat > conf.yml
```
cloudflare:
  zone_id: "xxxx"
  api_token: "xxxx"
  api_url: "api.cloudflare.com/client/v4/zones"

ip:
  check_url: "http://checkip.amazonaws.com"
```

build
```
go build -o cloudflare-cli main.g
```

## uso
no terminal
```
./cloudflare-cli update-dns
ID                                  | NAME                
------------------------------------------------------------------
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    | bla.wvcardoso.dev.br
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    | bla.wvcardoso.dev.br
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    | bla.wvcardoso.dev.br
------------------------------------------------------------------
Total registros: 3

Meu IP agora: xx.xx.xx.xx

DNS atualizado: xx.xx.xx.xx -> bla.wvcardoso.dev.br
DNS atualizado: xx.xx.xx.xx -> bla.wvcardoso.dev.br
DNS atualizado: xx.xx.xx.xx -> bla.wvcardoso.dev.br
```
