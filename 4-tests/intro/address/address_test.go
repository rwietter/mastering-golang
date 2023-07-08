package address_test

import (
	"address/address"
	"testing"
)

func TestAddressType(t *testing.T) {
	t.Run("should return 'Rua' when address is 'Rua 1'", func(t *testing.T) {
		addressType := address.AddressType("Rua 1")

		if addressType != "Rua" {
			t.Errorf("expected 'Rua', got %s", addressType)
		}
	})

	t.Run("should return 'Avenida' when address is 'Avenida 2'", func(t *testing.T) {
		addressType := address.AddressType("Avenida 2")

		if addressType != "Avenida" {
			t.Errorf("expected 'Avenida', got %s", addressType)
		}
	})

	t.Run("should return 'Estrada' when address is 'Estrada 3'", func(t *testing.T) {
		addressType := address.AddressType("Estrada 3")

		if addressType != "Estrada" {
			t.Errorf("expected 'Estrada', got %s", addressType)
		}
	})
}
