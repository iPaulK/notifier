package notifier

import (
	"text/template"
)

var (
	mailTemplate = template.Must(template.New("mail").Parse(`
<html>
<head>
<title></title>
<style type="text/css">
body {
background: #fff; margin-top: 30px; padding: 0;
}
</style>
</head>
<body style="background: #fff; margin-top: 30px; padding: 0;">
{{ .Body }}
</body>
</html>
`))
)
