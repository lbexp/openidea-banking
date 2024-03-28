package validation

import (
	"net/url"
	user_model "openidea-banking/src/model/user"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterValidation(validator *validator.Validate) {
	validator.RegisterStructValidation(mustValidRegisterRequest, user_model.UserRegisterRequest{})
	validator.RegisterStructValidation(mustValidLoginRequest, user_model.UserLoginRequest{})
	validator.RegisterValidation("imageurl", mustValidImageUrl)
}

func mustValidRegisterRequest(sl validator.StructLevel) {
	request := sl.Current().Interface().(user_model.UserRegisterRequest)

	email := request.Email
	if !isValidEmail(email) {
		sl.ReportError(request.Email, "Email", "Email", "email", "")
	}
}

func mustValidLoginRequest(sl validator.StructLevel) {
	request := sl.Current().Interface().(user_model.UserLoginRequest)

	email := request.Email
	if !isValidEmail(email) {
		sl.ReportError(request.Email, "Email", "Email", "email", "")
	}
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}

func mustValidImageUrl(fl validator.FieldLevel) bool {
	urlString := fl.Field().String()

	// Parse the URL
	_, err := url.Parse(urlString)
	if err != nil {
		return false
	}

	re := regexp.MustCompile(`(http[s]?:\/\/.*\.(?:png|jpg|gif|svg|jpeg))`)

	return re.MatchString(urlString)
}
