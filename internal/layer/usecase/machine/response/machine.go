package response

import (
	"time"
)

type Machine struct {
	ID                  uint       `json:"id"`
	Name                string     `json:"name"`
	SerialNumber        string     `json:"serial_number"`
	Location            string     `json:"location"`
	Type                string     `json:"type"`
	Vendor              string     `json:"vendor"`
	Slots               []Slot     `json:"slots"`
	Status              string     `json:"status"`
	SyncTime            *time.Time `json:"sync_time"`
	SyncSlotTime        *time.Time `json:"sync_slot_time"`
	SyncTransactionTime *time.Time `json:"sync_transaction_time"`
}
