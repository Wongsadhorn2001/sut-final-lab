package entity

import {
	"fmt"
	"gorm.io/gorm"
	"testing"

	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
}

type Customer struct{
	gorm.Model
	Name string
	Email string
	Customer string
}

func PositiveTest(t *testing T){
	g := NewGomegaWithT(t)

	t.Run("Positive Test", func(t *testing T)){
		custome := Customer{
			Name: "Minics",
			Email: "minic2001@gmail.com",
			Customer: "L1234567" || "M1234567" || "H1234567",
		}

		ok, err := govalidator.validateStruct(custome)
		g.expect(ok).To(BeTrue())
		g.expect(err).To(BeNil())
		fmt.println(err)
	}
}

func NegativeNameTest(t *testing T){
	g := NewGomegaWithT(t)

	t.Run("Negative Name Test", func(t *testing T)){
		custome := Customer{
			Name: "", //wrong
			Email: "minic2001@gmail.com",
			Customer: "L1234567" || "M1234567" || "H1234567",
		}

		ok, err := govalidator.validateStruct(custome)
		g.expect(ok).ToNot(BeTrue())
		g.expect(err).ToNot(BeNil())
		g.expect(err.Error()).To(Equal("กรุณากรอกชื่อ"))
	}
}