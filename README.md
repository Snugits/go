# Solution

## Main Idea

The idea is based on the finite number of possible IPv4 addresses. The total number of IPv4 addresses is 2^32 = 4,294,967,296. I created a `BitSet` structure that holds a sequence of bits. To store 4,294,967,296 bits, it takes approximately 512 MB of memory, which is reasonable for a modern computer.

In my solution, I read the file line by line and convert each IP address string into an integer. Then I check whether this integer is already present in the bit sequence. If it isn't, I increment the counter by one. After that, I set the corresponding bit in the sequence to 1 so that future occurrences of this IP can be detected.

Since Go does not provide a built-in bit sequence type, I emulated it using a `BitSet` structure. Internally, `BitSet` is represented as an array of `uint64`. Each `uint64` holds 64 bits. To store an IP address as a bit, I calculate two values:

- `idx` – the index in the `uint64` array where the bit should be stored.
- `pos` – the position of the bit within the `uint64` element.

Setting a bit is done with the operation `bits[idx] |= 1 << pos`, and checking a bit is done with `bits[idx] & (1 << pos) != 0`. This allows representing 2^32 possible IPv4 addresses in a compact fixed-size memory structure (~512 MB) without any dynamic allocation.

## Limitations

This approach works well for IPv4 addresses, but it would not scale to IPv6 addresses. The memory required for a BitSet for IPv6 would be prohibitively large.

To handle such cases, probabilistic data structures can be used. Two common choices are **Bloom Filter** and **HyperLogLog**:

- **Bloom Filter**: Suitable for large static datasets where approximate results are acceptable. However, it is not ideal for dynamically growing datasets, because increasing the dataset size without increasing the filter size raises the probability of false positives.
- **HyperLogLog**: Ideal for counting unique elements in large, dynamically growing datasets. It provides approximate results with a small, fixed error that does not depend on the dataset size.

## Performance Comparison

| Approach                | Memory Usage | Time Complexity | Accuracy           |
|-------------------------|-------------|----------------|------------------|
| HashSet (naive)         | O(N)        | O(N)           | Exact             |
| BitSet (this solution)  | ~512 MB     | O(N)           | Exact (IPv4 only) |
| Bloom Filter            | Adjustable  | O(N)           | Approximate       |
| HyperLogLog             | ~12 KB+     | O(N)           | Approximate (~2%) |

## Conclusion

The `BitSet` approach is efficient for counting unique IPv4 addresses in large files, with minimal memory usage and linear time complexity. For larger address spaces or situations where exact counts are not feasible, probabilistic structures like Bloom Filter or HyperLogLog provide scalable alternatives.
