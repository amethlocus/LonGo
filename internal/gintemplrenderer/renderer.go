package gintemplrenderer

// Importa los paquetes necesarios
import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin/render"

	"github.com/a-h/templ"
)

// Crea una variable global 'Default' que contiene una instancia vacía del renderizador HTMLTemplRenderer
var Default = &HTMLTemplRenderer{}

// Estructura principal del renderizador HTML para Templ
type HTMLTemplRenderer struct {
	// Si no se puede usar Templ, se usará este renderizador HTML como respaldo
	FallbackHtmlRenderer render.HTMLRender
}

// Método Instance: crea una instancia del renderizador
func (r *HTMLTemplRenderer) Instance(s string, d any) render.Render {

	// Intenta convertir los datos a un componente Templ y revisa si es o no un templ
	templData, ok := d.(templ.Component)
	if !ok {
		if r.FallbackHtmlRenderer != nil {
			// Usar el renderizador HTML normal de Gin
			return r.FallbackHtmlRenderer.Instance(s, d)
		}
	}

	// Si es un componente Templ, crea y retorna un Renderer configurado
	return &Renderer{
		Ctx:       context.Background(), // Contexto vacío por defecto
		Status:    -1,                   // -1 significa que no se establece un código HTTP por defecto
		Component: templData,            // Componente Templ a renderizar
	}
}

// Función para crear un Renderer directamente con contexto, estado y componente
func New(ctx context.Context, status int, component templ.Component) *Renderer {
	return &Renderer{
		Ctx:       ctx,       // Contexto pasado como argumento
		Status:    status,    // Código de estado HTTP a devolver
		Component: component, // Componente Templ que se renderizará
	}
}

// Estructura Renderer: define cómo renderizar un componente Templ
type Renderer struct {
	Ctx       context.Context // Contexto de ejecución
	Status    int             // Código de estado HTTP (por ejemplo 200, 404)
	Component templ.Component // Componente Templ a renderizar
}

// Método Render: se encarga de escribir la respuesta HTTP
func (t Renderer) Render(w http.ResponseWriter) error {
	t.WriteContentType(w) // Escribe el tipo de contenido "text/html"
	if t.Status != -1 {   // Si el código de estado no es -1...
		w.WriteHeader(t.Status) // Escribe el código HTTP en la respuesta
	}
	if t.Component != nil { // Si hay un componente Templ...
		// Renderiza el componente en la respuesta usando el contexto
		return t.Component.Render(t.Ctx, w)
	}
	return nil // Si no hay nada que renderizar, no devuelve error
}

// Método WriteContentType: define el encabezado "Content-Type" como HTML UTF-8
func (t Renderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
