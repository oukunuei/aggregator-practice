package repositories

import (
	"context"
	"time"

	"github.com/golang/mock/gomock"
)

var employeesMap = map[string]func(*gomock.Controller) *MockEmployeeModel{
	"1": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("1").AnyTimes()
		employee.EXPECT().Name().Return("Employee 1").AnyTimes()
		employee.EXPECT().DepartmentID().Return("1").AnyTimes()

		return employee
	},
	"2": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("2").AnyTimes()
		employee.EXPECT().Name().Return("Employee 2").AnyTimes()
		employee.EXPECT().DepartmentID().Return("2").AnyTimes()

		return employee
	},
	"3": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("3").AnyTimes()
		employee.EXPECT().Name().Return("Employee 3").AnyTimes()
		employee.EXPECT().DepartmentID().Return("3").AnyTimes()

		return employee
	},
	"4": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("4").AnyTimes()
		employee.EXPECT().Name().Return("Employee 4").AnyTimes()
		employee.EXPECT().DepartmentID().Return("4").AnyTimes()

		return employee
	},
	"5": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("5").AnyTimes()
		employee.EXPECT().Name().Return("Employee 5").AnyTimes()
		employee.EXPECT().DepartmentID().Return("5").AnyTimes()

		return employee
	},
	"6": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("6").AnyTimes()
		employee.EXPECT().Name().Return("Employee 6").AnyTimes()
		employee.EXPECT().DepartmentID().Return("6").AnyTimes()

		return employee
	},
	"7": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("7").AnyTimes()
		employee.EXPECT().Name().Return("Employee 7").AnyTimes()
		employee.EXPECT().DepartmentID().Return("7").AnyTimes()

		return employee
	},
	"8": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("8").AnyTimes()
		employee.EXPECT().Name().Return("Employee 8").AnyTimes()
		employee.EXPECT().DepartmentID().Return("8").AnyTimes()

		return employee
	},
	"9": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("9").AnyTimes()
		employee.EXPECT().Name().Return("Employee 9").AnyTimes()
		employee.EXPECT().DepartmentID().Return("9").AnyTimes()

		return employee
	},
	"10": func(ctrl *gomock.Controller) *MockEmployeeModel {
		employee := NewMockEmployeeModel(ctrl)
		employee.EXPECT().ID().Return("10").AnyTimes()
		employee.EXPECT().Name().Return("Employee 10").AnyTimes()
		employee.EXPECT().DepartmentID().Return("10").AnyTimes()

		return employee
	},
}

func EmployeeRepo(ctrl *gomock.Controller) EmployeeRepository {
	repo := NewMockEmployeeRepository(ctrl)
	repo.EXPECT().GetEmployee(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, id string) (EmployeeModel, error) {
		return getEmployeeByID(ctrl, id), nil
	}).AnyTimes()

	repo.EXPECT().GetEmployees(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, ids []string) ([]EmployeeModel, error) {
		return getEmployeesByID(ctrl, ids), nil
	}).AnyTimes()

	repo.EXPECT().GetEmployeesWithDepartment(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, ids []string) ([]EmployeeWithDepartmentModel, error) {
		employees := getEmployeesByID(ctrl, ids)
		ret := make([]EmployeeWithDepartmentModel, 0, len(employees))
		for _, employee := range employees {
			item := NewMockEmployeeWithDepartmentModel(ctrl)
			item.EXPECT().ID().Return(employee.ID()).AnyTimes()
			item.EXPECT().Name().Return(employee.Name()).AnyTimes()
			item.EXPECT().DepartmentID().Return(employee.DepartmentID()).AnyTimes()

			departmentName := ""
			departmentFunc := departmentsMap[employee.DepartmentID()]
			if departmentFunc != nil {
				department := departmentFunc(ctrl)
				departmentName = department.Name()

			}

			item.EXPECT().DepartmentName().Return(departmentName).AnyTimes()

			ret = append(ret, item)
		}

		return ret, nil
	}).AnyTimes()

	return repo
}

func getEmployeeByID(ctrl *gomock.Controller, id string) EmployeeModel {
	time.Sleep(time.Duration(50) * time.Millisecond)
	employeeFunc := employeesMap[id]
	if employeeFunc == nil {
		return nil
	}

	return employeeFunc(ctrl)
}

func getEmployeesByID(ctrl *gomock.Controller, ids []string) []EmployeeModel {
	time.Sleep(time.Duration(50) * time.Millisecond)
	ret := make([]EmployeeModel, 0, len(ids))
	for _, id := range ids {
		employeeFunc := employeesMap[id]
		if employeeFunc != nil {
			ret = append(ret, employeeFunc(ctrl))
		}
	}

	return ret
}
