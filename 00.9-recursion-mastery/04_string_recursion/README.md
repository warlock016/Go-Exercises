# Exercise 04: String Recursion

**Tier:** 2 - Patterns
**Estimated Time:** 50 minutes
**Concepts:** String processing with recursion, character manipulation, palindromes

---

## Learning Goal

Apply the slice recursion pattern to strings. Understand how strings behave differently from slices (runes vs bytes) and practice string manipulation recursively.

---

## Problem Description

Implement recursive functions that process strings character by character.

### 1. Reverse String

Reverse a string recursively.
- Input: "hello"
- Output: "olleh"

### 2. Is Palindrome

Check if a string reads the same forwards and backwards (ignoring case).
- "racecar" → true
- "hello" → false
- "A man a plan a canal Panama" → true (ignoring spaces and case)

### 3. Count Vowels

Count how many vowels (a, e, i, o, u) are in a string (case-insensitive).

### 4. Remove Character

Remove all occurrences of a specific character from a string.

---

## Function Signatures

```go
func ReverseString(s string) string

func IsPalindrome(s string) bool

func CountVowels(s string) int

func RemoveChar(s string, char rune) string
```

---

## Examples

### ReverseString
```go
ReverseString("")         // ""
ReverseString("a")        // "a"
ReverseString("hello")    // "olleh"
ReverseString("Go!")      // "!oG"
```

### IsPalindrome
```go
IsPalindrome("")              // true (empty string is palindrome)
IsPalindrome("a")             // true
IsPalindrome("racecar")       // true
IsPalindrome("hello")         // false
IsPalindrome("A man a plan a canal Panama")  // true
```

### CountVowels
```go
CountVowels("")           // 0
CountVowels("hello")      // 2 (e, o)
CountVowels("AEIOU")      // 5
CountVowels("sky")        // 1 (y is not a vowel)
CountVowels("rhythm")     // 0
```

### RemoveChar
```go
RemoveChar("hello", 'l')      // "heo"
RemoveChar("hello", 'z')      // "hello"
RemoveChar("aardvark", 'a')   // "rdvrk"
```

---

## Instructions

1. **Understand string indexing in Go**:
   ```go
   s := "hello"
   first := s[0]        // byte (not rune!)
   rest := s[1:]        // string

   // Better for Unicode:
   runes := []rune(s)
   firstRune := runes[0]
   restRunes := runes[1:]
   ```

2. **Implement ReverseString**:
   - Base case: empty or single character
   - Pattern: `reverse(rest) + first`
   - Example: `reverse("hello") = reverse("ello") + "h"`

3. **Implement IsPalindrome**:
   - Convert to lowercase first
   - Remove spaces and non-letters
   - Base case: 0 or 1 characters → true
   - Check: first == last AND IsPalindrome(middle)

4. **Implement CountVowels**:
   - Base case: empty string → 0
   - If first is vowel: 1 + CountVowels(rest)
   - Otherwise: 0 + CountVowels(rest)

5. **Implement RemoveChar**:
   - Base case: empty string → ""
   - If first matches char: RemoveChar(rest)
   - Otherwise: first + RemoveChar(rest)

6. **Run tests**: `go test -v`

---

## Hints

<details>
<summary><strong>Hint 1 - Basic (String Base Cases)</strong></summary>

**ReverseString:**
- Empty string or single character → return as is

**IsPalindrome:**
- 0 or 1 characters → always true
- Need to handle spaces and case

**CountVowels:**
- Empty string → 0
- Define vowels: `"aeiouAEIOU"`

**RemoveChar:**
- Empty string → ""

</details>

<details>
<summary><strong>Hint 2 - Intermediate (Working with Strings)</strong></summary>

**String slicing pattern:**
```go
if len(s) == 0 {
    return baseCase
}

// For ASCII strings:
first := s[0]      // byte
rest := s[1:]      // string

// For Unicode safety:
runes := []rune(s)
first := runes[0]
rest := string(runes[1:])
```

**Building strings:**
```go
return string(first) + RecursiveCall(rest)
```

**Character comparison:**
```go
if first == 'a' || first == 'e' || ... {
    // is vowel
}
```

</details>

<details>
<summary><strong>Hint 3 - Advanced (Palindrome Strategy)</strong></summary>

**IsPalindrome approach:**

1. **Clean the string first** (helper function):
   ```go
   func cleanString(s string) string {
       // Remove spaces, convert to lowercase
       // Only keep letters
   }
   ```

2. **Recursive check:**
   ```go
   if len(s) <= 1 {
       return true
   }

   if first != last {
       return false
   }

   return IsPalindrome(middle)  // s[1:len(s)-1]
   ```

**Alternative:** Convert to []rune for easier manipulation
```go
runes := []rune(s)
first := runes[0]
last := runes[len(runes)-1]
middle := runes[1:len(runes)-1]
```

</details>

<details>
<summary><strong>Hint 4 - Complete Solutions</strong></summary>

**ReverseString:**
```go
func ReverseString(s string) string {
    if len(s) <= 1 {
        return s
    }
    return ReverseString(s[1:]) + string(s[0])
}
```

**IsPalindrome:**
```go
func IsPalindrome(s string) bool {
    // Clean: remove non-letters, lowercase
    cleaned := ""
    for _, ch := range strings.ToLower(s) {
        if ch >= 'a' && ch <= 'z' {
            cleaned += string(ch)
        }
    }

    return isPalindromeHelper(cleaned)
}

func isPalindromeHelper(s string) bool {
    if len(s) <= 1 {
        return true
    }

    if s[0] != s[len(s)-1] {
        return false
    }

    return isPalindromeHelper(s[1:len(s)-1])
}
```

**CountVowels:**
```go
func CountVowels(s string) int {
    if len(s) == 0 {
        return 0
    }

    count := CountVowels(s[1:])

    ch := strings.ToLower(string(s[0]))
    if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" {
        return 1 + count
    }

    return count
}
```

**RemoveChar:**
```go
func RemoveChar(s string, char rune) string {
    if len(s) == 0 {
        return ""
    }

    runes := []rune(s)
    first := runes[0]
    rest := string(runes[1:])

    if first == char {
        return RemoveChar(rest, char)
    }

    return string(first) + RemoveChar(rest, char)
}
```

</details>

---

## Think About

1. **String immutability:** In Go, strings are immutable. Each concatenation creates a new string. For `ReverseString("hello")`, how many strings are created?

2. **Performance:** String concatenation in recursion is inefficient. How would you optimize this? (Hint: `strings.Builder`)

3. **Unicode considerations:** `s[0]` gives a byte, not a rune. When would this cause problems? (Try reversing "Hello 世界")

4. **IsPalindrome approaches:** We used a helper function. Could you do it without? Which is cleaner?

5. **Tail recursion:** None of these functions are tail-recursive. Why not? Could you convert them?

---

## What This Teaches

✅ **String manipulation** - Recursive string processing
✅ **Character checking** - Vowels, palindromes, character matching
✅ **Helper functions** - Cleaning/preprocessing before recursion
✅ **String slicing** - `s[1:]`, `s[0:len(s)-1]` patterns
✅ **Runes vs bytes** - Unicode awareness
✅ **Performance awareness** - String concatenation costs

**Connection to Module 00.5:** You learned about strings, runes, and bytes. Now you're applying that knowledge recursively!

---

**Next Exercise:** 05 - Helper Functions (mastering the accumulator pattern)
