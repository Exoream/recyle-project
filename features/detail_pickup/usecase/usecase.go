package usecase

import (
	"fmt"
	"log"
	"recycle/features/detail_pickup/entity"
	pickup "recycle/features/pickup/entity"
	rubbish "recycle/features/rubbish/entity"
	user "recycle/features/user/entity"

	"github.com/go-playground/validator/v10"
)

type detailPickupUseCase struct {
	detailPickupRepo entity.DetailPickupDataInterface
	rubbishRepo      rubbish.RubbishDataInterface
	pickupRepo       pickup.PickupDataInterface
	userRepo         user.UserDataInterface
	validate         *validator.Validate
}

func NewDetailPickupUsecase(detailPickupRepo entity.DetailPickupDataInterface, rubbishRepo rubbish.RubbishDataInterface, pickupRepo pickup.PickupDataInterface, userRepo user.UserDataInterface) entity.UseCaseInterface {
	return &detailPickupUseCase{
		detailPickupRepo: detailPickupRepo,
		rubbishRepo:      rubbishRepo,
		pickupRepo:       pickupRepo,
		userRepo:         userRepo,
		validate:         validator.New(),
	}
}

// Create implements entity.UseCaseInterface.
func (uc *detailPickupUseCase) Create(data []entity.Main) error {
	for _, detail := range data {
		errValidate := uc.validate.Struct(detail)
		if errValidate != nil {
			return errValidate
		}
	}

	// Mulai transaksi database
	tx, err := uc.detailPickupRepo.BeginTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	totalUserPoints := float64(0)
	// Simpan pickup IDs untuk perubahan status
	pickupIDs := []string{}
	pickupDataMap := make(map[string]pickup.Main)

	for _, detail := range data {
		pickupID := detail.PickupId
		pickup, exists := pickupDataMap[pickupID]
		if !exists {
			pickup, err := uc.pickupRepo.GetById(pickupID)
			if err != nil {
				return err
			}
			pickupDataMap[pickupID] = pickup
		}

		// Mendapatkan status pickup
		if pickup.Status == "done" {
			return fmt.Errorf("Cannot create detail pickup for a pickup with 'done' status")
		}

		// Hitung total poin berdasarkan itemWeight dan pointPerKg
		rubbish, err := uc.rubbishRepo.GetById(detail.RubbishId)
		if err != nil {
			return err
		}

		detail.TotalPoints = detail.ItemWeight * float64(rubbish.PointPerKg)
		totalUserPoints += detail.TotalPoints
		log.Println(detail.TotalPoints)
		// Buat detail pickup di dalam transaksi
		errCreate := uc.detailPickupRepo.Create([]entity.Main{detail})
		if errCreate != nil {
			return errCreate
		}

		pickupIDs = append(pickupIDs, pickupID)
	}
	log.Println(totalUserPoints)
	// Update status pickup hanya sekali setelah semua detail pickup berhasil dibuat
	for _, pickupID := range pickupIDs {
		_, errUpdateStatus := uc.pickupRepo.UpdateById(pickupID, pickup.Main{Status: "done"})
		if errUpdateStatus != nil {
			return errUpdateStatus
		}
	}

	// Perbarui saldo pengguna setelah menghitung semua detail pickup
	for _, pickup := range pickupDataMap {
		userId := pickup.UserId
		// Mengambil data pengguna
		user, err := uc.userRepo.GetById(userId)
		if err != nil {
			return err
		}

		// Tambahkan total poin dari detail pickup ke saldo pengguna
		user.SaldoPoints += totalUserPoints

		_, errUpdateUser := uc.userRepo.UpdateById(userId, user)
		if errUpdateUser != nil {
			return errUpdateUser
		}
	}

	// Commit transaksi jika semuanya berhasil
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}


// FindAllDetailPickup implements entity.UseCaseInterface.
func (uc *detailPickupUseCase) FindAllDetailPickup() ([]entity.Main, error) {
	detailPickup, err := uc.detailPickupRepo.FindAllDetailPickup()
	if err != nil {
		return nil, err
	}
	return detailPickup, nil
}
