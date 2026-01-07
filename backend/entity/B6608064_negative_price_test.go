package entity
import(
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestBookPrice(t *testing.T){
	g := NewGomegaWithT(t)

	book := Book{
		Title: "catoon",
		Price: "5",
		Code: "BK123457",
	}

	ok, err := govalidator.ValidateStruct(book)

	g.Expect(ok).To(BeFalse())
	g.Expect(err).NotTo(BeNil())
}