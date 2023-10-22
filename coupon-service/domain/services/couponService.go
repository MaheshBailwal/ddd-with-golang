package services

import (
	"strconv"
	"strings"
	"time"

	"github.com/MaheshBailwal/mscore/core"
	"github.com/MaheshBailwal/mscore/persistence"
	"github.com/jinzhu/copier"
	"invafresh.com/coupon-service/domain/constants"
	"invafresh.com/coupon-service/domain/dtos"
	"invafresh.com/coupon-service/domain/entities"
	"invafresh.com/coupon-service/domain/interfaces"
	"invafresh.com/coupon-service/domain/proxies"
	"invafresh.com/coupon-service/domain/queries"
	pb "invafresh.com/coupon-service/infrastructure/remoteServices/department"
)

type CouponService struct {
	UnitOfWork      interfaces.IUnitOfWork
	DepartmentProxy proxies.DepartmentProxy
}

func (c *CouponService) Add(serviceContext core.ServiceContext, couponProfileAddDto *dtos.CouponProfileAdd) (int64, error) {

	timeNow := time.Now()
	couponProfileEntity := entities.CouponProfiles{}
	copier.Copy(&couponProfileEntity, &couponProfileAddDto)
	time := strings.ReplaceAll(timeNow.Format(time.DateOnly), "-", "")
	couponProfileEntity.ModifyDate, _ = strconv.Atoi(time)
	couponProfileEntity.ModifyUser = serviceContext.CurrentUserId

	id, err := c.UnitOfWork.CouponProfileRepository().Create(serviceContext.Ctx, &couponProfileEntity)
	return id, err
}

func (c *CouponService) Update(serviceContext core.ServiceContext, couponProfileUpdateDto *dtos.CouponProfileUpdate) (int64, error) {

	couponProfileEntity := entities.CouponProfiles{}
	copier.Copy(&couponProfileEntity, &couponProfileUpdateDto)

	time := strings.ReplaceAll(time.Now().Format(time.DateOnly), "-", "")
	couponProfileEntity.ModifyDate, _ = strconv.Atoi(time)
	couponProfileEntity.ModifyUser = serviceContext.CurrentUserId

	var id int64
	var err error
	if couponProfileEntity.Id == 0 {
		id, err = c.UnitOfWork.CouponProfileRepository().Create(serviceContext.Ctx, &couponProfileEntity)
	} else {
		id, err = c.UnitOfWork.CouponProfileRepository().Update(serviceContext.Ctx, &couponProfileEntity)
		id = couponProfileEntity.Id
	}
	return id, err
}

func (c *CouponService) GetAll(serviceContext core.ServiceContext, filters persistence.QueryFilter) ([]dtos.CouponProfileDetail, error) {

	entities, _ := c.UnitOfWork.CouponProfileRepository().All(serviceContext.Ctx, filters)
	var couponProfilesDetail []dtos.CouponProfileDetail
	var couponProfileDetail dtos.CouponProfileDetail
	departments := c.DepartmentProxy.GetDepartmets(serviceContext)

	for _, entity := range entities {
		copier.Copy(&couponProfileDetail, &entity)
		couponProfileDetail.DeptName = getDepartmentName(*departments, int32(couponProfileDetail.DeptNo))
		couponProfileDetail.ModifyDate = strconv.Itoa(entity.ModifyDate)
		couponProfilesDetail = append(couponProfilesDetail, couponProfileDetail)
	}

	return couponProfilesDetail, nil
}

func getDepartmentName(dep pb.DepartmentDetails, number int32) string {
	for _, depart := range dep.Department {

		if depart.Number == number {
			return depart.Name
		}
	}

	return ""
}

func (r *CouponService) GetById(serviceContext core.ServiceContext, id int64) (dtos.CouponProfileDetail, error) {
	entity, _ := r.UnitOfWork.CouponProfileRepository().ById(serviceContext.Ctx, id)
	var couponProfileDetail dtos.CouponProfileDetail

	if entity == nil {
		return couponProfileDetail, core.NewServiceError(constants.CouponNotFound, id)
	}

	copier.Copy(&couponProfileDetail, &entity)
	return couponProfileDetail, nil
}

func (r *CouponService) Delete(serviceContext core.ServiceContext, ids []dtos.CouponProfileId) (bool, error) {

	for _, id := range ids {
		qry := queries.CouponInUse{
			CouponKey: id.Id,
		}

		response, _ := r.UnitOfWork.CouponProfileRepository().Query(serviceContext.Ctx, qry)

		if response.(bool) {
			r.UnitOfWork.Rollback()
			return false, core.NewServiceError(constants.CouponInUse, id)
		}

		_, err := r.UnitOfWork.CouponProfileRepository().Delete(serviceContext.Ctx, id.Id)

		if err != nil {
			r.UnitOfWork.Rollback()
			return false, err
		}
	}

	r.UnitOfWork.Commit()
	return true, nil
}
