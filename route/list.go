package route

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/fumeapp/tonic/render"
	"github.com/gofiber/fiber/v2"
)

type RouteInfo struct {
	Methods  []string
	Name     string
	Path     string
	Handlers string
}

// show a list of all routes
func List(c *fiber.Ctx) error {

	routes := []RouteInfo{}

	for _, routeStack := range c.App().Stack() {
		for _, route := range routeStack {
			method := []string{route.Method}
			var handlers string
			for i, handler := range route.Handlers {
				handlers += runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
				if i != len(route.Handlers)-1 {
					handlers += ", "
				}
				handlers = strings.Replace(handlers, "github.com/", "", -1)
			}
			if strings.Contains(handlers, "Benchmark") {
				continue
			}
			existing := routeExists(routes, route.Name, handlers, method[0])
			if existing != nil {
				existing.Methods = append(existing.Methods, route.Method)
			} else {
				routes = append(routes, RouteInfo{
					Methods:  method,
					Name:     route.Name,
					Path:     route.Path,
					Handlers: handlers,
				})
			}
		}
	}

	output := `
<style lang="css">
	:root {
		--gray-400: #9ca3af;
		--gray-600: #4b5563;
		--gray-800: #1f2937;
		--gray-900: #111827;
		--gray-300: #d1d5db;
	}
	body { 
		margin: 20px; 
		background-color: var(--gray-900);
		color: var(--gray-400);
	}
	.routes {
		border: 1px solid var(--gray-800);
		padding: 6px;
		width: 100%;
		border-collapse: collapse;
		font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Helvetica,Arial,serif;
		font-size: 14px;
	}
	.routes th {
		color: var(--gray-300);
	}
	.routes th, .routes td {
		padding: 4px 6px;
		border-right: 1px solid var(--gray-800);
	}
	.routes tr { 
		border: 1px solid var(--gray-800); 
	}
</style>

<table class="routes">
	<tr>
		<th>Method</th>
		<th>Name</th>
		<th>Path</th>
		<th>Handlers</th>
	</tr>
	`

	for _, route := range routes {
		output += `
		<tr>
			<td>` + methodList(route.Methods) + `</td>
			<td>` + route.Name + `</td>
			<td>` + route.Path + `</td>
			<td>` + route.Handlers + `</td>
		</tr>`
	}

	output += `</table>`

	return render.HTML(c, output)
}

func routeExists(routes []RouteInfo, name string, handlers string, method string) *RouteInfo {
	for index, route := range routes {
		if route.Handlers == handlers && (route.Name == name || method == "HEAD") {
			return &routes[index]
		}
	}
	return nil
}

func methodList(methods []string) string {
	var output string
	for i, method := range methods {
		if method == "GET" {
			output += `<span style="color: blue;">GET</span>`
		}
		if method == "HEAD" {
			output += `<span style="color: gray;">HEAD</span>`
		}
		if method == "POST" {
			output += `<span style="color: orange;">POST</span>`
		}
		if method == "PUT" {
			output += `<span style="color: yellow;">PUT</span>`
		}
		if method == "DELETE" {
			output += `<span style="color: red;">DELETE</span>`
		}
		if i != len(methods)-1 {
			output += `<span style="color: #374151;">|</span>`
		}
	}
	return output
}
