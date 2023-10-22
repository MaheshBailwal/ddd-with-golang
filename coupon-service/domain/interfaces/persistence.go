package interfaces

import (
	"github.com/MaheshBailwal/mscore/persistence"
	"invafresh.com/coupon-service/domain/entities"
)

//go:generate mockgen -destination=../mocks/mockPersistence.go -package=mocks github.com/ehsandavari/go-clean-architecture/application/common/interfaces IUnitOfWork

type (
	ICouponProfileRepository interface {
		persistence.IGenericRepository[entities.CouponProfiles]
	}
	IUnitOfWork interface {
		CouponProfileRepository() ICouponProfileRepository
		Commit()
		Rollback()
		Conclude(error)
	}
)
