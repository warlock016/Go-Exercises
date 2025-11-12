# Module 00.5: String Mastery

**Custom Remediation Module - Priority: CRITICAL ⚠️**

**Prerequisites:** Completed diagnostic assessment
**Estimated Time:** 5-7 hours
**Exercises:** 10 progressive exercises

## 🎯 Why This Module?

Your diagnostic revealed a significant knowledge gap with strings, runes, and bytes:
- **Reverse String exercise:** 45 minutes (longest time spent)
- **Self-assessment:** "Extremely hard due to lack of knowledge"
- **Approach:** Had to copy from documentation

**This is your #1 priority before moving forward.** String manipulation is fundamental to almost every Go program, and understanding the difference between strings, runes, and bytes is crucial for writing correct, bug-free code.

## 📚 Learning Objectives

By the end of this module, you will:
- Understand the difference between strings, runes, and bytes
- Know when to use each type
- Iterate over strings correctly (byte vs rune iteration)
- Handle Unicode characters properly
- Use `strings.Builder` for efficient string building
- Work confidently with string manipulation
- Debug string-related issues independently

## 🔑 Key Concepts

### The Three String Types in Go

1. **`string`** - Immutable sequence of bytes (UTF-8 encoded)
   - Read-only
   - Stored as `[]byte` internally
   - Can contain any UTF-8 text

2. **`rune`** - Alias for `int32`, represents a Unicode code point
   - A single character (including emoji, accents, etc.)
   - Use when you need to work with individual characters
   - `'a'` is a rune literal

3. **`byte`** - Alias for `uint8`, represents a single byte
   - Use for ASCII or binary data
   - Part of a UTF-8 encoded string

### Critical Understanding

```go
s := "café"

// Byte iteration - iterates over UTF-8 bytes
for i := 0; i < len(s); i++ {
    fmt.Printf("%c ", s[i])  // c a f Ã ©  (WRONG for Unicode!)
}

// Rune iteration - iterates over characters
for i, r := range s {
    fmt.Printf("%c ", r)  // c a f é  (CORRECT!)
}
```

**Key Insight:** `len(s)` returns the number of **bytes**, NOT characters!

## 📋 Module Structure

### Tier 1: Understanding the Basics (Ex 01-03)
- Counting bytes vs runes
- Iterating correctly
- Basic rune operations

### Tier 2: String Building (Ex 04-06)
- Efficient string concatenation
- Using strings.Builder
- String transformation

### Tier 3: Practical Applications (Ex 07-10)
- Validation and parsing
- Complex transformations
- Real-world string problems
- Performance considerations

## 🎓 Exercise List

| # | Name | Concept | Difficulty | Est. Time |
|---|------|---------|------------|-----------|
| 01 | Byte vs Rune Count | Understanding the difference | Easy | 15 min |
| 02 | Rune Iteration | Proper string iteration | Easy | 20 min |
| 03 | Character Types | Identifying rune properties | Medium | 25 min |
| 04 | String Builder Basics | Efficient concatenation | Medium | 20 min |
| 05 | Reverse String (Redux) | Apply rune knowledge | Medium | 25 min |
| 06 | Title Case | String transformation | Medium | 30 min |
| 07 | Character Counter | Count specific runes | Medium | 25 min |
| 08 | String Validator | Unicode validation | Medium-Hard | 30 min |
| 09 | Word Wrapper | Complex string manipulation | Hard | 35 min |
| 10 | String Compressor | Practical application | Hard | 40 min |

## 🚀 How to Use This Module

### Before You Start

1. **Read the learning guides:**
   ```bash
   cat ../resources/STRING_GUIDE.md
   cat ../resources/STRING_MASTERY_CONCEPTS.md
   ```
   These comprehensive guides explain strings, runes, and bytes in depth. The `STRING_MASTERY_CONCEPTS.md` file contains your consolidated learning journey notes and detailed slice slicing explanations from this diagnostic assessment.

2. **Set your expectation:**
   - This module is DESIGNED to be challenging based on your diagnostic
   - It's okay to struggle - that's how learning happens
   - You WILL understand strings deeply after this module

### Working Through Exercises

1. **Read each exercise README carefully**
2. **Attempt independently for 15-20 minutes**
3. **Use hints if stuck**
4. **Check the example code in hints**
5. **Write your EXPLANATION.md** after completion

### Success Criteria

To complete this module:
- [ ] Complete all 10 exercises (tests passing)
- [ ] Achieve >70% first-attempt accuracy (vs 0% on diagnostic)
- [ ] Write explanations demonstrating understanding
- [ ] Can explain difference between string/rune/byte to someone else
- [ ] Confidently iterate over strings

## 📖 Recommended Reading Order

1. **Start Here:** `../resources/STRING_GUIDE.md`
2. **Your Learning Journey:** `../resources/STRING_MASTERY_CONCEPTS.md` (your consolidated notes)
3. **Exercise 01** - Understand byte vs rune count
4. **Exercise 02** - Practice correct iteration
5. **Review:** Go Blog on Strings: https://go.dev/blog/strings
6. **Continue exercises 03-10**

## ⚡ Quick Reference

### When to Use Each Type

- **`string`** - Text data, immutable, for passing around
- **`[]rune`** - When you need to modify individual characters
- **`[]byte`** - For I/O operations, binary data, mutable text
- **`strings.Builder`** - For building strings efficiently

### Common Operations

```go
// Get rune count (character count)
runeCount := utf8.RuneCountInString(s)

// Get byte count
byteCount := len(s)

// Iterate over runes
for _, r := range s {  // r is a rune
    // Process character
}

// Convert between types
runes := []rune(s)     // string -> []rune
bytes := []byte(s)     // string -> []byte
s = string(runes)      // []rune -> string
s = string(bytes)      // []byte -> string
```

## 💡 Study Tips

1. **Visualize the difference** - Draw out how "café" is stored as bytes vs runes
2. **Test your assumptions** - Use `len()` and `utf8.RuneCountInString()` frequently
3. **Print debug info** - Use `%c` (character), `%v` (value), `%T` (type)
4. **Practice with emoji** - They're multi-byte: "👍" is 4 bytes, 1 rune!
5. **Use the playground** - Test small snippets at https://go.dev/play/

## 🎯 After This Module

Once completed, you will:
- Feel confident working with strings
- Understand why your reverse string took 45 minutes
- Write correct Unicode-aware code
- Choose the right type for the task
- Never confuse bytes and runes again!

---

**Ready to master strings?** Let's turn your biggest weakness into a strength!

Start with **Exercise 01: Byte vs Rune Count** 🚀
