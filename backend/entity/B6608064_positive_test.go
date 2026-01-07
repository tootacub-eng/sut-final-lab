package entity
import(
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestBook(t *testing.T){
	g := NewGomegaWithT(t)

	book := Book{
		Title: "anime",
		Price: 599,
		Code: "BK123456",
	}

	ok, err := govalidator.ValidateStruct(book)

	g.Expect(ok).To(BeTure())
	g.Expect(err).To(BeNil())
}