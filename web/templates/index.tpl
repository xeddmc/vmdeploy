{{ define "index.tpl" }}
<!DOCTYPE html><html>
<head>
	{{ if .logged }}
	<title>Cute</title>
	{{ else }}
	<title>Gotta login</title>
	{{ end }}
</head>

<body>
	{{ if .logged }}
	<p>WELCOME. YOU MAY <a href="/logout">LOGOUT</a>.</p>
	{{ else }}
	<p>This is the right way.</p>
	<form method="POST" action="/login">
		<label for="user">Username</label>
		<input type="text" name="user">

		<label for="pass">Password</label>
		<input type="password" name="pass">

		<input type="submit" value="Submit">
	</form>
	{{ end }}
</body>
</html>
{{ end }}