package repositories

import (
	"context"
	"time"

	"github.com/golang/mock/gomock"
)

var departmentsMap = map[string]func(*gomock.Controller) *MockDepartmentModel{
	"1": func(ctrl *gomock.Controller) *MockDepartmentModel {
		department := NewMockDepartmentModel(ctrl)
		department.EXPECT().ID().Return("1").AnyTimes()
		department.EXPECT().Name().Return("Department 1").AnyTimes()

		return department
	},
	"2": func(ctrl *gomock.Controller) *MockDepartmentModel {
		department := NewMockDepartmentModel(ctrl)
		department.EXPECT().ID().Return("2").AnyTimes()
		department.EXPECT().Name().Return("Department 2").AnyTimes()

		return department
	},
	"3": func(ctrl *gomock.Controller) *MockDepartmentModel {
		department := NewMockDepartmentModel(ctrl)
		department.EXPECT().ID().Return("3").AnyTimes()
		department.EXPECT().Name().Return("Department 3").AnyTimes()

		return department
	},
	"4": func(ctrl *gomock.Controller) *MockDepartmentModel {
		department := NewMockDepartmentModel(ctrl)
		department.EXPECT().ID().Return("4").AnyTimes()
		department.EXPECT().Name().Return("Department 4").AnyTimes()

		return department
	},
	"5": func(ctrl *gomock.Controller) *MockDepartmentModel {
		image := NewMockDepartmentModel(ctrl)
		image.EXPECT().ID().Return("5").AnyTimes()
		image.EXPECT().Name().Return("Department 5").AnyTimes()

		return image
	},
	"6": func(ctrl *gomock.Controller) *MockDepartmentModel {
		department := NewMockDepartmentModel(ctrl)
		department.EXPECT().ID().Return("6").AnyTimes()
		department.EXPECT().Name().Return("Department 6").AnyTimes()

		return department
	},
	"7": func(ctrl *gomock.Controller) *MockDepartmentModel {
		department := NewMockDepartmentModel(ctrl)
		department.EXPECT().ID().Return("7").AnyTimes()
		department.EXPECT().Name().Return("Department 7").AnyTimes()

		return department
	},
	"8": func(ctrl *gomock.Controller) *MockDepartmentModel {
		department := NewMockDepartmentModel(ctrl)
		department.EXPECT().ID().Return("8").AnyTimes()
		department.EXPECT().Name().Return("Department 8").AnyTimes()

		return department
	},
	"9": func(ctrl *gomock.Controller) *MockDepartmentModel {
		department := NewMockDepartmentModel(ctrl)
		department.EXPECT().ID().Return("9").AnyTimes()
		department.EXPECT().Name().Return("Department 9").AnyTimes()

		return department
	},
	"10": func(ctrl *gomock.Controller) *MockDepartmentModel {
		department := NewMockDepartmentModel(ctrl)
		department.EXPECT().ID().Return("10").AnyTimes()
		department.EXPECT().Name().Return("Department 10").AnyTimes()

		return department
	},
}

func DepartmentRepo(ctrl *gomock.Controller) DepartmentRepository {
	repo := NewMockDepartmentRepository(ctrl)
	repo.EXPECT().GetDepartment(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, id string) (DepartmentModel, error) {
		return getDepartmentByID(ctrl, id), nil
	}).AnyTimes()

	repo.EXPECT().GetDepartments(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, ids []string) ([]DepartmentModel, error) {
		return getDepartmentsByID(ctrl, ids), nil
	}).AnyTimes()

	return repo
}

func getDepartmentByID(ctrl *gomock.Controller, id string) DepartmentModel {
	time.Sleep(time.Duration(50) * time.Millisecond)
	departmentFunc := departmentsMap[id]
	if departmentFunc == nil {
		return nil
	}

	return departmentFunc(ctrl)
}

func getDepartmentsByID(ctrl *gomock.Controller, ids []string) []DepartmentModel {
	time.Sleep(time.Duration(50) * time.Millisecond)
	ret := make([]DepartmentModel, 0, len(ids))
	for _, id := range ids {
		departmentFunc := departmentsMap[id]
		if departmentFunc != nil {
			ret = append(ret, departmentFunc(ctrl))
		}
	}

	return ret
}
