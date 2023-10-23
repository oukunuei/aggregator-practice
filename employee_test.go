package aggregatorpractice

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	repo "github.com/oukunuei/aggregator-practice/repositories"
)

func Test_ShowEmployeeList_Example1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	employeeRepo := repo.EmployeeRepo(ctrl)
	departmentRepo := repo.DepartmentRepo(ctrl)
	employeeIDs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	employees, _ := employeeRepo.GetEmployees(ctx, employeeIDs)
	for _, employee := range employees {
		department, _ := departmentRepo.GetDepartment(ctx, employee.DepartmentID())
		if department == nil {
			continue
		}

		t.Log(employee.Name() + " " + department.Name())
	}
}

func Test_ShowEmployeeList_Example2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	employeeRepo := repo.EmployeeRepo(ctrl)
	departmentRepo := repo.DepartmentRepo(ctrl)
	employeeIDs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	employees, _ := employeeRepo.GetEmployees(ctx, employeeIDs)
	departmentIDs := make([]string, 0, len(employeeIDs))
	for _, employee := range employees {
		departmentIDs = append(departmentIDs, employee.DepartmentID())
	}

	departmentsMap := make(map[string]repo.DepartmentModel)
	departments, _ := departmentRepo.GetDepartments(ctx, departmentIDs)
	for _, department := range departments {
		departmentsMap[department.ID()] = department
	}

	for _, employee := range employees {
		relatedDepartment := departmentsMap[employee.DepartmentID()]
		if relatedDepartment == nil {
			continue
		}

		t.Log(employee.Name() + " " + relatedDepartment.Name())
	}
}

func Test_ShowEmployeeList_Example3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	employeeRepo := repo.EmployeeRepo(ctrl)
	employeeIDs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	// SELECT e.id as employee_id, e.name as employee.name, d.id as department_id, d.name as department_name
	// FROM employees e join departments d on e.department_id = d.id
	// WHERE e.id in ("1", "2", "3", "4", "5", "6", "7", "8", "9", "10");
	employees, _ := employeeRepo.GetEmployeesWithDepartment(ctx, employeeIDs)

	for _, employee := range employees {
		t.Log(employee.Name() + " " + employee.DepartmentName())
	}
}

func Test_ShowEmployeeList_Example4(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	employeeRepo := repo.EmployeeRepo(ctrl)
	departmentRepo := repo.DepartmentRepo(ctrl)
	employeeIDs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	employees, _ := employeeRepo.GetEmployees(ctx, employeeIDs)
	for _, employee := range employees {
		department, _ := departmentRepo.GetDepartment(ctx, employee.DepartmentID())
		if department == nil {
			continue
		}

		t.Log(employee.Name() + " " + department.Name())
	}
}

func sleepAndPrintHelloWorld() {
	time.Sleep(3 * time.Second)
	println("Hello World")
}

func sleepAndPrintHelloWorldWithWaitGroup(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	println("Hello World")
}

func Test_Goroutine_Example1(t *testing.T) {
	sleepAndPrintHelloWorld()
	sleepAndPrintHelloWorld()
	sleepAndPrintHelloWorld()
	sleepAndPrintHelloWorld()
	sleepAndPrintHelloWorld()
	sleepAndPrintHelloWorld()
}

func Test_Goroutine_Example2(t *testing.T) {
	go sleepAndPrintHelloWorld()
	go sleepAndPrintHelloWorld()
	go sleepAndPrintHelloWorld()
	go sleepAndPrintHelloWorld()
	go sleepAndPrintHelloWorld()
	go sleepAndPrintHelloWorld()
}

func Test_Goroutine_Example3(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go sleepAndPrintHelloWorldWithWaitGroup(&wg)
	wg.Add(1)
	go sleepAndPrintHelloWorldWithWaitGroup(&wg)
	wg.Add(1)
	go sleepAndPrintHelloWorldWithWaitGroup(&wg)
	wg.Add(1)
	go sleepAndPrintHelloWorldWithWaitGroup(&wg)
	wg.Add(1)
	go sleepAndPrintHelloWorldWithWaitGroup(&wg)
	wg.Add(1)
	go sleepAndPrintHelloWorldWithWaitGroup(&wg)

	wg.Wait()
}

func modifyInputWorker(inputCh chan string, outputCh chan string) {
	for inputStr := range inputCh {
		time.Sleep(3 * time.Second)
		outputCh <- inputStr + " is coming"
	}

}

func Test_Channel_Example(t *testing.T) {
	inputCh := make(chan string, 10)
	outputCh := make(chan string, 10)
	workerNum := 5
	inputCount := 10
	for i := 0; i < workerNum; i++ {
		go modifyInputWorker(inputCh, outputCh)
	}

	for i := 0; i < inputCount; i++ {
		inputCh <- fmt.Sprint(i)
	}

	close(inputCh)

	for i := 0; i < inputCount; i++ {
		outputStr := <-outputCh
		println(outputStr)
	}
}
