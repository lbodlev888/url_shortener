let qrcode = null;

const openModal = (modalId) => { document.getElementById(modalId).classList.remove('hidden'); }
const closeModal = (modalId) => { document.getElementById(modalId).classList.add('hidden'); }

const copyUrl = (url, btn) => {
	navigator.clipboard.writeText(window.location.href + url);
	const toast = document.getElementById('copyToast');
	const rect = btn.getBoundingClientRect();
	toast.classList.remove('hidden');
	toast.style.top  = (rect.top + window.scrollY - toast.offsetHeight - 10) + 'px';
	toast.style.left = (rect.left + window.scrollX + rect.width / 2 - toast.offsetWidth / 2) + 'px';
	// recalc after paint since offsetWidth is 0 while hidden
	requestAnimationFrame(() => {
		toast.style.top  = (rect.top + window.scrollY - toast.offsetHeight - 10) + 'px';
		toast.style.left = (rect.left + window.scrollX + rect.width / 2 - toast.offsetWidth / 2) + 'px';
	});
	setTimeout(() => toast.classList.add('hidden'), 2200);
};

const deleteShort = async(path) => {
	if(!confirm("Are you sure you want to delete the following short url?")) return;
	const resp = await fetch(`/api/short/${path}`, {
		method: "DELETE",
	});
	if(resp.ok) {
		const res = await resp.json();
		if (res.status === 'ok')
			document.getElementById(path).remove();
		
	} else alert('Something went wrong :(');
};

const displayQrCode = (path) => {
	if (qrcode === null) {
		qrcode = new QRCode("qrcode");
	}
	qrcode.clear();
	qrcode.makeCode(window.location.href + path);
	openModal('qrModal');
};

const submitNewUrl = async() => {
	const newUrl = document.getElementById('newUrl').value;

	const resp = await fetch('/api/short', {
		method: "POST",
		body: JSON.stringify({"url": newUrl})
	});
	if (!resp.ok)
		alert('Something failed :(');
	closeModal('addUrlModal');
	location.reload();
}

const copyImage = () => {
	const canvas = document.querySelector(`#qrcode canvas`);

	canvas.toBlob(async (blob) => {
		try {
			await navigator.clipboard.write([
				new ClipboardItem({ 'image/png': blob })
			]);
			document.getElementById('copied_status').classList.remove('hidden');
			setTimeout(() => {
				document.getElementById('copied_status').classList.add('hidden');
			}, 4000);
		} catch (err) {
			console.error('Failed to copy QR:', err);
		}
	});
};

const logout = () => {
	document.cookie = "token=;Expires=Thu, 01 Jan 1970 00:00:01 GMT;"
	window.location = "/login";
};
