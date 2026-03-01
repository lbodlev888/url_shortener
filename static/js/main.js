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

window.onload = () => {
	const savedTheme = localStorage.getItem('theme') || 'system';
    setTheme(savedTheme);
};
