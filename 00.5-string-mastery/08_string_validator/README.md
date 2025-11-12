# Exercise 08: String Validator

**Concept:** Validation logic, unicode properties, and complex boolean conditions
**Difficulty:** Medium-Hard
**Estimated Time:** 30-35 minutes

## Learning Goal

Master string validation by combining multiple conditions and using the unicode package to classify characters. Learn to write clear, maintainable validation logic that handles edge cases. This exercise teaches you how to build real-world validators like those used in forms, APIs, and security checks.

## The Problem

Validation is everywhere in software:
- Email addresses must follow a format
- Passwords must meet security requirements
- Input must contain only allowed characters
- User names must match constraints

You'll build three validators that combine multiple checks:

```go
IsValidEmail("user@example.com")     // → true
IsValidEmail("invalid.email")        // → false

IsValidPassword("Weak")              // → false (too short)
IsValidPassword("Strong123")         // → true (has upper, lower, digit, 8+ chars)

ContainsOnly("abc123", "abc123456789") // → true
ContainsOnly("abc!", "abc")           // → false (! not allowed)
```

## Your Tasks

Implement three validation functions:

### 1. IsValidEmail
Basic email validation (simplified, not RFC-compliant).

**Rules:**
- Must contain exactly one '@'
- Must have at least one character before '@' (local part)
- Must have at least one character after '@' (domain part)
- Domain must contain at least one '.'
- Domain must have at least one character after the last '.'

```go
func IsValidEmail(s string) bool
```

### 2. IsValidPassword
Password strength validator.

**Rules:**
- At least 8 characters long
- Contains at least one uppercase letter
- Contains at least one lowercase letter
- Contains at least one digit

```go
func IsValidPassword(s string) bool
```

### 3. ContainsOnly
Check if string contains only characters from an allowed set.

```go
func ContainsOnly(s string, allowed string) bool
```

## Examples

### IsValidEmail Examples
```go
IsValidEmail("user@example.com")       // → true
IsValidEmail("name.test@company.co.uk") // → true
IsValidEmail("simple@test.io")         // → true

IsValidEmail("invalid")                // → false (no @)
IsValidEmail("@example.com")           // → false (no local part)
IsValidEmail("user@")                  // → false (no domain)
IsValidEmail("user@@example.com")      // → false (multiple @)
IsValidEmail("user@nodot")             // → false (domain has no .)
IsValidEmail("user@example.")          // → false (nothing after .)
IsValidEmail("")                       // → false
```

### IsValidPassword Examples
```go
IsValidPassword("Strong123")           // → true
IsValidPassword("MyPassword1")         // → true
IsValidPassword("ValidPass9")          // → true

IsValidPassword("weak")                // → false (too short)
IsValidPassword("lowercase123")        // → false (no uppercase)
IsValidPassword("UPPERCASE123")        // → false (no lowercase)
IsValidPassword("NoDigits")            // → false (no digit)
IsValidPassword("Short1")              // → false (< 8 chars)
IsValidPassword("")                    // → false
```

### ContainsOnly Examples
```go
ContainsOnly("abc", "abcdef")          // → true
ContainsOnly("123", "0123456789")      // → true
ContainsOnly("hello", "helo")          // → true (l appears multiple times)

ContainsOnly("abc!", "abc")            // → false (! not in allowed)
ContainsOnly("123", "12")              // → false (3 not in allowed)
ContainsOnly("test", "")               // → false (no chars allowed)
ContainsOnly("", "abc")                // → true (empty string OK)
```

## Instructions

1. Open `string_validator.go`
2. Implement all three validation functions
3. Run `go test -v` to verify your solution
4. Pay special attention to edge cases

## Hints

### Hint 1: IsValidEmail Structure

```go
import "strings"

func IsValidEmail(s string) bool {
    // 1. Count '@' symbols
    atCount := strings.Count(s, "@")
    if atCount != 1 {
        return false
    }

    // 2. Split on '@' to get local and domain
    parts := strings.Split(s, "@")
    local := parts[0]
    domain := parts[1]

    // 3. Check local part exists
    if len(local) == 0 {
        return false
    }

    // 4. Check domain contains '.'
    if !strings.Contains(domain, ".") {
        return false
    }

    // 5. Check there's content after last '.'
    lastDot := strings.LastIndex(domain, ".")
    if lastDot == len(domain)-1 {  // '.' is last char
        return false
    }

    return true
}
```

### Hint 2: IsValidPassword Structure

```go
import "unicode"

func IsValidPassword(s string) bool {
    // Check length first
    if len(s) < 8 {
        return false
    }

    // Track what we've found
    hasUpper := false
    hasLower := false
    hasDigit := false

    // Check each rune
    for _, r := range s {
        if unicode.IsUpper(r) {
            hasUpper = true
        }
        if unicode.IsLower(r) {
            hasLower = true
        }
        if unicode.IsDigit(r) {
            hasDigit = true
        }
    }

    // All conditions must be true
    return hasUpper && hasLower && hasDigit
}
```

### Hint 3: ContainsOnly Approach

**Strategy 1: Set-based (using a map)**
```go
func ContainsOnly(s string, allowed string) bool {
    // Build a set of allowed runes
    allowedSet := make(map[rune]bool)
    for _, r := range allowed {
        allowedSet[r] = true
    }

    // Check each rune in s
    for _, r := range s {
        if !allowedSet[r] {
            return false  // Found a rune not in allowed set
        }
    }

    return true
}
```

**Strategy 2: Using strings.ContainsRune**
```go
import "strings"

func ContainsOnly(s string, allowed string) bool {
    for _, r := range s {
        if !strings.ContainsRune(allowed, r) {
            return false
        }
    }
    return true
}
```

Strategy 1 is more efficient for long strings (O(n) vs O(n*m)).

### Hint 4: Edge Cases to Consider

**IsValidEmail:**
- Empty string
- Multiple @ symbols
- @ at start or end
- Domain without dot
- Dot at end of domain

**IsValidPassword:**
- Empty string
- Exactly 7 characters (boundary)
- All same case
- No digits

**ContainsOnly:**
- Empty string to check (should return true)
- Empty allowed set (should return false if s is non-empty)
- Repeated characters in s

## Think About

1. **Why use `len(s)` vs `utf8.RuneCountInString(s)` for password length?**
   - Password complexity often counts bytes, not characters
   - "café" is 5 bytes but 4 runes
   - For passwords, byte count is often the requirement

2. **Why is email validation so complex in real systems?**
   - RFC 5322 allows many special characters
   - International domain names add complexity
   - Security concerns (injection attacks)
   - Our version is simplified but practical

3. **How would you make these validators more flexible?**
   - Pass requirements as parameters
   - Use a struct to configure validators
   - Return error messages, not just bool

4. **Why use boolean flags in IsValidPassword?**
   - Clear and readable
   - Easy to add more conditions
   - Short-circuit evaluation possible

## What This Teaches

- **Multi-condition validation** - combining multiple checks
- **Unicode classification** - using IsUpper, IsLower, IsDigit
- **String searching** - Count, Split, Contains, LastIndex
- **Set operations** - using maps for membership testing
- **Edge case handling** - thinking about boundary conditions
- **Boolean logic** - combining conditions with &&, ||, !
- **Code organization** - structuring validation logic clearly

## Common Mistakes to Avoid

1. **Forgetting edge cases**
   ```go
   // BUG - doesn't check for empty local part
   if strings.Contains(s, "@") {
       return true  // "@example.com" would pass!
   }

   // CORRECT
   parts := strings.Split(s, "@")
   if len(parts[0]) == 0 {
       return false
   }
   ```

2. **Using len() when you mean rune count**
   ```go
   // For password length, len() is usually correct
   // But be aware: len("café") = 5 (bytes), not 4 (runes)
   ```

3. **Not short-circuiting**
   ```go
   // INEFFICIENT - checks all runes even after finding violation
   valid := true
   for _, r := range s {
       if !allowedSet[r] {
           valid = false  // Should return immediately!
       }
   }
   return valid

   // EFFICIENT
   for _, r := range s {
       if !allowedSet[r] {
           return false  // Exit immediately
       }
   }
   return true
   ```

4. **Confusing && and ||**
   ```go
   // WRONG - only one condition needs to be true?
   return hasUpper || hasLower || hasDigit

   // RIGHT - ALL conditions must be true
   return hasUpper && hasLower && hasDigit
   ```

5. **Not handling empty strings**
   ```go
   // BUG - doesn't explicitly handle empty input
   func IsValidPassword(s string) bool {
       hasUpper := false
       // ... checks ...
       return hasUpper && hasLower && hasDigit
       // Empty string returns false, but not explicitly checked
   }

   // CLEARER
   if len(s) < 8 {
       return false  // Handles empty + too short
   }
   ```

## Challenge Extensions (Optional)

After completing the basic version, try these:

1. **IsValidEmailStrict**: Add more rules
   - No consecutive dots in local part
   - No special characters except . _ - in local part
   - Domain must be all lowercase or properly formatted

2. **PasswordStrength**: Return a strength score (0-100)
   - Length bonus
   - Character variety bonus
   - Penalize common patterns

3. **IsValidUsername**: Username validator
   - 3-20 characters
   - Alphanumeric and underscore only
   - Must start with a letter
   - Cannot end with underscore

4. **ContainsOnlyAlphanumeric**: Simplified version
   - Only letters and numbers allowed
   - Case insensitive option

## Real-World Applications

- **Web forms**: Email, password, username validation
- **APIs**: Input sanitization and validation
- **Security**: Password policy enforcement
- **Data cleaning**: Ensuring data meets format requirements
- **Configuration**: Validating config file values

## After Completing

You now understand:
- How to combine multiple validation conditions
- How to use unicode package for character classification
- How to use strings package for pattern matching
- How to handle edge cases systematically
- How to write clear, maintainable validation logic

Validation is a critical skill for building robust applications. Good validators prevent bugs, improve security, and enhance user experience!

---

**Next up:** Exercise 09 - Word Wrapper (advanced string building and layout)
