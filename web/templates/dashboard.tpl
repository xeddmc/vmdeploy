{{ define "dashboard.tpl" }}
<!DOCTYPE html><html>
<head>
	<title>Gotta login</title>
	<link rel="stylesheet" href="css/pure-min.css">
</head>

<body>
	<p>WELCOME [{{ .username }}]. YOU MAY <a href="/logout">LOGOUT</a>.</p>
</body>
</html>
{{ end }}