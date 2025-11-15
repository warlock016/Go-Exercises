# Luhn Algorithm

## Learning Goal

Implement the Luhn algorithm (also known as the modulus 10 algorithm) to validate credit card numbers and generate check digits, learning to work with digit manipulation and real-world checksum algorithms.

## Problem Description

The Luhn algorithm is a checksum formula used to validate credit card numbers, IMEI numbers, and other identification numbers. It's designed to catch common transcription errors like single-digit errors or swapping adjacent digits.

The algorithm works by:
1. Starting from the rightmost digit (check digit), moving left
2. Doubling every second digit
3. If the doubled value is greater than 9, subtract 9 from it
4. Summing all the digits
5. The number is valid if the total sum is divisible by 10

You'll implement two functions:
- One to validate if a given number passes the Luhn check
- One to generate the check digit for a partial number

## Function Signatures

```go
func IsValidLuhn(s string) bool
func GenerateCheckDigit(s string) int
```

## Examples

### IsValidLuhn

```go
IsValidLuhn("4532015112830366")  // true (valid Visa card)
IsValidLuhn("6011111111111117")  // true (valid Discover card)
IsValidLuhn("1234567812345670")  // true
IsValidLuhn("1234567812345678")  // false (invalid check digit)
IsValidLuhn("0000000000000000")  // true (all zeros is valid)
IsValidLuhn("42")                // false (too short, typically)
IsValidLuhn("4532-0151-1283-0366") // handle with spaces/hyphens removed
```

### GenerateCheckDigit

```go
GenerateCheckDigit("453201511283036")  // 6 (making "4532015112830366" valid)
GenerateCheckDigit("601111111111111")  // 7 (making "6011111111111117" valid)
GenerateCheckDigit("123456781234567")  // 0 (making "1234567812345670" valid)
GenerateCheckDigit("0")                // 0 (making "00" valid)
```

## Instructions

1. Implement `IsValidLuhn`
2. Implement `GenerateCheckDigit`
3. Run tests with `go test -v`

## Think About

1. Why does subtracting 9 from doubled digits work instead of summing the individual digits of the doubled number? (Hint: what's 10 + 6 vs (1 + 0) + 6?)

2. How does the Luhn algorithm catch single-digit transcription errors? What about transposition errors (swapping adjacent digits)?

3. What are the limitations of the Luhn algorithm? Can you think of errors it wouldn't catch?

4. Why might you want to allow hyphens and spaces in the input but still validate correctly?

## What This Teaches

- **Digit manipulation:** Converting between strings and numeric digits, processing individual digits
- **Positional algorithms:** Working with position-dependent operations (every second digit)
- **Checksum validation:** Understanding how checksum algorithms detect errors
- **String cleaning:** Handling real-world input with formatting characters
- **Modular arithmetic:** Using modulo operations for validation
- **Real-world algorithms:** Implementing an algorithm used in production systems worldwide
