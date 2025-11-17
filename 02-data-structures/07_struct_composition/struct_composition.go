package struct_composition

// TODO(human): Define Address struct with Street, City, State, Zip fields
type Address struct {
	Street, City, State, Zip string // should Zip be numeric or string?
}

// TODO(human): Define Employee struct with Name, Age, Address, Salary fields
type Employee struct {
	Name    string
	Age     int
	Address Address
	Salary  int
}

// TODO(human): Define Company struct with Name and Employees fields
type Company struct {
	Name      string
	Employees []Employee // slice of strucs to accomodate for multiple employees
}

// NewEmployee creates an employee with given details
func NewEmployee(name string, age int, street, city, state, zip string, salary int) Employee {
	// TODO(human): Return Employee with nested Address
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

// NewCompany creates a company with given name
func NewCompany(name string) Company {
	// TODO(human): Return Company with given name and empty employee slice
	return Company{Name: name, Employees: []Employee{}}
}

// AddEmployee adds an employee to the company
func AddEmployee(company *Company, employee Employee) {
	// TODO(human): Add employee to company's employees
	company.Employees = append(company.Employees, employee)
}

// GetEmployeesByCity returns all employees from a specific city
func GetEmployeesByCity(company Company, city string) []Employee {
	// TODO(human): Return slice of employees from the given city
	result := []Employee{}

	for _, v := range company.Employees {
		if v.Address.City == city {
			result = append(result, v)
		}
	}

	return result
}

// UpdateSalary updates an employee's salary by name (returns true if found)
func UpdateSalary(company *Company, employeeName string, newSalary int) bool {
	// TODO(human): Find employee by name and update their salary, return true if found

	for i := range company.Employees { // when we use _, v := we are creating a copy of the underlying value, therefore getting a warning about an unused variable and failing tests!
		if company.Employees[i].Name == employeeName {
			company.Employees[i].Salary = newSalary
			return true
		}
	}
	return false
}

// AverageSalary calculates average salary of all employees
func AverageSalary(company Company) float64 {
	// TODO(human): Calculate and return average salary of all employees

	var agg int
	var count int

	for _, v := range company.Employees {
		agg += v.Salary
		count++
	}

	if count == 0 {
		return 0
	}
	return float64(agg) / float64(count)
}

// GetHighestPaid returns the employee with highest salary
func GetHighestPaid(company Company) (Employee, bool) {
	// TODO(human): Return employee with highest salary, or false if no employees

	highestSalary := 0
	bestPaid := Employee{}

	if len(company.Employees) == 0 {
		return Employee{}, false
	}

	for i := range company.Employees { // when we use _, v := we are creating a copy of the underlying value, therefore getting a warning about an unused variable and failing tests!
		if company.Employees[i].Salary > highestSalary {
			bestPaid = company.Employees[i]
			highestSalary = company.Employees[i].Salary
		}
	}

	return bestPaid, true
}

// CountByState returns count of employees per state
func CountByState(company Company) map[string]int {
	// TODO(human): Return map of employee counts by state

	result := map[string]int{}

	for _, v := range company.Employees {
		result[v.Address.State]++
	}

	return result
}
