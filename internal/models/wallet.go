package models

import "time"

type Wallet struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	UserID    int       `json:"user_id" gorm:"column:user_id;unique"`
	Balance   float64   `json:"balance" gorm:"column:balance;type:decimal(15,2)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*Wallet) TableName() string {
	return "wallets"
}

type WalletTransaction struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	WalletID  int       `json:"wallet_id" gorm:"column:wallet_id"`
	Amount    float64   `json:"amount" gorm:"column:amount;type:decimal(15,2)"`
	Type      string    `json:"type" gorm:"column:type;type:ENUM('CREDIT', 'DEBIT')"`
	Reference string    `json:"reference" gorm:"column:reference;type:varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*WalletTransaction) TableName() string {
	return "wallet_transactions"
}
