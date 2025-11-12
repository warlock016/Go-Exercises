# Diagnostic Assessment

Welcome to your initial skill assessment! This diagnostic will help us understand your current Go proficiency and create a personalized learning path.

## 🎯 Purpose

This assessment evaluates your understanding across 10 core Go areas:
1. Variables and basic types
2. Control structures (loops, conditionals)
3. Functions and parameters
4. Arrays and slices
5. Maps
6. Structs
7. Error handling
8. Basic testing
9. Pointers (awareness)
10. Interfaces (awareness)

**Time Estimate:** 30-60 minutes (don't rush - this is for learning, not scoring)

## 📋 Instructions

### Before You Start

1. **No AI code generation** - This is crucial for accurate assessment
2. **Documentation is allowed** - Use Go.dev, your notes, or previous code
3. **Take your time** - Understanding matters more than speed
4. **It's okay not to know** - Identifying gaps is the whole point!

### How to Complete

1. Read each exercise README carefully
2. Implement the function in the corresponding `.go` file
3. Run tests to check your solution: `go test -v`
4. If you don't know how to solve something, that's valuable data!
5. Mark exercises you couldn't complete in `results.md`

### Scoring Yourself

After completing (or attempting) all exercises:

**For each exercise, honestly assess:**
- ✅ **Completed independently** - Solved without help
- 🤔 **Completed with references** - Needed documentation/examples
- ❌ **Could not complete** - Don't know how to approach
- ⏭️ **Skipped** - Didn't attempt

## 📂 Exercise List

| # | Exercise | Concept | Difficulty |
|---|----------|---------|------------|
| 01 | Temperature Converter | Variables, types, functions | Easy |
| 02 | FizzBuzz | Control structures, conditionals | Easy |
| 03 | Sum Numbers | Slices, loops | Easy |
| 04 | Count Words | Maps, strings | Easy-Medium |
| 05 | Find Max | Slices, comparison | Easy |
| 06 | Reverse String | Strings, runes | Medium |
| 07 | Person Struct | Structs, methods | Medium |
| 08 | Divide Safely | Error handling | Medium |
| 09 | Filter Even | Slices, functions | Medium |
| 10 | Simple Validator | Interfaces (basic) | Medium-Hard |
| 11 | Update Value | Pointers | Medium |
| 12 | Basic Test | Testing awareness | Easy |

## 🚀 Getting Started

### Step 1: Run Initial Test
```bash
cd /Users/mode/Documents/Code/Go\ Exercises/00-diagnostic
go test -v
```

You should see 12 failing tests - that's expected!

### Step 2: Complete Exercises in Order
Start with `01_temperature/` and work through each exercise:

```bash
cd 01_temperature
cat README.md              # Read instructions
cat temperature.go         # See the TODO
# Implement the function
go test -v                 # Test your solution
```

### Step 3: Record Your Results
After completing all exercises, fill out `results.md` with your honest self-assessment.

### Step 4: Review Your Path
Run the diagnostic analyzer to get your personalized learning recommendations:
```bash
go run ../analyze_diagnostic.go
```

This will update your `PROGRESS.md` with:
- Your current skill level
- Identified strengths and gaps
- Recommended starting module
- Customized exercise sequence

## 💡 Tips

- **Read compiler errors carefully** - They often tell you exactly what's wrong
- **Start simple** - Get something working first, then improve
- **Test frequently** - Run tests after small changes
- **It's okay to struggle** - That means you're learning!
- **Don't spend >15min per exercise** - If stuck, mark it and move on

## ❓ Common Questions

**Q: What if I can't complete an exercise?**
A: That's perfectly fine! Mark it as incomplete in `results.md`. Identifying gaps is the point of this diagnostic.

**Q: Can I use Google/Go.dev?**
A: Yes! Looking up documentation is a real programming skill. Just don't copy code without understanding it.

**Q: What if I get syntax errors?**
A: Read the error message carefully. The Go compiler is very helpful. If you're stuck after 5 minutes, that's data about syntax familiarity.

**Q: How accurate does my code need to be?**
A: The tests check correctness. If tests pass, you're good! If they fail, try to fix them, but don't spend too long.

**Q: Should I write the most efficient solution?**
A: No - write the solution that makes sense to you. We'll work on optimization later.

## 📊 What Happens Next?

After completing the diagnostic:

1. **Your PROGRESS.md will be updated** with skill assessment results
2. **You'll get a personalized learning path** based on your strengths/gaps
3. **Some modules may be accelerated** if you show proficiency
4. **You'll start with targeted exercises** addressing your specific needs

---

**Ready? Let's begin!** Start with `01_temperature/README.md`

Remember: This is about discovering where you are, not proving where you should be. Honesty in this assessment will make your learning journey much more effective.
