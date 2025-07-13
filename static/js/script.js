function generateShortUrl() {
    const urlInput = document.getElementById('urlInput');
    const shortUrl = document.getElementById('shortUrl');
    const resultContainer = document.getElementById('resultContainer');
    
    fetch('/newlink', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({ url: urlInput.value })
    }).then(response => response.json())
    .then(data => {
        if (data.status && data.status == 'success') {
            const shortUrlText = data.shortUrl;
            shortUrl.textContent = `${window.location.origin}/${shortUrlText}`;
            resultContainer.classList.remove('hidden');
        } else {
            alert('Error generating short URL. Message: ' + data.message);
        }
    })
    .catch(error => {
        console.error('Error:', error);
        alert('An error occurred while generating the short URL. Please try again.');
    });
}

function copyToClipboard() {
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
}

function toggleDropdown() {
    document.getElementById('themeDropdown').classList.toggle('hidden');
}

function setTheme(mode) {
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
}

// Apply saved or system theme
(function () {
    const savedTheme = localStorage.getItem('theme') || 'system';
    setTheme(savedTheme);
})();