#!bash


if [ -f "jwtRS256.key" ]; then
    rm jwtRS256.key
fi

if [ -f "jwtRS256.key.pub" ]; then
    rm jwtRS256.key.pub
fi

echo "DON'T ADD PASSPHRASE! DON'T ADD PASSPHRASE! DON'T ADD PASSPHRASE!"

sleep 1

ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key -N ""
openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub