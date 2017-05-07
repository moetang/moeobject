# 1. PhyFile

```
--> 16 bytes
4 bytes, Header: magic number, 0x76, 0x3E, 0xB1, 0x3F
2 bytes, ID: 0-65535, LE
2 bytes, type: 0x0001-normal
4 bytes, reserved: 0x00, 0x00, 0x00, 0x00
4 bytes, Footer: magic number, 0x52, 0x1F, 0x15, 0xEA
```

# 2. File Segment

```
--> 24 bytes, fixed part
4 bytes, Header: magic number, 0xF1, 0x1E, 0xB1, 0x0C
2 bytes, type: first byte 1->simple file, 2->replica, 3->ec / second byte 1->variable length, 2 fixed length
2 bytes, ID: 0-65535, LE
4 bytes, flags:
2 bytes, file segment header size: max 512K, include fixed part, LE
2 bytes, version: 0x0001->current
4 bytes, fixed part checksum: LE
4 bytes, Footer: magic number 0x52, 0x1F, 0x15, 0xEB
--> rest of header
N bytes, rest data: length step, 1K, 4K, 16K, 32K, 256K, ALL; fields are in order
  --> 56 bytes
  16 bytes, file segment full id: none {0x00 ...}
  16 bytes, reserved id: none {0x00 ...}
  8 bytes, header checksum of full header: default: {0x00 ...}, LE
  8 bytes, sealed file checksum: default: {0x00 ...}, LE
  4 bytes, replica/ec info: 2 bytes current part no, 2 bytes total part no, LE
  4 bytes, filled: {0x00}
--> 0-65536 x File Item
```

# 3. File Item

```
--> 24 bytes
4 bytes, header: magic number, 0xD3, 0x14, 0xEF, 0x1A
4 bytes, owner id / acl id: 0->reserved, LE
4 bytes, cookie:
2 bytes, key: seq 0-65535, LE
2 bytes, alt key: 0-default, LE
4 bytes, size: max 32GB including header+footer+padding... complete File Item
4 bytes, flags:
        first byte -> high 4 bits padding length(0-7 bytes) + low 4 bits pointer type(default 0)
        second byte -> 0 bit->delete flag
--> N bytes data
--> 8 + 0-7 bytes
4 bytes footer: 0xF0, 0x07, 0xDA, 0x7A
4 bytes data checksum: LE
0-7 bytes padding: {0x88 ...}
```

# 4. index file

```
--> File Seg
2 bytes, id: LE
8 bytes, offset: absolute file offset in PhyFile
6 bytes, reserved:
--> File Seg -> File Item
2 bytes, object id: LE
4 bytes, offset: from start of File Seg, without File Seg Header
4 bytes, length: same as size field of File Item
2 bytes, alt key: LE
4 bytes, flags: same as flags field of File Item
4 bytes, padding: {0x88 ...}
```
