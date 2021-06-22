package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name        string
	Role        string
	Age         int32
	EmployeCode int64 `copier:"EmployeNum"` // 指定 field name

	// Explicitly ignored in the destination struct.
	// 在目标结构体里明确指定忽略复制
	Salary int
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

// Tags in the destination Struct provide instructions to copier.Copy to ignore
// or enforce copying and to panic or return an error if a field was not copied.
// 目标结构体里tag指定must must,nopanic  - 三种赋值类型情况
type Employee struct {
	// Tell copier.Copy to panic if this field is not copied.
	// 告诉copier.Copy若该字段未复制则panic
	Name string `copier:"must"`

	// Tell copier.Copy to return an error if this field is not copied.
	// 告诉copier.Copy若该字段未复制则返回error
	Age int32 `copier:"must,nopanic"`

	// Tell copier.Copy to explicitly ignore copying this field.
	// 告诉copier.Copy若该字段不要复制
	Salary int `copier:"-"`

	DoubleAge int32
	EmployeId int64 `copier:"EmployeNum"` //指定 field name
	SuperRole string
}

func (employee *Employee) Role(role string) {
	employee.SuperRole = "Super " + role
}

func main() {
	var (
		user      = User{Name: "Paopao", Age: 18, Role: "Programer", Salary: 666000}
		users     = []User{{Name: "Paopao", Age: 18, Role: "Programer", Salary: 100000}, {Name: "Paopao 2", Age: 30, Role: "Dev", Salary: 60000}}
		employee  = Employee{Salary: 150000}
		employees = []Employee{}
	)

	copier.Copy(&employee, &user)

	fmt.Printf("%#v \n", employee)
	// Employee{
	//    Name: "Paopao",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    Salary:150000,            // Copying explicitly ignored
	//    DoubleAge: 36,            // Copy from method
	//    EmployeeId: 0,            // Ignored
	//    SuperRole: "Super Programer", // Copy to method
	// }

	// Copy struct to slice
	// 拷贝结构体至切片
	copier.Copy(&employees, &user)
	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Paopao", Age: 18, Salary:0, DoubleAge: 36, EmployeId: 0, SuperRole: "Super Programer"}
	// }

	// Copy slice to slice
	// 拷贝切片至切片
	employees = []Employee{}
	copier.Copy(&employees, &users)
	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Paopao", Age: 18, Salary:0, DoubleAge: 36, EmployeId: 0, SuperRole: "Super Programer"},
	//   {Name: "Paopao 2", Age: 30, Salary:0, DoubleAge: 60, EmployeId: 0, SuperRole: "Super Dev"},
	// }

	// Copy map to map
	// 拷贝map至map
	map1 := map[int]int{1: 8, 8: 64}
	map2 := map[int32]int8{}
	copier.Copy(&map2, map1)

	fmt.Printf("%#v \n", map2)
	// map[int32]int8{1:8, 8:64}
}

//readme:
//1 . go get github.com/Paopao/copier
//2.https://github.com/Paopao/copier
//3. go mod init github.com/wsqyouth/coopers_go_code; go mod tidy
