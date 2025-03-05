package validation

import (
	"fmt"
	proto "order_open_delivery_go_service/api/order-proto"

	"github.com/go-playground/validator"
)

type ItemValidation struct {
	Id                  string  `validate:"required"`
	Index               int32   `validate:"required,gt=0"`
	Name                string  `validate:"required"`
	ExternalCode        string  `validate:"required"`
	Unit                string  `validate:"required"`
	Ean                 string  `validate:"required"`
	Quantity            int32   `validate:"required,gt=0"`
	SpecialInstructions string  `validate:"omitempty"`
	UnitPriceValue      float32 `validate:"required,gt=0"`
	OptionsPriceValue   float32 `validate:"required,gt=0"`
	TotalPriceValue     float32 `validate:"required,gt=0"`
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

	if orderRequest == nil {
		return fmt.Errorf("orderRequest is nil")
	}

	if orderRequest.Items == nil {
		return fmt.Errorf("orderRequest.Items is nil")
	}

	orderValidation := &OrderRequestValidation{
		Id:    orderRequest.Id,
		Items: orderRequest.Items,
	}

	if err := validate.Struct(orderValidation); err != nil {
		return err
	}

	for _, item := range orderRequest.Items {
		if item == nil {
			return fmt.Errorf("found nil item in orderRequest.Items")
		}

		var unitPriceValue, optionsPriceValue, totalPriceValue float32

		if item.UnitPrice != nil {
			unitPriceValue = item.UnitPrice.Value
		} else {
			return fmt.Errorf("item.UnitPrice is nil")
		}

		if item.OptionsPrice != nil {
			optionsPriceValue = item.OptionsPrice.Value
		} else {
			return fmt.Errorf("item.OptionsPrice is nil")
		}

		if item.TotalPrice != nil {
			totalPriceValue = item.TotalPrice.Value
		} else {
			return fmt.Errorf("item.TotalPrice is nil")
		}

		itemValidation := &ItemValidation{
			Id:                  item.Id,
			Index:               item.Index,
			Name:                item.Name,
			ExternalCode:        item.ExternalCode,
			Unit:                item.Unit,
			Ean:                 item.Ean,
			Quantity:            item.Quantity,
			SpecialInstructions: item.SpecialInstructions,
			UnitPriceValue:      unitPriceValue,
			OptionsPriceValue:   optionsPriceValue,
			TotalPriceValue:     totalPriceValue,
		}

		if err := validate.Struct(itemValidation); err != nil {
			return err
		}
	}

	return nil
}
