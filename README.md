# gmcrypto
Golang crypto library based on Chinese National Standard

## SM3 - Cryptographic Hash Algorithm
The algorithm is defined by GB/T 32905-2016.

The `gmcrypto/sm3` package implements
- [hash.Hash](https://golang.org/pkg/hash/#Hash), which can be further used in HMAC or KDF.
- [encoding.BinaryMarshaler](https://golang.org/pkg/encoding/#BinaryMarshaler) and [encoding.BinaryUnmarshaler](https://golang.org/pkg/encoding/#BinaryUnmarshaler), which implies that this SM3 implementation is **resumable**, and its state can be encoded to or decoded from a JSON object.

### Performance
This implementation:
| Content Size | Speed        | Throughput | Memory Usage | Memory Alloc |
| ------------ | ------------ | ---------- | ------------ | ------------ |
| 8 Bytes      | 1212 ns/op   | 6.60 MB/s  | 176 B/op     | 2 allocs/op  |
| 320 Bytes    | 6009 ns/op   | 53.26 MB/s | 176 B/op     | 2 allocs/op  |
| 1 KiB        | 16392 ns/op  | 62.47 MB/s | 176 B/op     | 2 allocs/op  |
| 8 KiB        | 122731 ns/op | 66.75 MB/s | 176 B/op     | 2 allocs/op  |

Other implementation: [github.com/tjfoc/gmsm/sm3](https://github.com/tjfoc/gmsm)
| Content Size | Speed        | Throughput | Memory Usage | Memory Alloc |
| ------------ | ------------ | ---------- | ------------ | ------------ |
| 8 Bytes      | 1560 ns/op   | 5.13 MB/s  | 120 B/op     | 4 allocs/op  |
| 320 Bytes    | 7176 ns/op   | 44.59 MB/s | 440 B/op     | 5 allocs/op  |
| 1 KiB        | 19077 ns/op  | 53.68 MB/s | 1144 B/op    | 5 allocs/op  |
| 8 KiB        | 139148 ns/op | 58.87 MB/s | 8312 B/op    | 5 allocs/op  |
