# chal3_single_byte_xor

We're given the following hex-encoded string encrypted with single-byte XOR, and to decrypt by hand:
```
1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
```

We perform frequency analysis:
```
0x37: 5
0x78: 5
0x36: 3
0x33: 2
0x31: 2
0x1b: 2
0x39: 2
0x15: 1
0x3f: 1
0x7f: 1
0x2b: 1
0x78: 1
0x34: 1
0x3a: 1
0x3b: 1
0x3c: 1
0x3d: 1
0x3e: 1
0x2d: 1
0x28: 1
```

We can make several hypotheses here:
- `0x37`/`0x78` could be common letters in the English alphabet like `e`.
- `0x78` could be a space (`0x20`), judging from the way it is placed within the hex-encoded string.

When we test our second hypotheses (the other two give jumbled text), we XOR the entire string with `0x58` such that `0x78` maps to `0x20` (ASCII for space), and we get the following string:

```
Cooking MC's like a pound of bacon
```
