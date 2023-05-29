package migration

import (
	"time"
	"vm-backend/internal/core/domain/payment"
	"vm-backend/internal/core/domain/store"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                  uint       `json:"id" gorm:"primarykey"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	MachineID           uint       `json:"machine_id"`
	MachineName         string     `json:"machine_name"`
	MerchantOrderID     string     `json:"merchant_order_id" gorm:"not null,uniqueIndex"` // key to find order
	MachineSerialNumber string     `json:"machine_serial_number" gorm:"not null"`         // key to find machine
	Branch              string     `json:"branch"`                                        // ordered
	Location            string     `json:"location"`                                      // ordered
	RawCart             string     `json:"raw_cart"`                                      // ordered
	Note                string     `json:"note"`                                          // ordered
	OrderQuantity       int        `json:"order_quantity" gorm:"not null"`                // ordered
	OrderPrice          float64    `json:"order_price" gorm:"not null"`                   // ordered
	OrderStatus         string     `json:"order_status" gorm:"not null"`                  // ordered
	OrderedAt           time.Time  `json:"ordered_at"`                                    // ordered
	PaymentChannel      string     `json:"payment_channel"`                               // ordered, key to find payment channel - MakeTransactionCreateRequest
	PaymentRequestedAt  *time.Time `json:"payment_requested_at"`                          // ordered - MakeTransactionCreateRequest
	RawReference        *string    `json:"raw_reference"`                                 // raw_reference
	Reference1          *string    `json:"reference1"`                                    // reference1 - MakeTransactionCreateResult
	Reference2          *string    `json:"reference2"`                                    // reference2
	Reference3          *string    `json:"reference3"`                                    // reference3
	CancelledBy         *string    `json:"cancelled_by"`                                  // cancelled - MakeTransactionCancel
	CancelledAt         *time.Time `json:"cancelled_at"`                                  // cancelled - MakeTransactionCancel
	ConfirmedPaidBy     *string    `json:"confirmed_paid_by" gorm:"default:null"`         // paid - MakeTransactionPaid
	ConfirmedPaidAt     *time.Time `json:"confirmed_paid_at" gorm:"default:null"`         // paid - MakeTransactionPaid
	RefError            *string    `json:"ref_error"`                                     // MakeTransactionError
	RefundAt            *time.Time `json:"refund_at" gorm:"default:null"`                 // refund
	RefundPrice         float64    `json:"refund_price"`                                  // refund
	ReceivedItemAt      *time.Time `json:"received_item_at" gorm:"default:null"`          // received - MakeTransactionDone
	ReceivedQuantity    int        `json:"received_quantity"`                             // received, refund - MakeTransactionDone
	PaidPrice           float64    `json:"paid_price"`                                    // received, refund - MakeTransactionDone
	IsError             bool       `json:"is_error" gorm:"default:false"`                 // error
	Error               *string    `json:"error" gorm:"default:null"`                     // error - MakeTransactionError
	ErrorAt             *time.Time `json:"error_at" gorm:"default:null"`                  // MakeTransactionRefund
}

func (e *Transaction) TableName() string {
	return "transactions"
}

func MigrateTransaction(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})

	var transactions []Transaction
	db.Find(&transactions)

	for _, transaction := range transactions {
		var paymentTransaction payment.Transaction

		var branch store.Branch
		if isNotFound(db.Where("name = ?", transaction.Branch), &branch) {
			branch = store.Branch{Name: transaction.Branch}
			db.Debug().Create(&branch)
			log.Info().Str("name", transaction.Branch).Msg("migrated branch")
		}

		var channel payment.Channel
		if transaction.PaymentChannel != "" {
			if isNotFound(db.Where("channel = ?", transaction.PaymentChannel), &channel) {
				channel = payment.Channel{
					Name:       transaction.PaymentChannel,
					Channel:    transaction.PaymentChannel,
					Vendor:     "",
					IsEnable:   true,
					Host:       "",
					MerchantID: "",
					BillerCode: "",
					BillerID:   "",
					Token:      "",
					StoreID:    "",
					TerminalID: "",
				}
				db.Debug().Create(&channel)
				log.Info().Str("channel", transaction.PaymentChannel).Msg("migrated channel")
			}
		} else {
			continue
		}

		// Check if the user already exists in the new table
		if isNotFound(db.Where("merchant_order_id = ?", transaction.MerchantOrderID), &paymentTransaction) {
			paymentTransaction = payment.Transaction{
				BranchID:           branch.ID,
				MachineID:          transaction.MachineID,
				ChannelID:          channel.ID,
				MerchantOrderID:    transaction.MerchantOrderID,
				RawCart:            transaction.RawCart,
				OrderQuantity:      transaction.OrderQuantity,
				OrderPrice:         transaction.OrderPrice,
				OrderStatus:        transaction.OrderStatus,
				OrderedAt:          transaction.OrderedAt,
				PaymentRequestedAt: transaction.PaymentRequestedAt,
				RawReference:       transaction.RawReference,
				Reference1:         transaction.Reference1,
				Reference2:         transaction.Reference2,
				Reference3:         transaction.Reference3,
				CancelledBy:        transaction.CancelledBy,
				CancelledAt:        transaction.CancelledAt,
				ConfirmedPaidBy:    transaction.ConfirmedPaidBy,
				ConfirmedPaidAt:    transaction.ConfirmedPaidAt,
				RefError:           transaction.RefError,
				RefundAt:           transaction.RefundAt,
				RefundPrice:        transaction.RefundPrice,
				ReceivedItemAt:     transaction.ReceivedItemAt,
				ReceivedQuantity:   transaction.ReceivedQuantity,
				PaidPrice:          transaction.PaidPrice,
				IsError:            transaction.IsError,
				Error:              transaction.Error,
				ErrorAt:            transaction.ErrorAt,
			}
			db.Debug().Create(&paymentTransaction)
			log.Info().Str("merchant_order_id", paymentTransaction.MerchantOrderID).Msg("migrated transaction")
		}
	}
}
