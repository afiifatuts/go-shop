package models

//disini nantinya method untuk meregister models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	//meregister model apa saja yang mau di migration ke method initializeDb yang ada di server.go
	return []Model{
		{Model: User{}},
		{Model: Address{}},
		{Model: Product{}},
		{Model: ProductImage{}},
		{Model: Section{}},
		{Model: Category{}},
		{Model: Order{}},
		{Model: OrderItem{}},
		{Model: OrderCustomer{}},
		{Model: Payment{}},
		{Model: Shipment{}},
	}
}
