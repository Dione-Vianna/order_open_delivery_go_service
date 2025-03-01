package validation

import (
	proto "order_open_delivery_go_service/api/order-proto"

	"github.com/go-playground/validator"
)

type ItemValidation struct {
	Id                string  `validate:"required"`
	Index             int32   `validate:"required,gt=0"`
	Name              string  `validate:"required"`
	ExternalCode      string  `validate:"required"`
	Unit              string  `validate:"required"`
	Ean               string  `validate:"required"`
	Quantity          int32   `validate:"required,gt=0"`
	SpecialInstructions string `validate:"omitempty"`
	UnitPriceValue    float32 `validate:"required,gt=0"`
	OptionsPriceValue float32 `validate:"required,gt=0"`
	TotalPriceValue   float32 `validate:"required,gt=0"`
}

type OrderRequestValidation struct {
	Id    string        `validate:"required"`
	Items []*proto.Item `validate:"required,dive"` 
}

func NewOrderRequestValidation() *OrderRequestValidation {
	return &OrderRequestValidation{}
}

func NewItemValidation() *ItemValidation {
	return &ItemValidation{}
}

func ValidateOrderRequest(orderRequest *proto.OrderRequest) error {
	validate := validator.New()

	// OrderRequest Validation
	orderValidation := &OrderRequestValidation{
		Id:    orderRequest.Id,
		Items: orderRequest.Items,
	}

	// Validates the OrderRequest itself
	if err := validate.Struct(orderValidation); err != nil {
		return err
	}

	// Validation of items (each item in the list)
	for _, item := range orderRequest.Items {
		itemValidation := &ItemValidation{
			Id:                item.Id,
			Index:             item.Index,
			Name:              item.Name,
			ExternalCode:      item.ExternalCode,
			Unit:              item.Unit,
			Ean:               item.Ean,
			Quantity:          item.Quantity,
			SpecialInstructions: item.SpecialInstructions,
			UnitPriceValue:    item.UnitPrice.Value,
			OptionsPriceValue: item.OptionsPrice.Value,
			TotalPriceValue:   item.TotalPrice.Value,
		}

		if err := validate.Struct(itemValidation); err != nil {
			return err
		}
	}

	return nil
}
