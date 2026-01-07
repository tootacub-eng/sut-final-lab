package entity
import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)
type Book struct{
	gorm.Model
	Title string  `valid:"stringlength(3|100)"`
	Price float64 `valid:"range(50|5000)"`
	Code string `valid:"matched(^B[K][0-9]{6}$)"`
}
func init(){
	govalidator.SetFieldsRequiredByDefault(false)
}