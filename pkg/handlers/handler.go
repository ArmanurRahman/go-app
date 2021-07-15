package handlers

import (
	"net/http"

	"github.com/ArmanurRahman/go-app/pkg/config"
	"github.com/ArmanurRahman/go-app/pkg/models"
	"github.com/ArmanurRahman/go-app/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
	//_, _ = fmt.Fprintf(w, "This is Home page")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strngMap := make(map[string]string)

	strngMap["test"] = "Hello world"
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: strngMap})
	//sum := addNumber(2, 2)
	//_, _ = fmt.Fprintf(w, "This is about page and 2 + 2 is %d", sum)
}

/*
func Devide(w http.ResponseWriter, r *http.Request) {
	rs, err := devideValue(100.00, 10.5)
	if err != nil {
		fmt.Fprintf(w, "can not devided by zero")
		return
	}
	fmt.Fprintf(w, "%f devided by %f is %f ", 100.00, 10.5, rs)
}

func devideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("can not devided by zero")
		return 0, err
	}
	return x / y, nil
}
func addNumber(x, y int) int {
	return x + y
}
*/
