## generate cert

```bash
# 生成私钥
openssl genrsa -des3 -out server.key 2048
# 生成CSR（证书签名请求）
openssl req -new -key server.key -out server.csr
# 删除私钥中的密码
cp server.key server.key.org
openssl rsa -in server.key.org -out server.key
# 生成自签名证书
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
```

## reference
> https://github.com/sensiblecodeio/tiny-ssl-reverse-proxy
