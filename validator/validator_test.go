package validator

import "testing"

type UserRequest struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func TestValidatorSuccess(t *testing.T) {
	validate := New()

	payload := UserRequest{
		Name:  "Raka",
		Email: "raka@mail.com",
	}

	err := validate.Validate(payload)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidatorFailed(t *testing.T) {
	validate := New()

	payload := UserRequest{
		Name:  "",
		Email: "invalid-email",
	}

	err := validate.Validate(payload)
	if err == nil {
		t.Fatal("expected validation error")
	}
}
