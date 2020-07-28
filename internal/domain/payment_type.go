package domain

type PaymentType string

//noinspection ALL
const (
	CASH            PaymentType = "cash"
	CREDIT_CARD     PaymentType = "credit_card"
	DEBIT_CARD      PaymentType = "debit_card"
	PERSONAL_CHEQUE PaymentType = "personal_cheque"
	TRANSFER        PaymentType = "transfer"
)
