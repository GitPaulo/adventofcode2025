# Day 2

- For Part 1 invalid IDs must
  - be even in length
  - be in the range provided inclusive
  - be a repetition of its first half (e.g. 123123)

```lua
for each range [L, R]:
    determine number of digits d
    if d is odd: continue
    k = d/2
    for A from 10^(k-1) to 10^k - 1:
        X = AA
        if L <= X <= R:
            add X to answer
```
