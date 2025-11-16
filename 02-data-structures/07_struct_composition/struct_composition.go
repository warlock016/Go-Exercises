package struct_composition

// TODO(human): Define Address struct with Street, City, State, Zip fields

// TODO(human): Define Employee struct with Name, Age, Address, Salary fields

// TODO(human): Define Company struct with Name and Employees fields

// NewEmployee creates an employee with given details
func NewEmployee(name string, age int, street, city, state, zip string, salary int) Employee {
	// TODO(human): Return Employee with nested Address
	return Employee{}
}

// NewCompany creates a company with given name
func NewCompany(name string) Company {
	// TODO(human): Return Company with given name and empty employee slice
	return Company{}
}

// AddEmployee adds an employee to the company
func AddEmployee(company *Company, employee Employee) {
	// TODO(human): Add employee to company's employees
}

// GetEmployeesByCity returns all employees from a specific city
func GetEmployeesByCity(company Company, city string) []Employee {
	// TODO(human): Return slice of employees from the given city
	return nil
}

// UpdateSalary updates an employee's salary by name (returns true if found)
func UpdateSalary(company *Company, employeeName string, newSalary int) bool {
	// TODO(human): Find employee by name and update their salary, return true if found
	return false
}

// AverageSalary calculates average salary of all employees
func AverageSalary(company Company) float64 {
	// TODO(human): Calculate and return average salary of all employees
	return 0.0
}

// GetHighestPaid returns the employee with highest salary
func GetHighestPaid(company Company) (Employee, bool) {
	// TODO(human): Return employee with highest salary, or false if no employees
	return Employee{}, false
}

// CountByState returns count of employees per state
func CountByState(company Company) map[string]int {
	// TODO(human): Return map of employee counts by state
	return nil
}
