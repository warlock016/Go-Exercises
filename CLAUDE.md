# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

This is a **personalized Go learning curriculum** designed as an adaptive tutoring system. It's NOT a traditional codebase - it's an educational framework where the learner progresses through exercises with diagnostic assessment, targeted remediation, and progressive difficulty tiers.

**Key Context:**
- Student completed diagnostic on 2025-11-12
- Identified gaps: Strings/runes (45-min struggle), Testing (15-min struggle)
- Custom remediation modules created (String Mastery, Testing Fundamentals)
- Learning style: Active/hands-on with "Learn by Doing" approach
- Anti-AI-reliance protocol: Weeks 1-4 = no AI code generation, documentation only

## Architecture Overview

### Learning System Structure

```
Go Exercises/
├── PROGRESS.md          # Skill matrix, completion tracking, diagnostic results
├── JOURNAL.md           # Daily reflections, learning insights
├── go.mod              # Module: github.com/warlock016/go-exercises
├── XX-module-name/     # Each module = 10-17 progressive exercises
│   ├── README.md       # Module overview, learning objectives, success criteria
│   ├── XX_exercise/    # Individual exercise directory
│   │   ├── README.md       # Problem description, hints, examples
│   │   ├── exercise.go     # Function stubs with TODO(human) markers
│   │   └── exercise_test.go # Comprehensive test suite (table-driven)
└── resources/          # Learning guides, cheat sheets, documentation
```

### Module Taxonomy

**Diagnostic Modules** (`00-*`): Baseline assessment
- `00-diagnostic/` - 12 exercises, completed 2025-11-12
- `00.5-string-mastery/` - Custom remediation, 10 exercises (runes/bytes/strings)
- `00.6-testing-fundamentals/` - Custom remediation, 8 exercises (planned)

**Standard Modules** (`01-11`): Core curriculum
- Numbered 01-11, each with 12-20 exercises
- Progressive difficulty: Tier 1 (intro) → Tier 4 (mastery)

**Projects** (`projects/`): Capstone applications
- Milestone projects every 4 weeks
- Apply multiple modules' concepts

### Exercise Structure Pattern

Every exercise follows this format:
1. **README.md** - Problem statement, learning goals, examples, progressive hints (basic → complete solution)
2. **{name}.go** - Package with function stubs containing `// TODO(human):` markers
3. **{name}_test.go** - Table-driven tests, edge cases, benchmarks
4. **EXPLANATION.md** (student creates) - Line-by-line code explanation, alternative approaches, design decisions

## Testing Commands

### Run All Tests in a Module
```bash
cd 00-diagnostic  # or any module
go test ./...
```

### Run Single Exercise Tests
```bash
cd 00-diagnostic/01_temperature
go test -v
```

### Run Specific Test
```bash
go test -v -run TestCelsiusToFahrenheit
```

### Run Tests with Coverage
```bash
go test -cover ./...
```

### Run Benchmarks
```bash
go test -bench=. -benchmem
```

## Working with This Codebase

### When Creating New Exercises

**Required files per exercise:**
1. `README.md` - Include: Learning Goal, Problem Description, Function Signatures, Examples (3-5), Instructions, Hints (progressive: basic → intermediate → solution), "Think About" questions, "What This Teaches" section
2. `{name}.go` - **MINIMAL SCAFFOLDING ONLY**: Function signatures, brief 1-line comments describing WHAT (not HOW), TODO(human) markers. NO pre-defined structs, NO commented-out solutions, NO step-by-step pseudocode
3. `{name}_test.go` - Table-driven tests (6-8+ cases), edge cases (empty, Unicode, boundaries), benchmarks for performance awareness

**IMPORTANT - Scaffolding Philosophy (Updated 2025-11-16):**
- Student prefers **productive struggle** over hand-holding
- .go files should have MINIMAL guidance (function signature + brief comment only)
- Let student define structs, figure out algorithms, make mistakes, and learn from errors
- READMEs can be detailed (student chooses whether to read them)
- Remove: commented-out solutions, step-by-step pseudocode, pre-defined types, implementation hints

**Exercise Naming Convention:**
- Directory: `XX_exercise_name/` (e.g., `01_temperature/`, `05_reverse_string/`)
- Files: `exercise_name.go`, `exercise_name_test.go`
- Package: `package exercise_name` (matches directory name, no hyphens)

**Difficulty Progression:**
- Tier 1 (Exercises 1-3): Introduction - basic syntax, single concept
- Tier 2 (Exercises 4-7): Application - common patterns, error handling
- Tier 3 (Exercises 8-12): Integration - multiple concepts, edge cases
- Tier 4 (Exercises 13-17): Mastery - optimization, algorithms, production concerns

### Understanding Student Progress

**Check `PROGRESS.md` for:**
- Diagnostic results (skill matrix with 🔴🟡🟢🔵 levels)
- Current module status
- Identified knowledge gaps
- Completion rates and time metrics

**Check `JOURNAL.md` for:**
- Recent learning sessions
- Concepts that "clicked"
- Struggles and breakthroughs
- Self-assessment ratings (1-5)

### Learning Style Guidelines (Active Mode)

When this student is working:

**DO provide:**
- Educational insights (★ Insight boxes) before/after code
- "Learn by Doing" prompts for 2-10 line code pieces when generating 20+ lines involving design decisions
- Explanations of WHY, not just HOW
- Connections to broader patterns
- Minimal TODO(human) markers (brief directive only, no implementation steps)

**DON'T:**
- Generate complete solutions upfront (violates anti-AI-reliance protocol for Weeks 1-4)
- Use TodoWrite tool for simple tasks (only for complex multi-step work)
- Create documentation files (.md) unless explicitly requested
- Skip explanation of idiomatic Go patterns vs working-but-suboptimal code
- **Add excessive scaffolding in .go files** (no pre-defined structs, no commented-out solutions, no step-by-step pseudocode)
- **Provide "how to implement" hints in TODO markers** (describe outcome, not steps)

**Request human collaboration when:**
- Design decisions with multiple valid approaches
- Business logic with >1 correct solution
- Key algorithms or interface definitions
- Must add TODO(human) marker in code BEFORE making request

### Module-Specific Notes

**String Mastery (00.5-string-mastery/):**
- Focus: runes vs bytes vs strings, UTF-8, unicode package, strings.Builder
- Context: Student struggled 45 min on reverse string in diagnostic, needed to copy from docs
- Now understands: range over string gives runes automatically, slice slicing syntax `[:len-i]`, strings.Builder for efficiency
- Exercises 01-05 completed, 06-10 ready
- Key resource: `resources/STRING_MASTERY_CONCEPTS.md` (consolidates student's learning notes + slice slicing guide)

**Diagnostic (00-diagnostic/):**
- Completed 2025-11-12, all 12 exercises, 120 minutes total
- FizzBuzz: 5 min (used redundant conditionals, works but could simplify with i%15)
- Reverse String: 45 min (major struggle, used utf8.DecodeLastRune with += instead of strings.Builder)
- Testing: 15 min (needed examples, confused by table-driven test syntax)

## Key Resources Files

- `resources/STRING_GUIDE.md` - Comprehensive strings/runes/bytes technical guide (300+ lines)
- `resources/STRING_MASTERY_CONCEPTS.md` - Student's consolidated learning journey + slice slicing guide
- `resources/GETTING_STARTED.md` - First-day setup and workflow guide

## Anti-AI-Reliance Protocol

The student is actively trying to reduce Claude Code over-reliance. Current phase (Weeks 1-4):

**Allowed:**
- Reading documentation (go.dev)
- Searching "how to X in Go" (articles)
- Compiler error interpretation
- Asking "what does this error mean?"

**Not Allowed:**
- AI writing functions ("write a function that...")
- Copying AI-generated code without understanding
- AI debugging code directly
- AI-generated solutions

**When student asks for code:** Provide TODO(human) markers, pseudocode, hints, explanations - but make them implement. If generating >20 lines with design decisions, request human contribution on 2-10 line portions using "Learn by Doing" format.

## Important Patterns to Maintain

### Test Structure
All tests use table-driven pattern:
```go
tests := []struct {
    name string
    input Type
    want Type
}{
    {"Description", input, expected},
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        got := Function(tt.input)
        if got != tt.want {
            t.Errorf("Function(%v) = %v, want %v", tt.input, got, tt.want)
        }
    })
}
```

### Exercise README Template
- Learning Goal (1-2 sentences)
- Problem Description (with context)
- Function Signatures (exact syntax)
- Examples (3-5 with expected output)
- Instructions (numbered steps)
- Hints (progressive: basic concept → intermediate → full solution)
- "Think About" (3-4 questions for deeper understanding)
- "What This Teaches" (key takeaways)

**Note:** READMEs can be detailed - student has agency to read or skip them. The key is keeping .go files minimal.

### Exercise .go File Template (MINIMAL SCAFFOLDING)
```go
package exercise_name

// TODO(human): Import required packages if needed

// TODO(human): Define any required types/structs (if applicable)

// FunctionName does X and returns Y
func FunctionName(params) returnType {
    // TODO(human): Implement
    return zeroValue
}
```

**What NOT to include:**
- ❌ Pre-defined struct definitions with fields
- ❌ Commented-out solutions
- ❌ Step-by-step pseudocode ("Create slice, then iterate, then return")
- ❌ Specific function hints ("Use strings.Fields", "Use math.Sqrt")
- ❌ Implementation details in comments

**What TO include:**
- ✅ Package declaration
- ✅ Function signatures (for test compatibility)
- ✅ Brief 1-line comment describing function purpose
- ✅ Minimal TODO(human) marker
- ✅ Zero-value return statement (so code compiles)

### Success Criteria Updates
When student completes exercises, update `PROGRESS.md`:
- Mark exercise status (✅ completed)
- Note time spent and accuracy (independent vs with references)
- Update skill matrix levels if proficiency demonstrated
- Add "Key Insight" line summarizing learning

## Go Module Information

- Module path: `github.com/warlock016/go-exercises`
- Go version: 1.25.3
- Dependencies: `golang.org/x/sync v0.18.0`
- Each exercise is a separate package (enables isolated testing)

## Special Considerations

1. **Student is in accelerated beginner track** - Skip basic exercises they've mastered, focus on targeted gaps
2. **Diagnostic showed 58% independent completion** - Good problem-solving ability, needs syntax/idiom practice
3. **String mastery is priority #1** - 45-min diagnostic struggle now being addressed systematically
4. **Testing is priority #2** - 15-min diagnostic struggle, planned remediation module
5. **Student wants advanced struct/interface exercises** - Plan for this after remediation modules complete
