# gmcrypto

![CI Test](https://github.com/need-being/gmcrypto/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/need-being/gmcrypto.svg)](https://pkg.go.dev/github.com/need-being/gmcrypto)

Golang crypto library based on Chinese National Standard

## SM2 - Signature Algorithm

The algorithm is defined by GB/T 32918.1-2016, GB/T 32918.2-2016, and GB/T 32918.5-2017.

The `gmcrypto/sm2` package implements

- [crypto.Signer](https://pkg.go.dev/crypto#Signer)

### Performance

This implementation: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz

| Operation    | Speed         | Allocated Mem | Mem Allocs      |
| ------------ | ------------- | ------------- | --------------- |
| Sign         | 1592825 ns/op | 871452 B/op   | 9294 allocs/op  |
| Verify       | 3211443 ns/op | 1724630 B/op  | 18393 allocs/op |
| VerifyFailed | 3149854 ns/op | 1695985 B/op  | 18101 allocs/op |

## SM3 - Cryptographic Hash Algorithm

The algorithm is defined by GB/T 32905-2016.

The `gmcrypto/sm3` package implements

- [hash.Hash](https://pkg.go.dev/hash#Hash), which can be further used in HMAC or KDF.
- [encoding.BinaryMarshaler](https://pkg.go.dev/encoding/#BinaryMarshaler) and [encoding.BinaryUnmarshaler](https://pkg.go.dev/encoding/#BinaryUnmarshaler), which implies that this SM3 implementation is **resumable**, and its state can be encoded to or decoded from a JSON object.

### Performance

This implementation:

| Content Size | Speed        | Throughput | Allocated Mem | Mem Allocs  |
| ------------ | ------------ | ---------- | ------------- | ----------- |
| 8 Bytes      | 1212 ns/op   | 6.60 MB/s  | 176 B/op      | 2 allocs/op |
| 320 Bytes    | 6009 ns/op   | 53.26 MB/s | 176 B/op      | 2 allocs/op |
| 1 KiB        | 16392 ns/op  | 62.47 MB/s | 176 B/op      | 2 allocs/op |
| 8 KiB        | 122731 ns/op | 66.75 MB/s | 176 B/op      | 2 allocs/op |

Other implementation: [github.com/tjfoc/gmsm/sm3](https://github.com/tjfoc/gmsm)

| Content Size | Speed        | Throughput | Allocated Mem | Mem Allocs  |
| ------------ | ------------ | ---------- | ------------- | ----------- |
| 8 Bytes      | 1560 ns/op   | 5.13 MB/s  | 120 B/op      | 4 allocs/op |
| 320 Bytes    | 7176 ns/op   | 44.59 MB/s | 440 B/op      | 5 allocs/op |
| 1 KiB        | 19077 ns/op  | 53.68 MB/s | 1144 B/op     | 5 allocs/op |
| 8 KiB        | 139148 ns/op | 58.87 MB/s | 8312 B/op     | 5 allocs/op |

## SM4 - Block Cipher Algorithm

The algorithm is defined by GB/T 32907-2016.

The `gmcrypto/sm4` package implements

- [crypto/cipher.Block](https://pkg.go.dev/crypto/cipher/#Block), which can be further used in GCM, CBC, CFB, CTR, OFB, and many other block cipher modes.

### Performance

This implementation:

| Operation | Speed     | Throughput | Allocated Mem | Mem Allocs  |
| --------- | --------- | ---------- | ------------- | ----------- |
| NewCipher | 473 ns/op | -          | 128 B/op      | 1 allocs/op |
| Encrypt   | 347 ns/op | 46.06 MB/s | 0 B/op        | 0 allocs/op |
| Decrypt   | 344 ns/op | 46.48 MB/s | 0 B/op        | 0 allocs/op |

Other implementation: [github.com/tjfoc/gmsm/sm4](https://github.com/tjfoc/gmsm)

| Operation | Speed      | Throughput | Allocated Mem | Mem Allocs  |
| --------- | ---------- | ---------- | ------------- | ----------- |
| NewCipher | 1002 ns/op | -          | 240 B/op      | 4 allocs/op |
| Encrypt   | 363 ns/op  | 44.05 MB/s | 0 B/op        | 0 allocs/op |
| Decrypt   | 357 ns/op  | 44.78 MB/s | 0 B/op        | 0 allocs/op |
