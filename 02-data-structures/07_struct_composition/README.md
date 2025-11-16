# Exercise 07: Struct Composition

## 🎯 Learning Goal
Master nested structs, composition patterns, and working with hierarchical data structures. Understand field access, struct copying, and updating nested values.

## 📝 Problem Description

Structs can contain other structs, creating hierarchical relationships. This is called "composition" - building complex types from simpler ones. You'll work with a company hierarchy: Address → Employee → Company.

**Key Concepts:**
- Nested struct definitions
- Accessing nested fields with dot notation
- Struct literals with nested initialization
- Copying structs (value semantics)
- Updating nested fields
- Slices of structs

## 🔧 Type Definitions

```go
// Address represents a physical address
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// Employee represents a company employee
type Employee struct {
	Name    string
	Age     int
	Address Address  // Nested struct
	Salary  int
}

// Company represents a company with employees
type Company struct {
	Name      string
	Employees []Employee  // Slice of structs
}
```

## 🔧 Function Signatures

```go
// NewEmployee creates an employee with given details
func NewEmployee(name string, age int, street, city, state, zip string, salary int) Employee

// NewCompany creates a company with given name
func NewCompany(name string) Company

// AddEmployee adds an employee to the company
func AddEmployee(company *Company, employee Employee)

// GetEmployeesByCity returns all employees from a specific city
func GetEmployeesByCity(company Company, city string) []Employee

// UpdateSalary updates an employee's salary by name (returns true if found)
func UpdateSalary(company *Company, employeeName string, newSalary int) bool

// AverageSalary calculates average salary of all employees
func AverageSalary(company Company) float64

// GetHighestPaid returns the employee with highest salary
func GetHighestPaid(company Company) (Employee, bool)

// CountByState returns count of employees per state
func CountByState(company Company) map[string]int
```

## 💡 Examples

```go
// Create employee with nested address
emp := NewEmployee(
    "Alice Smith", 30,
    "123 Main St", "Boston", "MA", "02101",
    75000,
)
// emp.Name == "Alice Smith"
// emp.Address.City == "Boston"

// Create company and add employees
company := NewCompany("TechCorp")
AddEmployee(&company, emp)
AddEmployee(&company, NewEmployee("Bob Jones", 35, "456 Oak Ave", "Boston", "MA", "02102", 85000))

// Filter by city
bostonEmps := GetEmployeesByCity(company, "Boston")
// Returns all employees in Boston

// Update salary
found := UpdateSalary(&company, "Alice Smith", 80000)
// found=true, Alice's salary now 80000

// Calculate average
avg := AverageSalary(company)
// avg = 82500.0

// Find highest paid
emp, found := GetHighestPaid(company)
// emp.Name == "Bob Jones", found=true

// Count by state
counts := CountByState(company)
// map[string]int{"MA": 2}
```

## 📋 Instructions

1. **NewEmployee:** Return Employee struct literal with nested Address literal
2. **NewCompany:** Return Company with name and empty employee slice
3. **AddEmployee:** Append employee to company.Employees (use pointer!)
4. **GetEmployeesByCity:** Iterate employees, filter by Address.City
5. **UpdateSalary:** Iterate employees by index, update salary if name matches
6. **AverageSalary:** Sum all salaries, divide by employee count
7. **GetHighestPaid:** Track max salary while iterating
8. **CountByState:** Build map of state → count

## 🧪 Testing

```bash
go test -v
```

Expected test count: ~30+ tests

## 🤔 Think About

1. **Why does AddEmployee take *Company?**
   - Modifying a struct requires a pointer, otherwise you modify a copy!

2. **How do you access nested fields?**
   - Use dot notation: `employee.Address.City`

3. **What happens when you copy a struct?**
   - Full copy of all fields (value semantics), nested structs are also copied

4. **Why iterate with index for UpdateSalary?**
   - `range` gives you a copy, you need `company.Employees[i]` to modify original

5. **What's the zero value of a struct?**
   - All fields set to their zero values (0 for int, "" for string, etc.)

## 💡 Hints

<details>
<summary>Hint 1: Creating nested structs</summary>

```go
func NewEmployee(name string, age int, street, city, state, zip string, salary int) Employee {
    return Employee{
        Name: name,
        Age:  age,
        Address: Address{  // Nested struct literal
            Street: street,
            City:   city,
            State:  state,
            Zip:    zip,
        },
        Salary: salary,
    }
}
```
</details>

<details>
<summary>Hint 2: Modifying struct with pointer</summary>

```go
func AddEmployee(company *Company, employee Employee) {
    company.Employees = append(company.Employees, employee)
    // Without pointer, this would modify a copy!
}
```
</details>

<details>
<summary>Hint 3: Filtering by nested field</summary>

```go
func GetEmployeesByCity(company Company, city string) []Employee {
    result := []Employee{}
    for _, emp := range company.Employees {
        if emp.Address.City == city {
            result = append(result, emp)
        }
    }
    return result
}
```
</details>

<details>
<summary>Hint 4: Updating with index</summary>

```go
func UpdateSalary(company *Company, employeeName string, newSalary int) bool {
    for i := range company.Employees {
        if company.Employees[i].Name == employeeName {
            company.Employees[i].Salary = newSalary
            return true
        }
    }
    return false
}
```
</details>

<details>
<summary>Complete Solution</summary>

```go
func NewEmployee(name string, age int, street, city, state, zip string, salary int) Employee {
    return Employee{
        Name: name,
        Age:  age,
        Address: Address{
            Street: street,
            City:   city,
            State:  state,
            Zip:    zip,
        },
        Salary: salary,
    }
}

func NewCompany(name string) Company {
    return Company{
        Name:      name,
        Employees: []Employee{},
    }
}

func AddEmployee(company *Company, employee Employee) {
    company.Employees = append(company.Employees, employee)
}

func GetEmployeesByCity(company Company, city string) []Employee {
    result := []Employee{}
    for _, emp := range company.Employees {
        if emp.Address.City == city {
            result = append(result, emp)
        }
    }
    return result
}

func UpdateSalary(company *Company, employeeName string, newSalary int) bool {
    for i := range company.Employees {
        if company.Employees[i].Name == employeeName {
            company.Employees[i].Salary = newSalary
            return true
        }
    }
    return false
}

func AverageSalary(company Company) float64 {
    if len(company.Employees) == 0 {
        return 0.0
    }
    total := 0
    for _, emp := range company.Employees {
        total += emp.Salary
    }
    return float64(total) / float64(len(company.Employees))
}

func GetHighestPaid(company Company) (Employee, bool) {
    if len(company.Employees) == 0 {
        return Employee{}, false
    }
    highest := company.Employees[0]
    for _, emp := range company.Employees {
        if emp.Salary > highest.Salary {
            highest = emp
        }
    }
    return highest, true
}

func CountByState(company Company) map[string]int {
    counts := make(map[string]int)
    for _, emp := range company.Employees {
        counts[emp.Address.State]++
    }
    return counts
}
```
</details>

## 🎓 What This Teaches

- **Struct composition** - Building complex types from simpler ones
- **Nested field access** - Using dot notation for hierarchical data
- **Pointer semantics** - When to use pointers vs values for modification
- **Struct literals** - Initializing nested structs in one expression
- **Iteration patterns** - range vs index-based when modifying
- **Aggregation** - Computing statistics over collections
- **Filtering** - Extracting subsets based on criteria
- **Map building** - Grouping and counting patterns

---

**Next Exercise:** `08_collections` - Implementing Stack and Queue with slices
