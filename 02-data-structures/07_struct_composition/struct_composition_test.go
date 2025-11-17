package struct_composition

import (
	"reflect"
	"testing"
)

func TestNewEmployee(t *testing.T) {
	tests := []struct {
		name   string
		eName  string
		age    int
		street string
		city   string
		state  string
		zip    string
		salary int
	}{
		{"basic employee", "Alice Smith", 30, "123 Main St", "Boston", "MA", "02101", 75000},
		{"different employee", "Bob Jones", 35, "456 Oak Ave", "Seattle", "WA", "98101", 85000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emp := NewEmployee(tt.eName, tt.age, tt.street, tt.city, tt.state, tt.zip, tt.salary)

			if emp.Name != tt.eName {
				t.Errorf("NewEmployee().Name = %q, want %q", emp.Name, tt.eName)
			}
			if emp.Age != tt.age {
				t.Errorf("NewEmployee().Age = %d, want %d", emp.Age, tt.age)
			}
			if emp.Address.Street != tt.street {
				t.Errorf("NewEmployee().Address.Street = %q, want %q", emp.Address.Street, tt.street)
			}
			if emp.Address.City != tt.city {
				t.Errorf("NewEmployee().Address.City = %q, want %q", emp.Address.City, tt.city)
			}
			if emp.Address.State != tt.state {
				t.Errorf("NewEmployee().Address.State = %q, want %q", emp.Address.State, tt.state)
			}
			if emp.Address.Zip != tt.zip {
				t.Errorf("NewEmployee().Address.Zip = %q, want %q", emp.Address.Zip, tt.zip)
			}
			if emp.Salary != tt.salary {
				t.Errorf("NewEmployee().Salary = %d, want %d", emp.Salary, tt.salary)
			}
		})
	}
}

func TestNewCompany(t *testing.T) {
	tests := []struct {
		name  string
		cName string
	}{
		{"basic company", "TechCorp"},
		{"another company", "StartupInc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			company := NewCompany(tt.cName)

			if company.Name != tt.cName {
				t.Errorf("NewCompany().Name = %q, want %q", company.Name, tt.cName)
			}
			if company.Employees == nil {
				t.Error("NewCompany().Employees should not be nil")
			}
			if len(company.Employees) != 0 {
				t.Errorf("NewCompany().Employees length = %d, want 0", len(company.Employees))
			}
		})
	}
}

func TestAddEmployee(t *testing.T) {
	company := NewCompany("TechCorp")
	emp1 := NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 75000)
	emp2 := NewEmployee("Bob", 35, "456 Oak", "Boston", "MA", "02102", 85000)

	AddEmployee(&company, emp1)
	if len(company.Employees) != 1 {
		t.Fatalf("After adding 1 employee, len = %d, want 1", len(company.Employees))
	}
	if company.Employees[0].Name != "Alice" {
		t.Errorf("First employee name = %q, want %q", company.Employees[0].Name, "Alice")
	}

	AddEmployee(&company, emp2)
	if len(company.Employees) != 2 {
		t.Fatalf("After adding 2 employees, len = %d, want 2", len(company.Employees))
	}
	if company.Employees[1].Name != "Bob" {
		t.Errorf("Second employee name = %q, want %q", company.Employees[1].Name, "Bob")
	}
}

func TestGetEmployeesByCity(t *testing.T) {
	company := NewCompany("TechCorp")
	AddEmployee(&company, NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 75000))
	AddEmployee(&company, NewEmployee("Bob", 35, "456 Oak", "Seattle", "WA", "98101", 85000))
	AddEmployee(&company, NewEmployee("Charlie", 28, "789 Elm", "Boston", "MA", "02103", 70000))

	tests := []struct {
		name      string
		city      string
		wantLen   int
		wantNames []string
	}{
		{"Boston employees", "Boston", 2, []string{"Alice", "Charlie"}},
		{"Seattle employees", "Seattle", 1, []string{"Bob"}},
		{"No matches", "Portland", 0, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetEmployeesByCity(company, tt.city)

			if len(got) != tt.wantLen {
				t.Fatalf("GetEmployeesByCity(%q) length = %d, want %d", tt.city, len(got), tt.wantLen)
			}

			for i, name := range tt.wantNames {
				if got[i].Name != name {
					t.Errorf("GetEmployeesByCity(%q)[%d].Name = %q, want %q", tt.city, i, got[i].Name, name)
				}
			}
		})
	}
}

func TestUpdateSalary(t *testing.T) {
	tests := []struct {
		name       string
		empName    string
		newSalary  int
		wantFound  bool
		wantSalary int
	}{
		{"update Alice", "Alice", 80000, true, 80000},
		{"update Bob", "Bob", 90000, true, 90000},
		{"not found", "David", 100000, false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			company := NewCompany("TechCorp")
			AddEmployee(&company, NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 75000))
			AddEmployee(&company, NewEmployee("Bob", 35, "456 Oak", "Seattle", "WA", "98101", 85000))

			found := UpdateSalary(&company, tt.empName, tt.newSalary)

			if found != tt.wantFound {
				t.Errorf("UpdateSalary(%q) found = %v, want %v", tt.empName, found, tt.wantFound)
			}

			if found {
				// Find the employee and check salary
				for _, emp := range company.Employees {
					if emp.Name == tt.empName {
						if emp.Salary != tt.wantSalary {
							t.Errorf("After UpdateSalary(%q, %d), salary = %d, want %d",
								tt.empName, tt.newSalary, emp.Salary, tt.wantSalary)
						}
						break
					}
				}
			}
		})
	}
}

func TestAverageSalary(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee
		want      float64
	}{
		{
			"two employees",
			[]Employee{
				NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 75000),
				NewEmployee("Bob", 35, "456 Oak", "Seattle", "WA", "98101", 85000),
			},
			80000.0,
		},
		{
			"three employees",
			[]Employee{
				NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 60000),
				NewEmployee("Bob", 35, "456 Oak", "Seattle", "WA", "98101", 90000),
				NewEmployee("Charlie", 28, "789 Elm", "Boston", "MA", "02103", 75000),
			},
			75000.0,
		},
		{
			"single employee",
			[]Employee{
				NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 100000),
			},
			100000.0,
		},
		{
			"empty company",
			[]Employee{},
			0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			company := NewCompany("TechCorp")
			for _, emp := range tt.employees {
				AddEmployee(&company, emp)
			}

			got := AverageSalary(company)
			if got != tt.want {
				t.Errorf("AverageSalary() = %f, want %f", got, tt.want)
			}
		})
	}
}

func TestGetHighestPaid(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee
		wantName  string
		wantFound bool
	}{
		{
			"multiple employees",
			[]Employee{
				NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 75000),
				NewEmployee("Bob", 35, "456 Oak", "Seattle", "WA", "98101", 95000),
				NewEmployee("Charlie", 28, "789 Elm", "Boston", "MA", "02103", 70000),
			},
			"Bob",
			true,
		},
		{
			"single employee",
			[]Employee{
				NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 100000),
			},
			"Alice",
			true,
		},
		{
			"empty company",
			[]Employee{},
			"",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			company := NewCompany("TechCorp")
			for _, emp := range tt.employees {
				AddEmployee(&company, emp)
			}

			got, found := GetHighestPaid(company)

			if found != tt.wantFound {
				t.Errorf("GetHighestPaid() found = %v, want %v", found, tt.wantFound)
			}

			if found && got.Name != tt.wantName {
				t.Errorf("GetHighestPaid() name = %q, want %q", got.Name, tt.wantName)
			}
		})
	}
}

func TestCountByState(t *testing.T) {
	tests := []struct {
		name      string
		employees []Employee
		want      map[string]int
	}{
		{
			"multiple states",
			[]Employee{
				NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 75000),
				NewEmployee("Bob", 35, "456 Oak", "Seattle", "WA", "98101", 85000),
				NewEmployee("Charlie", 28, "789 Elm", "Boston", "MA", "02103", 70000),
				NewEmployee("Diana", 32, "321 Pine", "Portland", "OR", "97201", 80000),
			},
			map[string]int{"MA": 2, "WA": 1, "OR": 1},
		},
		{
			"all same state",
			[]Employee{
				NewEmployee("Alice", 30, "123 Main", "Boston", "MA", "02101", 75000),
				NewEmployee("Bob", 35, "456 Oak", "Cambridge", "MA", "02102", 85000),
			},
			map[string]int{"MA": 2},
		},
		{
			"empty company",
			[]Employee{},
			map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			company := NewCompany("TechCorp")
			for _, emp := range tt.employees {
				AddEmployee(&company, emp)
			}

			got := CountByState(company)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountByState() = %v, want %v", got, tt.want)
			}
		})
	}
}
