# Exercise 12: Pipeline Builder

**Learning Goal:** Master functional transformation pipelines

---

## 📝 Problem Description

Pipelines chain transformations where each stage's output feeds into the next stage's input. This pattern is used for:
- Data processing
- Stream transformations
- ETL (Extract, Transform, Load)
- Functional composition

---

## 🎯 Function Signatures

```go
type IntTransform func(int) int

type Pipeline struct {
    transforms []IntTransform
}

func NewPipeline() *Pipeline

func (p *Pipeline) Add(transform IntTransform) *Pipeline

func (p *Pipeline) Execute(value int) int

// Helper transform functions
func Add(n int) IntTransform
func Multiply(n int) IntTransform
func Square() IntTransform
```

---

## 📖 Examples

```go
pipeline := NewPipeline().
    Add(Add(10)).
    Add(Multiply(2)).
    Add(Square())

result := pipeline.Execute(5)
// 5 -> Add(10) -> 15 -> Multiply(2) -> 30 -> Square() -> 900

// Reusable pipeline
doubleAndAdd10 := NewPipeline().
    Add(Multiply(2)).
    Add(Add(10))

doubleAndAdd10.Execute(5)   // 20
doubleAndAdd10.Execute(10)  // 30
```

---

## 📋 Instructions

1. Define `IntTransform` function type
2. Implement `Pipeline` struct
3. Implement `NewPipeline()` constructor
4. Implement `Add()` for building pipeline
5. Implement `Execute()` to run all transforms
6. Create helper transform functions
7. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Complete Solution</summary>

```go
package pipeline_builder

type IntTransform func(int) int

type Pipeline struct {
	transforms []IntTransform
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		transforms: []IntTransform{},
	}
}

func (p *Pipeline) Add(transform IntTransform) *Pipeline {
	p.transforms = append(p.transforms, transform)
	return p
}

func (p *Pipeline) Execute(value int) int {
	result := value
	for _, transform := range p.transforms {
		result = transform(result)
	}
	return result
}

func Add(n int) IntTransform {
	return func(x int) int {
		return x + n
	}
}

func Multiply(n int) IntTransform {
	return func(x int) int {
		return x * n
	}
}

func Square() IntTransform {
	return func(x int) int {
		return x * x
	}
}
```

</details>

---

## 🤔 Think About

1. How is this different from function composition?
2. What are the benefits of explicit pipelines?
3. How would you add error handling to pipelines?
4. When would you use pipelines vs direct function calls?

---

## 🎓 What This Teaches

- **Pipeline pattern**: Chaining transformations
- **Function composition**: Combining functions
- **Data flow**: Explicit transformation stages
- **Reusability**: Creating reusable transformation chains
- **Functional programming**: Pure functions and composition

---

**Tier:** 3 - Integration
**Estimated Time:** 35-45 minutes
