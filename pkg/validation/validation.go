package validation

import (
	"abude-backend/pkg/exception"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"
	"text/template"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Validation struct {
	*validator.Validate
}

func (v *Validation) Setup(db *gorm.DB) {
	validate := validator.New()

	validate.RegisterValidation("exist", exist(db))
	validate.RegisterValidation("not_exist", unique(db))

	v.Validate = validate
}

func New(db *gorm.DB) *Validation {
	validation := new(Validation)
	validation.Setup(db)

	return validation
}

// Body parses the request body into a given struct and validates it.
func (v *Validation) Body(data interface{}, ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(data); err != nil {
		return exception.Validation()
	}

	return v.ValidateStruct(data)
}

// Query parses the query parameters into a given struct and validates it.
func (v *Validation) Query(data interface{}, ctx *fiber.Ctx) error {
	if err := ctx.QueryParser(data); err != nil {
		return exception.Validation()
	}

	return v.ValidateStruct(data)
}

// Retrieves an integer value from the Fiber context using the specified key.
func (v *Validation) ParamsInt(ctx *fiber.Ctx, keys ...string) (int, error) {
	key := "id"
	if len(keys) > 0 && keys[0] != "" {
		key = keys[0]
	}

	value, err := ctx.ParamsInt(key)
	if err != nil {
		return -1, exception.BadRequest("ID tidak valid")
	}

	return value, nil
}

func (v *Validation) Field(field interface{}, tag string) error {
	if err := v.Var(field, tag); err != nil {
		return fmt.Errorf("%v", getMessage(err.(validator.FieldError)))
	}

	return nil
}

// Body parses the request form data into a given struct and validates it.
func (v *Validation) FormData(data interface{}, ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(data); err != nil {
		return exception.Validation()
	}

	form, _ := ctx.MultipartForm()
	if form == nil {
		return v.ValidateStruct(data)
	}

	structType := reflect.Indirect(reflect.ValueOf(data)).Type() // Get indirect type
	structValue := reflect.ValueOf(data).Elem()                  // Get the value of the struct

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Type == reflect.TypeOf(&multipart.FileHeader{}) {
			fieldName := field.Tag.Get("form")

			// Get the field's value from the form
			fileHeaders := form.File[fieldName]

			// Assuming there is only one file in the form, assign it to the field
			if len(fileHeaders) > 0 {
				fileHeader := reflect.ValueOf(fileHeaders[0])

				// Set the *multipart.FileHeader field of the struct
				structValue.FieldByName(field.Name).Set(fileHeader)
			}
		}
	}

	return v.ValidateStruct(data)
}

// ValidateStruct validates a given struct and returns a map of validation errors.
func (v *Validation) ValidateStruct(values interface{}) error {
	if err := v.Struct(values); err != nil {
		st := reflect.Indirect(reflect.ValueOf(values)).Type() // Get indirect type
		messages := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := st.FieldByName(err.Field())

			key := field.Tag.Get("json")

			if key == "" {
				key = strings.ToLower(err.Field())
			}

			messages[key] = getMessage(err)
		}

		return exception.Validation(messages)
	}

	return nil
}

// getMessage returns a validation error message based on the given validator.FieldError.
func getMessage(err validator.FieldError) string {
	attribute, format := attributes[err.Field()], messages[err.Tag()]
	if format == "" {
		format = messages["default"]
	}

	if attribute == "" {
		attribute = err.Field()
	}

	param := strings.Join(strings.Split(err.Param(), " "), ", ")

	t := template.Must(template.New("").Parse(format))
	b := new(strings.Builder)
	if err := t.Execute(b, map[string]interface{}{
		"Attribute": attribute,
		"Param":     param,
	}); err != nil {
		return err.Error()
	}

	return b.String()
}
