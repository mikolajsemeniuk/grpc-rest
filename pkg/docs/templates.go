package docs

import _ "embed"

//go:embed templates/cart.swagger.yaml
var docs string

//go:embed templates/elements.tpl
var elements string
