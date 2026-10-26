# Complete Go DSA Preparation Roadmap

## Goal

Learn all the important data types, operations, manipulations, built-in functions, and problem-solving patterns required before starting medium-level DSA problems in Go (Golang).

The goal is to master the operations, patterns, and logic repeatedly used in LeetCode problems.

---

# Level 1: Integer Manipulation

Integer manipulation is essential for number-based DSA problems, mathematics, arrays, and algorithms.

## 1. Basic Integer Operations

- [ ] Declare integer variables using `int`, `int32`, and `int64`.
- [ ] Perform addition, subtraction, multiplication, division, and modulo.
- [ ] Understand integer division.
- [ ] Use increment and decrement operators: `i++`, `i--`.
- [ ] Understand negative numbers.
- [ ] Swap two variables.
- [ ] Understand operator precedence.
- [ ] Understand integer type conversions.
- [ ] Understand integer overflow and underflow.

### Example

```go
package main

import "fmt"

func main() {
    a := 17
    b := 5

    fmt.Println(a + b) // 22
    fmt.Println(a - b) // 12
    fmt.Println(a * b) // 85
    fmt.Println(a / b) // 3
    fmt.Println(a % b) // 2
}
```

## 2. Digit Manipulation

- [ ] Extract the last digit using `% 10`.
- [ ] Remove the last digit using `/ 10`.
- [ ] Count the digits in a number.
- [ ] Reverse an integer.
- [ ] Calculate the sum of digits.
- [ ] Calculate the product of digits.
- [ ] Check whether a number is a palindrome.
- [ ] Check whether a number is an Armstrong number.
- [ ] Extract digits using a loop.
- [ ] Check whether a number is prime.
- [ ] Find GCD and LCM.
- [ ] Calculate powers and factorials.
- [ ] Handle negative numbers.
- [ ] Understand integer overflow when reversing numbers.

### Example: Reverse an Integer

```go
package main

import "fmt"

func main() {
    n := 1234
    reverse := 0

    for n > 0 {
        digit := n % 10
        reverse = reverse*10 + digit
        n /= 10
    }

    fmt.Println(reverse) // 4321
}
```

## 3. Integer-Related Functions

Learn the `math` package:

- [ ] `math.Abs()`
- [ ] `math.Max()`
- [ ] `math.Min()`
- [ ] `math.Sqrt()`
- [ ] `math.Pow()`

**Important:** Many functions in `math` work with `float64`. Understand conversions between integers and floating-point types.

### Practice Problems

- Reverse Integer
- Palindrome Number
- Plus One
- Add Digits
- Happy Number
- Sqrt(x)

---

# Level 2: Array Manipulation

Arrays have a fixed length in Go. Learn arrays first, then focus on slices, which are more commonly used in DSA.

## 1. Array Basics

- [ ] Declare an array.
- [ ] Initialize an array.
- [ ] Access an element using an index.
- [ ] Update an element.
- [ ] Traverse an array using a `for` loop.
- [ ] Traverse an array using `range`.
- [ ] Find the length using `len()`.
- [ ] Find the first and last elements.
- [ ] Find the maximum and minimum elements.
- [ ] Calculate the sum and average.
- [ ] Count positive and negative elements.
- [ ] Count even and odd elements.
- [ ] Search for an element.
- [ ] Copy an array.
- [ ] Compare arrays.
- [ ] Understand array bounds.

### Example

```go
package main

import "fmt"

func main() {
    arr := [5]int{10, 20, 30, 40, 50}

    fmt.Println(arr[0])  // 10
    fmt.Println(arr[4])  // 50
    fmt.Println(len(arr)) // 5

    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
}
```

## 2. Array Modification

- [ ] Update an element at a given index.
- [ ] Swap two elements.
- [ ] Reverse an array.
- [ ] Reverse a portion of an array.
- [ ] Shift elements left.
- [ ] Shift elements right.
- [ ] Rotate an array left.
- [ ] Rotate an array right.
- [ ] Move zeros to the end.
- [ ] Remove duplicates from a sorted array.
- [ ] Insert an element using a slice.
- [ ] Delete an element using a slice.
- [ ] Merge two arrays.
- [ ] Merge two sorted arrays.
- [ ] Find the second-largest element.
- [ ] Find missing numbers.
- [ ] Find duplicate numbers.

## 3. Essential Array Patterns

| Pattern | What to Learn |
|---|---|
| Two pointers | Manipulate elements using two indices |
| Sliding window | Maintain a window over consecutive elements |
| Prefix sum | Calculate cumulative sums |
| Suffix sum | Calculate cumulative values from the right |
| Frequency counting | Count occurrences of elements |
| In-place modification | Modify data without unnecessary copies |
| Kadane's algorithm | Find the maximum-sum contiguous subarray |
| Binary search | Search efficiently in sorted arrays |
| Sorting | Arrange elements in ascending or descending order |

### Practice Problems

- Two Sum
- Best Time to Buy and Sell Stock
- Move Zeroes
- Rotate Array
- Product of Array Except Self
- Maximum Subarray

---

# Level 3: Slice Manipulation

**Slices are among the most important Go concepts for DSA.**

## 1. Basic Slice Operations

- [ ] Declare a slice.
- [ ] Initialize a slice.
- [ ] Create an empty slice.
- [ ] Access and update elements.
- [ ] Find the length using `len()`.
- [ ] Find the capacity using `cap()`.
- [ ] Understand the difference between length and capacity.
- [ ] Create a slice from another slice.
- [ ] Understand slice expressions.
- [ ] Understand how slices share their underlying array.
- [ ] Understand `nil` slices versus empty slices.
- [ ] Understand slice bounds and index errors.

### Example

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30, 40, 50}

    fmt.Println(nums[0])
    fmt.Println(nums[1:4])
    fmt.Println(len(nums))
    fmt.Println(cap(nums))
}
```

## 2. Adding and Removing Elements

- [ ] Append one element.
- [ ] Append multiple elements.
- [ ] Append one slice to another.
- [ ] Insert an element at an index.
- [ ] Remove an element at an index.
- [ ] Remove the first element.
- [ ] Remove the last element.
- [ ] Remove an element by value.
- [ ] Remove duplicates.
- [ ] Copy one slice into another.
- [ ] Understand when `append()` reallocates the underlying array.

### Example

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30}

    nums = append(nums, 40)

    fmt.Println(nums) // [10 20 30 40]
}
```

**Important:** `append()` returns a slice. Assign its result back to your variable.

## 3. Standard Library Operations

Learn the `slices` package, available in Go 1.21 and later.

- [ ] `slices.Contains()`
- [ ] `slices.Index()`
- [ ] `slices.Sort()`
- [ ] `slices.Reverse()`
- [ ] `slices.Equal()`
- [ ] `slices.Clone()`
- [ ] `slices.Delete()`
- [ ] `slices.Insert()`
- [ ] `slices.Compact()`
- [ ] `slices.Max()`
- [ ] `slices.Min()`

Learn to implement operations manually before relying on library functions.

## 4. Advanced Slice Concepts

- [ ] Two-dimensional slices.
- [ ] Slicing a slice.
- [ ] Shallow copying versus independent copying.
- [ ] Passing slices to functions.
- [ ] Modifying slices inside functions.
- [ ] Understanding shared underlying arrays.
- [ ] Avoiding unexpected changes after `append()`.
- [ ] Understanding slice bounds.

### Practice Problems

- Remove Duplicates from Sorted Array
- Remove Element
- Merge Sorted Array
- Rotate Array
- Sort Colors

---

# Level 4: String Manipulation

In Go, a string is an immutable sequence of bytes. You cannot directly modify an individual character in a string.

## 1. Basic String Operations

- [ ] Declare and initialize strings.
- [ ] Find string length using `len()`.
- [ ] Access individual bytes using an index.
- [ ] Understand bytes versus characters.
- [ ] Compare strings.
- [ ] Concatenate strings using `+`.
- [ ] Iterate through strings using a `for` loop.
- [ ] Iterate using `range`.
- [ ] Extract substrings.
- [ ] Check whether a string is empty.
- [ ] Compare strings lexicographically.

### Example

```go
package main

import "fmt"

func main() {
    s := "hello"

    fmt.Println(len(s))      // 5
    fmt.Println(s[0])        // 104: byte value of 'h'
    fmt.Println(string(s[0])) // h
    fmt.Println(s[1:4])      // ell
}
```

## 2. String Modification

- [ ] Reverse a string.
- [ ] Convert uppercase letters to lowercase.
- [ ] Convert lowercase letters to uppercase.
- [ ] Count vowels and consonants.
- [ ] Count letters, digits, and spaces.
- [ ] Remove spaces.
- [ ] Remove a particular character.
- [ ] Replace characters.
- [ ] Replace substrings.
- [ ] Split a string into parts.
- [ ] Join strings.
- [ ] Trim leading and trailing whitespace.
- [ ] Check whether a string is a palindrome.
- [ ] Check whether a string is a substring of another string.
- [ ] Find the first occurrence of a character.
- [ ] Find the last occurrence of a character.
- [ ] Find duplicate characters.
- [ ] Find the first non-repeating character.

## 3. Important Functions from the `strings` Package

| Function | Purpose |
|---|---|
| `strings.Contains()` | Check whether a substring exists |
| `strings.Index()` | Find the first occurrence |
| `strings.LastIndex()` | Find the last occurrence |
| `strings.Count()` | Count non-overlapping occurrences |
| `strings.Split()` | Split a string |
| `strings.Join()` | Join strings |
| `strings.Replace()` | Replace occurrences |
| `strings.ReplaceAll()` | Replace all occurrences |
| `strings.ToLower()` | Convert to lowercase |
| `strings.ToUpper()` | Convert to uppercase |
| `strings.TrimSpace()` | Remove surrounding whitespace |
| `strings.HasPrefix()` | Check the beginning |
| `strings.HasSuffix()` | Check the ending |
| `strings.Fields()` | Split on whitespace |
| `strings.Compare()` | Compare strings |

## 4. String-to-Number Conversion

Learn the `strconv` package:

- [ ] `strconv.Atoi()`
- [ ] `strconv.Itoa()`
- [ ] `strconv.ParseInt()`
- [ ] `strconv.FormatInt()`

Understand errors returned by conversion functions.

## 5. Bytes, Runes, and Unicode

- [ ] Understand what a byte is.
- [ ] Understand what a rune is.
- [ ] Understand why `len(s)` returns the number of bytes, not necessarily the number of characters.
- [ ] Understand that `range` iterates over Unicode code points.
- [ ] Convert a string to `[]byte`.
- [ ] Convert a string to `[]rune`.
- [ ] Modify a string using a rune slice.
- [ ] Understand ASCII character comparisons.
- [ ] Understand basic character encoding.

### Example

```go
package main

import "fmt"

func main() {
    s := "hello"

    chars := []rune(s)
    chars[0] = 'H'

    s = string(chars)

    fmt.Println(s) // Hello
}
```

### Practice Problems

- Valid Palindrome
- Reverse String
- Valid Anagram
- First Unique Character in a String
- Longest Common Prefix
- Longest Substring Without Repeating Characters

---

# Level 5: Character Manipulation

Character manipulation is closely related to strings, but it is worth practicing separately.

## Operations to Learn

- [ ] Understand ASCII values.
- [ ] Convert a character to its integer value.
- [ ] Convert an integer to a character.
- [ ] Check whether a character is uppercase.
- [ ] Check whether a character is lowercase.
- [ ] Check whether a character is a digit.
- [ ] Convert uppercase to lowercase.
- [ ] Convert lowercase to uppercase.
- [ ] Count character frequencies.
- [ ] Find duplicate characters.
- [ ] Find the most frequent character.
- [ ] Find the first non-repeating character.
- [ ] Compare character frequencies between strings.

## Useful Functions from the `unicode` Package

- `unicode.IsLetter()`
- `unicode.IsDigit()`
- `unicode.IsUpper()`
- `unicode.IsLower()`
- `unicode.ToLower()`
- `unicode.ToUpper()`

### Example

```go
package main

import "fmt"

func main() {
    ch := 'a'

    if ch >= 'a' && ch <= 'z' {
        fmt.Println("Lowercase letter")
    }

    fmt.Println(int(ch)) // 97
}
```

### Practice Problems

- Valid Anagram
- Ransom Note
- Isomorphic Strings
- Find the Difference

---

# Level 6: Map Manipulation

Maps are essential for frequency counting, lookup problems, grouping, and many medium-level DSA problems.

## Operations to Learn

- [ ] Declare a map.
- [ ] Initialize a map using `make()`.
- [ ] Insert key-value pairs.
- [ ] Retrieve values using keys.
- [ ] Update values.
- [ ] Delete entries using `delete()`.
- [ ] Check whether a key exists using the comma-ok idiom.
- [ ] Iterate over a map.
- [ ] Count the frequency of elements.
- [ ] Count character frequencies.
- [ ] Find duplicate values.
- [ ] Group elements by a key.
- [ ] Use maps with strings as keys.
- [ ] Use maps with integer keys.
- [ ] Understand that map iteration order is unspecified.
- [ ] Understand map reference-like behavior.
- [ ] Understand which key types are comparable.

### Example: Frequency Counting

```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 2, 3, 3, 3}

    freq := make(map[int]int)

    for _, num := range nums {
        freq[num]++
    }

    fmt.Println(freq)
}
```

Conceptual output:

```text
map[1:1 2:2 3:3]
```

### Practice Problems

- Two Sum
- Contains Duplicate
- Valid Anagram
- Group Anagrams
- Top K Frequent Elements

---

# Level 7: Sorting and Searching

Sorting and searching are foundational for medium-level DSA.

## 1. Sorting

- [ ] Understand ascending and descending order.
- [ ] Sort an integer slice.
- [ ] Sort a string slice.
- [ ] Sort characters in a string.
- [ ] Sort using a custom comparison.
- [ ] Sort structs by a field.
- [ ] Understand stable versus unstable sorting.
- [ ] Understand sorting time complexity.

### Useful Functions

- `sort.Ints()`
- `sort.Strings()`
- `sort.Slice()`
- `sort.SliceStable()`
- `slices.Sort()`
- `slices.SortFunc()`

## 2. Searching

- [ ] Linear search.
- [ ] Binary search.
- [ ] Binary search on sorted slices.
- [ ] Find the first occurrence of an element.
- [ ] Find the last occurrence of an element.
- [ ] Find the insertion position.
- [ ] Find the lower bound.
- [ ] Find the upper bound.
- [ ] Search in a rotated sorted array.
- [ ] Understand binary search on the answer.

### Useful Functions

- `sort.Search()`
- `slices.BinarySearch()`

### Practice Problems

- Binary Search
- Search Insert Position
- Search in Rotated Sorted Array
- Find First and Last Position of Element in Sorted Array
- Kth Largest Element in an Array

---

# Level 8: Two-Dimensional Arrays and Matrices

Matrix problems require you to manipulate rows, columns, and indices carefully.

## Operations to Learn

- [ ] Declare a two-dimensional array.
- [ ] Create a two-dimensional slice.
- [ ] Traverse a matrix row by row.
- [ ] Traverse a matrix column by column.
- [ ] Find the number of rows and columns.
- [ ] Calculate row sums.
- [ ] Calculate column sums.
- [ ] Find diagonal elements.
- [ ] Calculate diagonal sums.
- [ ] Transpose a matrix.
- [ ] Reverse rows.
- [ ] Reverse columns.
- [ ] Rotate a matrix by 90 degrees.
- [ ] Search in a sorted matrix.
- [ ] Traverse in spiral order.
- [ ] Understand matrix boundary conditions.

### Example

```go
package main

import "fmt"

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Print(matrix[i][j], " ")
        }
        fmt.Println()
    }
}
```

### Practice Problems

- Set Matrix Zeroes
- Spiral Matrix
- Rotate Image
- Search a 2D Matrix
- Search a 2D Matrix II

---

# Level 9: Other Important Go Data Types

You do not need to spend as much time on these as you spend on arrays, slices, strings, and maps. However, you should understand their fundamentals.

| Data Type | What to Learn | DSA Applications |
|---|---|---|
| `bool` | True/false and logical operators | Conditions and validation |
| `float32`, `float64` | Arithmetic, comparisons, precision | Mathematical problems |
| `byte` | Byte values and ASCII | String manipulation |
| `rune` | Unicode code points | Character processing |
| `int32`, `int64` | Integer ranges and conversions | Large numbers |
| `uint` | Unsigned integer behavior | Specialized numerical problems |
| `struct` | Fields, initialization, field access | Custom records and sorting |
| `map` | Lookup, insertion, deletion | Frequency counting |
| `slice` | Dynamic sequences and modification | Most array problems |
| `array` | Fixed-size sequences | Basic array manipulation |
| `string` | Searching, slicing, conversion | String problems |
| `pointer` | Addresses, dereferencing, nil | Linked lists and trees |

## Additional Data Structures

Learn to implement:

- [ ] Linked lists.
- [ ] Stacks.
- [ ] Queues.
- [ ] Deques.
- [ ] Binary trees.
- [ ] Binary search trees.
- [ ] Heaps and priority queues.
- [ ] Graphs.
- [ ] Tries.

Also understand recursion, backtracking, and basic dynamic programming.

These are algorithmic concepts rather than data types, but they are essential for medium-level DSA.

---

# Level 10: Go Language Features Required for DSA

Even if you understand the data types, you can still struggle with DSA if you do not know how to use Go effectively.

## 1. Loops and Conditions

- [ ] `if`, `else if`, and `else`.
- [ ] `for` loops.
- [ ] Nested loops.
- [ ] `range`.
- [ ] `break` and `continue`.
- [ ] Multiple loop variables.
- [ ] Looping backward.
- [ ] Infinite loops and termination conditions.

## 2. Functions

- [ ] Define functions.
- [ ] Pass arguments.
- [ ] Return values.
- [ ] Return multiple values.
- [ ] Pass slices to functions.
- [ ] Modify slices inside functions.
- [ ] Understand recursion.
- [ ] Write helper functions.
- [ ] Understand value copying versus pointer-based modification.

## 3. Built-in Functions

| Function | Purpose |
|---|---|
| `len()` | Length of a string, array, slice, map, or channel |
| `cap()` | Capacity of arrays, slices, and channels |
| `append()` | Add elements to a slice |
| `copy()` | Copy slice elements |
| `make()` | Initialize slices, maps, and channels |
| `new()` | Allocate zero-valued storage and return a pointer |
| `delete()` | Remove a map entry |
| `clear()` | Clear map entries or zero slice elements in modern Go |
| `min()` | Find the minimum of supported ordered values in modern Go |
| `max()` | Find the maximum of supported ordered values in modern Go |

**Version note:** Built-in `min()`, `max()`, and `clear()` require Go 1.21 or later. The `slices` package was also introduced in Go 1.21.

Check your installed Go version:

```bash
go version
```

---

# Level 11: Problem-Solving Patterns Before Medium-Level DSA

Understanding individual data types is only the beginning. You must also understand how to combine them to solve problems.

| Pattern | Key Concepts | Example Problems |
|---|---|---|
| Two pointers | Slices, indices, comparisons | Two Sum II, 3Sum |
| Sliding window | Slices, strings, maps | Longest Substring Without Repeating Characters |
| Prefix sum | Integers, slices | Subarray Sum Equals K |
| Frequency counting | Maps, strings, slices | Group Anagrams |
| Binary search | Slices, loops, boundaries | Search in Rotated Sorted Array |
| Fast and slow pointers | Pointers, linked lists | Linked List Cycle |
| Monotonic stack | Slices, stacks | Daily Temperatures |
| Intervals | Slices, sorting | Merge Intervals |
| Matrix traversal | Two-dimensional slices | Spiral Matrix |
| Backtracking | Recursion, slices | Subsets, Permutations |
| Heap | Slices, custom data structures | Top K Frequent Elements |
| Tree traversal | Recursion, structs, pointers | Binary Tree Level Order Traversal |
| Graph traversal | Maps, slices, queues | Number of Islands |
| Dynamic programming | Arrays, recurrence relations | House Robber |

---

# Level 12: Time and Space Complexity

Before attempting medium-level problems, understand the following complexities.

| Complexity | Meaning | Example |
|---|---|---|
| O(1) | Constant time | Access an array element |
| O(log n) | Logarithmic time | Binary search |
| O(n) | Linear time | Traverse a slice |
| O(n log n) | Linearithmic time | Typical comparison sorting |
| O(n²) | Quadratic time | Compare every pair |
| O(2ⁿ) | Exponential time | Some brute-force subset algorithms |
| O(n!) | Factorial time | Generate all permutations |

## Complexity Topics to Learn

- [ ] Time complexity of nested loops.
- [ ] Space complexity of additional slices.
- [ ] Space complexity of maps.
- [ ] Auxiliary space versus total space.
- [ ] Best-case, average-case, and worst-case complexity.
- [ ] Identify unnecessary nested loops.
- [ ] Understand trading extra memory for faster execution.

---

# Final Checklist Before Starting Medium-Level LeetCode Problems

| No. | Topic | Priority |
|---:|---|---|
| 1 | Integer manipulation | Must learn |
| 2 | Arrays | Must learn |
| 3 | Slices and `append()` | Must learn |
| 4 | Strings | Must learn |
| 5 | Bytes and runes | Must learn |
| 6 | Character manipulation | Must learn |
| 7 | Maps and frequency counting | Must learn |
| 8 | Sorting and searching | Must learn |
| 9 | Two-dimensional arrays | Must learn |
| 10 | Loops and functions | Must learn |
| 11 | Recursion | Must learn |
| 12 | Time and space complexity | Must learn |
| 13 | Two pointers | Must learn |
| 14 | Sliding window | Must learn |
| 15 | Prefix sums | Must learn |
| 16 | Stacks and queues | Next |
| 17 | Linked lists | Next |
| 18 | Trees and heaps | Next |
| 19 | Graphs | Next |
| 20 | Backtracking and dynamic programming | Next |

---

# Recommended Learning Order

Follow this sequence instead of learning everything simultaneously.

## Phase 1: Data Manipulation

1. Integers.
2. Arrays.
3. Slices.
4. Strings.
5. Characters, bytes, and runes.

## Phase 2: Data Structures and Operations

6. Maps.
7. Sorting and searching.
8. Two-dimensional slices.
9. Structs and pointers.
10. Stacks and queues.

## Phase 3: Problem-Solving Patterns

11. Two pointers.
12. Sliding window.
13. Frequency counting.
14. Prefix sums.
15. Binary search.
16. Recursion and backtracking.

## Phase 4: Medium-Level DSA

17. Linked lists.
18. Trees and heaps.
19. Graphs.
20. Dynamic programming.

---

# How to Practice Each Topic

Do not just read these topics. Write small programs for every operation.

For each topic:

1. Understand the syntax.
2. Write a simple program.
3. Solve 5–10 small exercises.
4. Try solving the exercises without looking at the solution.
5. Test edge cases.
6. Analyze time and space complexity.
7. Solve easy LeetCode problems.
8. Move to medium-level problems after mastering the basics.

## Example: Preparation for Longest Substring Without Repeating Characters

Before solving this problem, make sure you can independently:

- [ ] Iterate through a string.
- [ ] Understand character representation.
- [ ] Manipulate two indices.
- [ ] Use a map to track character positions.
- [ ] Update a maximum value.
- [ ] Analyze time complexity.
- [ ] Analyze space complexity.

## Final Recommendation

Complete Phase 1 first, practice the operations with small coding exercises, and then solve easy LeetCode problems before progressing to medium-level DSA.

**The goal is not to memorize every Go function. The goal is to understand the data types and operations well enough to implement the required logic independently.**
