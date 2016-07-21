{{ define "dashboard.tpl" }}
<!DOCTYPE html><html>
<head>
	<title>Gotta login</title>
</head>

<body>
	<p>WELCOME [{{ .username }}]. YOU MAY <a href="/logout">LOGOUT</a>.</p>
</body>
</html>
{{ end }}