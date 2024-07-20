document.getElementById('shorten-form').addEventListener('submit', async function(e) {
    e.preventDefault();
    const url = document.getElementById('url-input').value;
    const response = await fetch('/shorten', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url })
    });
    const data = await response.json();
    document.getElementById('result').innerText = `Shortened URL: ${data.shortenedUrl}`;
});