rm *.pem
# 1. 生成 CA 的私钥和自签证书
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=CN/ST=guangdong/L=guangzhou/O=pure coder/OU=coder/CN=*.coder.com/emailAddress=1173325467@qq.com"

eho "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text
# 2. 生成 web 站点私钥和 CSR (certificate signing request)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-csr.pem -subj "/C=CN/ST=jiangxi/L=ganzhou/O=some coder/OU=some/CN=*.some.com/emailAddress=some@qq.com"

# 3. 使用 CA 的私钥去签 web 站点的 CSR，得到证书
openssl x509 -req -in server-csr.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf
echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text