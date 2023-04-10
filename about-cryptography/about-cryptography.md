### encode example

```
./encode.sh "this is my test string"
# => U2FsdGVkX18fjSHGEzTXU08q+GG2I4xrV+GGtfDLg+T3vxpllUc/IHnRWLgoUl9q
```
### decode example
```
./decode.sh U2FsdGVkX18fjSHGEzTXU08q+GG2I4xrV+GGtfDLg+T3vxpllUc/IHnRWLgoUl9q
# => this is my test string
```

## Cipher Flags Available

```
Cipher commands (see the `enc' command for more details)
aes-128-cbc       aes-128-ecb       aes-192-cbc       aes-192-ecb
aes-256-cbc       aes-256-ecb       base64            bf
bf-cbc            bf-cfb            bf-ecb            bf-ofb
camellia-128-cbc  camellia-128-ecb  camellia-192-cbc  camellia-192-ecb
camellia-256-cbc  camellia-256-ecb  cast              cast-cbc
cast5-cbc         cast5-cfb         cast5-ecb         cast5-ofb
chacha            des               des-cbc           des-cfb
des-ecb           des-ede           des-ede-cbc       des-ede-cfb
des-ede-ofb       des-ede3          des-ede3-cbc      des-ede3-cfb
des-ede3-ofb      des-ofb           des3              desx
rc2               rc2-40-cbc        rc2-64-cbc        rc2-cbc
rc2-cfb           rc2-ecb           rc2-ofb           rc4
rc4-40
```

## Zero Knowledge Proofs

https://medium.com/asecuritysite-when-bob-met-alice/can-i-prove-that-i-have-used-a-random-number-generator-b7e9fb40da0e

https://asecuritysite.com/zero/