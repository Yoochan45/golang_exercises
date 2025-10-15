package entity

type Orders struct {
	ID         int
	UserID     int
	TotalPrice float64
}

type OrderItems struct {
	ProductName string
	OrderID     int
	ProductID   int
	Quantity    int
	Price       float64
}
