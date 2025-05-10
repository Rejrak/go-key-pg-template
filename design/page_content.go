package design

import (
	"be/design/errors"

	. "goa.design/goa/v3/dsl"
)

var PageContent = Type("PageContent", func() {
	Description("Contenuto libero JSON di una pagina")
	Attribute("id", Int, "ID univoco della pagina", func() {
		Example(1)
	})
	Attribute("data", Any, "Dati JSON liberi della pagina", func() {
		Example(map[string]interface{}{
			"title": "Welcome",
			"body":  "This is a sample page",
		})
	})
	Required("id", "data")
})

var CreatePagePayload = Type("CreatePagePayload", func() {
	Description("Payload per creare un nuovo contenuto di pagina")
	Attribute("data", Any, "Contenuto JSON libero", func() {
		Example(map[string]interface{}{"title": "New Page"})
	})
	Required("data")
})

var UpdatePagePayload = Type("UpdatePagePayload", func() {
	Description("Payload per aggiornare un contenuto di pagina")
	Attribute("id", Int, "ID della pagina da aggiornare")
	Attribute("data", Any, "Contenuto JSON aggiornato")
	Required("id", "data")
})

var PageContentService = Service("page_content", func() {
	Security(OAuth2, func() {
		Scope("openid")
	})
	Description("Gestione dei contenuti dinamici delle pagine in formato JSON")

	HTTP(func() {
		Path("/page-content")
	})

	Error("notFound", errors.NotFound)
	Error("internalServerError", errors.InternalServerError)
	Error("badRequest", errors.BadRequest)

	Method("create", func() {
		Description("Crea un nuovo contenuto di pagina")
		Payload(CreatePagePayload)
		Result(PageContent)

		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("get", func() {
		Description("Recupera un contenuto di pagina per ID")
		Payload(func() {
			Attribute("id", Int, "ID del contenuto da recuperare")
			Required("id")
		})
		Result(PageContent)

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			errors.CommonResponses()
		})
	})

	Method("update", func() {
		Description("Aggiorna un contenuto esistente")
		Payload(UpdatePagePayload)
		Result(PageContent)

		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
			errors.CommonResponses()
		})
	})

	Method("delete", func() {
		Description("Elimina un contenuto di pagina")
		Payload(func() {
			Attribute("id", Int, "ID del contenuto da eliminare")
			Required("id")
		})
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
			errors.CommonResponses()
		})
	})

	Method("list", func() {
		Description("Restituisce la lista dei contenuti di pagina")
		Payload(func() {
			Attribute("limit", Int, "Limite massimo di risultati", func() {
				Default(10)
				Minimum(1)
				Maximum(100)
			})
			Attribute("offset", Int, "Offset per la paginazione", func() {
				Default(0)
				Minimum(0)
			})
		})
		Result(ArrayOf(PageContent))
		HTTP(func() {
			GET("/")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})
})
