# gmcrypto

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

Other implementation: [github.com/tjfoc/gmsm/sm2](https://github.com/tjfoc/gmsm)

| Operation    | Speed         | Allocated Mem | Mem Allocs     |
| ------------ | ------------- | ------------- | -------------- |
| Sign         | 276723 ns/op  | 4613 B/op     | 90 allocs/op   |
| Verify       | 1503850 ns/op | 84577 B/op    | 1738 allocs/op |
| VerifyFailed | 1522215 ns/op | 75034 B/op    | 1537 allocs/op |

## SM3 - Cryptographic Hash Algorithm

The algorithm is defined by GB/T 32905-2016.

The `gmcrypto/sm3` package implements

- [hash.Hash](https://pkg.go.dev/hash#Hash), which can be further used in HMAC or KDF.
- [encoding.BinaryMarshaler](https://pkg.go.dev/encoding/#BinaryMarshaler) and [encoding.BinaryUnmarshaler](https://pkg.go.dev/encoding/#BinaryUnmarshaler), which implies that this SM3 implementation is **resumable**, and its state can be encoded to or decoded from a JSON object.

### Performance

This implementation: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz

| Content Size | Speed       | Throughput  | Allocated Mem | Mem Allocs  |
| ------------ | ----------- | ----------- | ------------- | ----------- |
| 8 Bytes      | 453.5 ns/op | 17.64 MB/s  | 176 B/op      | 2 allocs/op |
| 320 Bytes    | 2220 ns/op  | 144.11 MB/s | 176 B/op      | 2 allocs/op |
| 1 KiB        | 6079 ns/op  | 168.44 MB/s | 176 B/op      | 2 allocs/op |
| 8 KiB        | 45115 ns/op | 181.58 MB/s | 176 B/op      | 2 allocs/op |

Other implementation: [github.com/tjfoc/gmsm/sm3](https://github.com/tjfoc/gmsm)

| Content Size | Speed       | Throughput  | Allocated Mem | Mem Allocs  |
| ------------ | ----------- | ----------- | ------------- | ----------- |
| 8 Bytes      | 711.6 ns/op | 11.24 MB/s  | 120 B/op      | 4 allocs/op |
| 320 Bytes    | 3243 ns/op  | 98.66 MB/s  | 440 B/op      | 5 allocs/op |
| 1 KiB        | 8393 ns/op  | 122.01 MB/s | 1144 B/op     | 5 allocs/op |
| 8 KiB        | 62547 ns/op | 130.97 MB/s | 8312 B/op     | 5 allocs/op |

## SM4 - Block Cipher Algorithm

The algorithm is defined by GB/T 32907-2016.

The `gmcrypto/sm4` package implements

- [crypto/cipher.Block](https://pkg.go.dev/crypto/cipher/#Block), which can be further used in GCM, CBC, CFB, CTR, OFB, and many other block cipher modes.

### Performance

This implementation: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz

| Operation | Speed       | Throughput  | Allocated Mem | Mem Allocs  |
| --------- | ----------- | ----------- | ------------- | ----------- |
| NewCipher | 181.6 ns/op | -           | 128 B/op      | 1 allocs/op |
| Encrypt   | 131.4 ns/op | 121.77 MB/s | 0 B/op        | 0 allocs/op |
| Decrypt   | 134.4 ns/op | 119.08 MB/s | 0 B/op        | 0 allocs/op |

Other implementation: [github.com/tjfoc/gmsm/sm4](https://github.com/tjfoc/gmsm)

| Operation | Speed       | Throughput  | Allocated Mem | Mem Allocs  |
| --------- | ----------- | ----------- | ------------- | ----------- |
| NewCipher | 324.8 ns/op | -           | 240 B/op      | 4 allocs/op |
| Encrypt   | 146.4 ns/op | 109.32 MB/s | 0 B/op        | 0 allocs/op |
| Decrypt   | 148.1 ns/op | 108.04 MB/s | 0 B/op        | 0 allocs/op |
