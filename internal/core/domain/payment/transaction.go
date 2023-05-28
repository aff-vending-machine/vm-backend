package payment

import (
	"context"
	"time"
	"vm-backend/internal/core/domain/machine"
	"vm-backend/internal/core/domain/store"
	"vm-backend/internal/core/infrastructure/strorage/postgresql/service"
	"vm-backend/internal/layer/usecase/payment_transaction/request"
	"vm-backend/internal/layer/usecase/payment_transaction/response"

	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	ID                 uint            `json:"id" gorm:"primarykey"`
	BranchID           uint            `json:"branch_id"`
	Branch             store.Branch    `json:"branch"`
	MachineID          uint            `json:"machine_id"`
	Machine            machine.Machine `json:"machine"`
	ChannelID          uint            `json:"channel_id"`
	Channel            Channel         `json:"channel"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	MerchantOrderID    string          `json:"merchant_order_id" gorm:"not null,uniqueIndex"` // key to find order
	RawCart            string          `json:"raw_cart"`                                      // ordered
	Note               string          `json:"note"`                                          // ordered
	OrderQuantity      int             `json:"order_quantity" gorm:"not null"`                // ordered
	OrderPrice         float64         `json:"order_price" gorm:"not null"`                   // ordered
	OrderStatus        string          `json:"order_status" gorm:"not null"`                  // ordered
	OrderedAt          time.Time       `json:"ordered_at"`
	PaymentRequestedAt *time.Time      `json:"payment_requested_at"`                  // ordered - MakeTransactionCreateRequest
	RawReference       *string         `json:"raw_reference"`                         // raw_reference
	Reference1         *string         `json:"reference1"`                            // reference1 - MakeTransactionCreateResult
	Reference2         *string         `json:"reference2"`                            // reference2
	Reference3         *string         `json:"reference3"`                            // reference3
	CancelledBy        *string         `json:"cancelled_by"`                          // cancelled - MakeTransactionCancel
	CancelledAt        *time.Time      `json:"cancelled_at"`                          // cancelled - MakeTransactionCancel
	ConfirmedPaidBy    *string         `json:"confirmed_paid_by" gorm:"default:null"` // paid - MakeTransactionPaid
	ConfirmedPaidAt    *time.Time      `json:"confirmed_paid_at" gorm:"default:null"` // paid - MakeTransactionPaid
	RefError           *string         `json:"ref_error"`                             // MakeTransactionError
	RefundAt           *time.Time      `json:"refund_at" gorm:"default:null"`         // refund
	RefundPrice        float64         `json:"refund_price"`                          // refund
	ReceivedItemAt     *time.Time      `json:"received_item_at" gorm:"default:null"`  // received - MakeTransactionDone
	ReceivedQuantity   int             `json:"received_quantity"`                     // received, refund - MakeTransactionDone
	PaidPrice          float64         `json:"paid_price"`                            // received, refund - MakeTransactionDone
	IsError            bool            `json:"is_error" gorm:"default:false"`         // error
	Error              *string         `json:"error" gorm:"default:null"`             // error - MakeTransactionError
	ErrorAt            *time.Time      `json:"error_at" gorm:"default:null"`          // MakeTransactionRefund
}

func (e *Transaction) TableName() string {
	return "payment_transactions"
}

type TransactionRepository interface {
	service.Repository[Transaction]
}

type TransactionUsecase interface {
	Count(context.Context, *request.Filter) (int64, error)
	Get(context.Context, *request.Get) (*response.Transaction, error)
	List(context.Context, *request.Filter) ([]response.Transaction, error)
	Done(context.Context, *request.Done) error
	Cancel(context.Context, *request.Cancel) error
}

type TransactionTransport interface {
	Read(ctx *fiber.Ctx) error    // GET	{transactions}
	Count(ctx *fiber.Ctx) error   // GET 	{transactions/count}
	ReadOne(ctx *fiber.Ctx) error // GET 	{transactions/:id}
	Done(ctx *fiber.Ctx) error    // POST 	{transactions/:id/done}
	Cancel(ctx *fiber.Ctx) error  // POST 	{transactions/:id/cancel}
}
