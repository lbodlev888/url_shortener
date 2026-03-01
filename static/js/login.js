const handleLogin = async() => {
	const login = document.getElementById("username").value;
	const password = document.getElementById("password").value;

	const resp = await fetch("/api/login", {
		method: "POST",
		body: JSON.stringify({"user": login, "pass": password})
	});
	if (resp.ok) {
		const res = await resp.json();
		document.cookie = `token=${res.token}`;
		window.location = "/";
	} else if (resp.status === 403) {
		alert("Invalid credentials");
	} else {
		alert("Something failed :(");
	}
};
