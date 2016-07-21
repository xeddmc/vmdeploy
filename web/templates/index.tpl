{{ define "index.tpl" }}
<!DOCTYPE html><html>
<head>
	<title>Gotta login</title>
</head>

<body>
	<p>This is the right way.</p>

	{{ if .credError }}
	<p>You have provided invalid credentials.</p>
	{{ end }}

	<form method="POST" action="/login">
		<label for="user">Username</label>
		<input type="text" name="user">

		<label for="pass">Password</label>
		<input type="password" name="pass">

		<input type="submit" value="Submit">
	</form>
</body>
</html>
{{ end }}