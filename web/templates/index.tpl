{{ define "index.tpl" }}
<!DOCTYPE html><html>
<head>
	<title>Gotta login</title>
	<link rel="stylesheet" href="css/pure-min.css">
	<link rel="stylesheet" href="css/index.css">
</head>

<body>
	<div class="container">
		{{ if .credError }}
		<p>You have provided invalid credentials.</p>
		{{ end }}

		<form class="pure-form" method="POST" action="/login">
			<fieldset class="pure-group">
				<input type="text" class="pure-input-1" name="user" placeholder="Username">
				<input type="password" class="pure-input-1" name="pass" placeholder="Password">
			</fieldset>

			<button type="submit" class="pure-button pure-input-1-2 pure-button-primary" style="width: 100%">Submit</button>
		</form>
	</div>
</body>
</html>
{{ end }}