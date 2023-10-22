package services

import (
	"strconv"
	"strings"
	"time"

	"github.com/MaheshBailwal/mscore/core"
	"github.com/MaheshBailwal/mscore/persistence"
	"github.com/jinzhu/copier"
	dtos "invafresh.com/department-service/domain/dtos"
	"invafresh.com/department-service/domain/entities"
	"invafresh.com/department-service/domain/interfaces"
)

type DepartmentService struct {
	UnitOfWork interfaces.IUnitOfWork
}

func (r *DepartmentService) Update(serviceContext core.ServiceContext, departmentSetDto *dtos.DepartmentSet) (int64, error) {

	departmentEntity := entities.Departments{}
	copier.Copy(&departmentEntity, &departmentSetDto)

	time := strings.ReplaceAll(time.Now().Format(time.DateOnly), "-", "")
	departmentEntity.ModifyDate, _ = strconv.Atoi(time)
	departmentEntity.ModifyUser = serviceContext.CurrentUserId

	var err error
	department, _ := r.UnitOfWork.DepartmentRepository().ById(serviceContext.Ctx, departmentSetDto.Number)

	if department == nil {
		_, err = r.UnitOfWork.DepartmentRepository().Create(serviceContext.Ctx, &departmentEntity)
	} else {
		_, err = r.UnitOfWork.DepartmentRepository().Update(serviceContext.Ctx, &departmentEntity)
	}

	r.UnitOfWork.Conclude(err)

	return departmentEntity.Number, err
}

func (r *DepartmentService) GetAll(serviceContext core.ServiceContext, filters persistence.QueryFilter) ([]dtos.DepartmentDetail, error) {

	entities, _ := r.UnitOfWork.DepartmentRepository().All(serviceContext.Ctx, filters)
	var departmentsDetail []dtos.DepartmentDetail

	for _, entity := range entities {
		departmentDetail := dtos.DepartmentDetail{}
		copier.Copy(&departmentDetail, &entity)
		departmentDetail.ModifyDate = strconv.Itoa(entity.ModifyDate)
		departmentsDetail = append(departmentsDetail, departmentDetail)
	}

	return departmentsDetail, nil
}

func (r *DepartmentService) Delete(serviceContext core.ServiceContext, departNumbers []dtos.DepartmentNumber) (bool, error) {

	for _, no := range departNumbers {
		_, err := r.UnitOfWork.DepartmentRepository().Delete(serviceContext.Ctx, no.Number)

		if err != nil {
			r.UnitOfWork.Rollback()
			return false, err
		}
	}

	r.UnitOfWork.Commit()
	return true, nil
}
