# detect_ecb
**[Challenge URL](https://cryptopals.com/sets/1/challenges/8)**

We detect ECB by splitting the string into 16-byte blocks, and inspecting each one for the number of repeats.

In `8.txt`, we identify line 132:
```
The line is line 132 with 3 repeats.
d880619740a8a19b7840a8a31c810a3d08649af70dc06f4fd5d2d69c744cd283e2dd052f6b641dbf9d11b0348542bb5708649af70dc06f4fd5d2d69c744cd2839475c9dfdbc1d46597949d9c7e82bf5a08649af70dc06f4fd5d2d69c744cd28397a93eab8d6aecd566489154789a6b0308649af70dc06f4fd5d2d69c744cd283d403180c98c8f6db1f2a3f9c4040deb0ab51b29933f2c123c58386b06fba186a
```
