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
            shortUrl.textContent = `${window.location.origin}/g/${shortUrlText}`;
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

    // Re-trigger animation
    notification.classList.remove('animate-fade-in-out');
    void notification.offsetWidth; // Trigger reflow
    notification.classList.add('animate-fade-in-out');

    // Hide after animation
    setTimeout(() => {
        notification.classList.add('hidden');
    }, 2000);
    });
}