const handleRegister = async() => {
	const checkBox = document.getElementById("terms");
	if (!checkBox.checked) {
		alert("Agree with terms and conditions");
		return;
	}

	const username = document.getElementById("username").value;
	const email = document.getElementById("email").value;
	const password = document.getElementById("password").value;
	const confirmedPassword = document.getElementById("confirmPassword").value;

	if (password !== confirmedPassword) {
		document.getElementById("passwordError").classList.remove("hidden");
		return;
	} else document.getElementById("passwordError").classList.add("hidden");

	const resp = await fetch("/api/register", {
		method: "POST",
		body: JSON.stringify({"user": username, "email": email, "pass": password})
	});
	if(resp.ok) {
		const res = await resp.json();
		if (res.status === 'ok') {
			alert("Register successfull. Proceed to login");
			window.location = "/login";
		}
	} else {
		alert("Something went wrong :(");
	}
};
