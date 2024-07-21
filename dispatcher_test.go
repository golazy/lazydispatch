package lazydispatch

import (
	"fmt"
	"net/http"
	"testing"
)

var langs = []string{"posts", "publicaciones", "veroffentlichung"}

type AuthorsController struct{}

func TestDispatcher(t *testing.T) {

	d := New()
	d.Draw(func(d *Scope) {

		for _, p := range langs {
			d.Resources(&PostsController{}).Path(p).Draw(func(posts *Scope) {
				posts.Resource(&AuthorsController{})
			})
		}

		d.Namespace("admin").Draw(func(admin *Scope) {

		})

		//d.Get("/patientes/:id").As("patient").To(&PostsController{}.Index)

	})

}

var postsController = &PostsController{}

func ExampleDispatcherUse(t *testing.T) {
	d := New()
	d.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL)
			next.ServeHTTP(w, r)
		})
	})

}

func ExampleDispatcherDraw(t *testing.T) {
	dispatcher := New()
	dispatcher.Draw(func(routes *Scope) {
		routes.Path("").Resource(sessionController).NewName()
		routes.Resources(postsController)
		routes.Namespace("admin").Draw(func(admin lazydispatch.Scope) {
			admin.Resources(usersController)
		})

		routes.Path("/backend/*").Use(otherHandler)
		routes.Path("/").Resources(pagesController)
	})

	fmt.Printf("%20s\t%20s\t%20s\n", "Path", "Target", "Name")
	for _, r := range dispatcher.Routes {
		fmt.Printf("%20s\t", r.URL, r.Target, r.Name)
	}

	// Output:
	// Path Target Name
	//
}
