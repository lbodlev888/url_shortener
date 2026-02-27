const generateShortUrl = () => {
}

const copyToClipboard = () => {
    const shortUrlText = document.getElementById('shortUrl').textContent;
    navigator.clipboard.writeText(shortUrlText).then(() => {
    const notification = document.getElementById('copyNotification');
    notification.classList.remove('hidden');

    notification.classList.remove('animate-fade-in-out');
    void notification.offsetWidth;
    notification.classList.add('animate-fade-in-out');

    setTimeout(() => {
			notification.classList.add('hidden');
		}, 2000);
    });
};

function toggleDropdown() {
    document.getElementById('themeDropdown').classList.toggle('hidden');
}

const setTheme = (mode) => {
    const html = document.documentElement;
    localStorage.setItem('theme', mode);
    if (mode === 'dark') {
    html.classList.add('dark');
    document.getElementById('theme-icon').textContent = '🌙';
    } else if (mode === 'light') {
    html.classList.remove('dark');
    document.getElementById('theme-icon').textContent = '☀️';
    } else {
    if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
        html.classList.add('dark');
        document.getElementById('theme-icon').textContent = '🌙';
    } else {
        html.classList.remove('dark');
        document.getElementById('theme-icon').textContent = '☀️';
    }
    }
    document.getElementById('themeDropdown').classList.add('hidden');
};

const handleLogin = async() => {
	const login = document.getElementById("email").value;
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
		console.log(res);
	} else {
		alert("Something went wrong :(");
	}
};

// Apply saved or system theme
(() => {
	const savedTheme = localStorage.getItem('theme') || 'system';
    setTheme(savedTheme);
})();
