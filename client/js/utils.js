function RndNum(n) {
    rnd = Math.floor(Math.random() * n);
    return rnd;
}

const PUBLIC_KEY = `
-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1PvP0k0bP9oJd2k9lUVt
u7YH7q0jmgcUPqL4jOQen6D37mFOKF4fhcqSdt65CbSgdjGP4Kln6Mvy4BHAR54i
a4fXmbxXb9pc4mhnTFGoi5eXGgcs/VfR2M5DE2Ik8Hy7XU3uhf74iVCW9MGU4LJZ
ehAa478SPHkaEnf2us0AFlbyp2IXOp11ZpIvUEsr7JC1E1wsnABVcORjSlzCE2Sz
MEBKjLBYOTDuBXRnpLCkgl1OPJ215t5FVo0tAc3LGzjNAR0xrhEAI0cKGofZaXUF
7H7ifO4Hmk/hAVv45+f3gEcqlJ5uKMoB34BH8ejTwx30SupHFP06E4c5uNGjqZ4U
ZwIDAQAB
-----END RSA PUBLIC KEY-----`;


const PRIVATE_KEY = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAxPn0RlSBN7LniniD/Oc+hb289b6OELsnBHoVe6x/Q337a5DK
SHf5tl+OBYhEAHK8F3xpYQdbCNi+iB7PUQVYFnJZAPGuuU9voDAMBpMt2kRh4DM7
R3Xjt9irfNHM/aaE8x9YbXOuDIRc/1+u5KMQGUzyoGKJtPDYYeZcFGRjUBKwXZRN
otHelnFl1fwUmSesZbXvsJYPjUuD6UCuoU3VuHgS8sxd0WMD4kCNlOGM39X48ckZ
9CmXKmc881/mG3TnAZpWXjJR5e2R3E1FiBMj0mrKQ/RZU96Jya22PPw3j68LgkvR
1alIaP94lHRIEsr06VDELoqIn1DmidyFgEkCQQIDAQABAoIBABl6b1NNiO1IcdSi
ZAgpbRflg/SRclTNsG1O0UqO9GMpf4TZVKDtC9rAH2Gtz4XzUUsEZ6kKR4csafC3
c396XnuAzOxnVn9XvAuPS27qSKsL31Edr1Q99neISh46EPbaPCYqbsixhtjNoi3S
FZAzW1i6cEO4mzFRWdvH7S8iEQqBNeQ4GS4TMPf/lGtL7klti2bIWLv7Br2WzOfm
i8GsWP8LUMxKHA2DGTfVmnphf+gR2WwSIYDk+um5A03ok0ctbQ9mjcUfVPuVdHQb
FB4dcOgZclsCB7/YtEw0wCgnaLNUspTpzxDK3c7H1+6uvvcyEBEHMtwwnC0JmVtb
nWfwYAECgYEAyNBj5kcapKzInCPd3XbYMcKnN5J4U+Nv/F4v8eZHgjhBYxJUmptT
Q5QsYDINEBeWCD0mDD7LHlA+YxZbMDlTh2mFdjeGpiZTPwd0QGj+0Mn4UqPYn8UC
oM9nIl+mNGiZbD60yDfraFG+qyooUNMUXE1xSe1tEbaL0ZUzo3gqoIECgYEA+xuU
UkYcqiC59W2dhKzXDJShOxzTJih3lbc7wGDJLrTIAtit3pVORaMgKDE2kPfpZ3ts
MbSRuekcSYsAJQ39IINJLPzulunWqSeMN14HN5h8yB+Jt97fpoL04Jqjgh5gAXvu
Ihb2yHSUZyTwQs84Q6fDdLdU1s2JbRvjt+rFgcECgYBktK0EtK/V2ZiZRRtkjs8I
1VdKdTfGyg5E/28H3rYJxfB0oKKxDigJgetnKnKGmW1yIEhOZ8cxIojG5FVCr90F
0ZNOn7X06M9iknhoPL9dMYxI1UYziXcx8hEEmfcd3T7jm+bJadGydRwrdm7VaaWS
THv92QWTlnoz1qEY35BxAQKBgQCwFg3zh9DwCFT+4ygzcpS28L8DWdpDhtsc9MoE
xJovk3wyasm6LU8fDB4vsRsHm6Fj/KvJS0tpaCt68uteEKoxk37L+m5BC6eJJP9x
kBFVjivqlhsYAkUpWenoWuhQBYbjY4mBBEN4HDA5CDEnWHKnFsy8Wxc2Lhmxln4g
aZSowQKBgCywTI1ongTB3LGqc0K+EPYW+fgM2QuDmiyztUy8YFhUjHKzyOEhWcnN
Lnc1ZEqyxPLBri/yDxMocOIVYVcCPz5DMj4k8ICPLAjBqXF+UcHCsCjqsVwIkaaQ
SjSbrbsBrgIWkR+9G7ikweoDwSmMXvHKoihV11LwFhVQbpL/yLta
-----END RSA PRIVATE KEY-----`;

function EncryptNonce(rnd) {
    let encrypt = new JSEncrypt();
    encrypt.setPublicKey(PUBLIC_KEY);
    let enc = encrypt.encrypt(rnd);
    console.log("nonce:", enc);

    // let decrypt = new JSEncrypt;
    // decrypt.setPrivateKey(PRIVATE_KEY);
    // let msg = decrypt.decrypt(enc);
    //
    // console.log("原文:", rnd);
    // console.log("加密:", enc);
    // console.log("解密:", msg);
    return enc
}

function DecryptNonce(nonce) {
    let decrypt = new JSEncrypt;
    decrypt.setPrivateKey(PRIVATE_KEY);
    let oNonce = decrypt.decrypt(nonce);
    return oNonce
}