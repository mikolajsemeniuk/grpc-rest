package cart

import _ "embed"

//go:embed templates/cart.swagger.yaml
var swagger string

//go:embed templates/elements.tpl
var elements string
