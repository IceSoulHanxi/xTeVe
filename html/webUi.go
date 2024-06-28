package html

import "embed"

//go:embed *.html css img js lang video
var WebUI embed.FS
