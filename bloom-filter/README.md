# Bloom Filter

Bloom filters are a probabilistic data structure that checks for presence of an element in a set. Bloom Filters are very space-efficient and fast.

A Bloom filter enables you to check if an element is present in a set using a <b>very small memory space of a fixed size.</b>
Instead of storing all of the elements in the set, Bloom Filters hashes elements and set the corresponding bit.
While checking if element exists, it hashes the element again and check if the corresponding bit is set or not.

Bloom Filter can guarantee the absence of element in set. It is becasue for a <em>given value hash functions always generate same value.</em>
But it only give an estimate about its presence.

Bloom filters can be useful where negative answer will prevent more costly operations. For example checking if username is taken, if credit card is marked stolen.
Bloom filters are also used in <b>Databases such as Cassandra</b> for checking values in their SSTables.


### Hash Functions
Bloom Filter require hash functions which are fast, and does not require cryptographically secure hash functions such as SHA which are secure but slow.
A very common hash function used by Bloom filters is MurmurHash. It is not a cryptographic hash function It is not a cryptographic hash function.

Also to prevent collisions and reduce the false positive rate multiple hash functions are used. Each element if passed through k different hash functions which generate k diffent indices and all these indices are set to 1.
While checking if a element exists it is passed through same k hash functions and if any of the index is not set then element does not exist in the set.


### False Positive Rate and Expected Capacity

False Positive rate of bloom filter is the number of times it gives positive result but element actually does not exist in the set.
The expected false positive rate for most applications is between `1% - 10%`.  

It is important to select the right capacity for the Bloom Filter. Selecting a large number will waste storage and selecting a smaller number will increase collisions thus increasing False Positive Rate.


The actual memory used by Bloom Filter is function of choosen error rate
```
bits_per_item = -ln(error_rate) / ln(2)

memory = capacity * bits_per_item 
```

- 1% error rate require 10.08 bits per item.

Apparentely there is much more mathematics involved here.
This [link](https://en.wikipedia.org/wiki/Bloom_filter) contains all the details.

### Performance

- Insertion is O(k) operation where k is the number of hash functions used.

- Checking for element is O(k) or O(k * N) in case of stacked filter, where N is number of stacked filters.
