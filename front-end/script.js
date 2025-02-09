document.getElementById('shorten-form').addEventListener('submit', async function(e) {
    e.preventDefault();
    const url = document.getElementById('url-input').value;

    try {
        const response = await fetch('/api/shorten', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ url: url })
        });

        if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
        }

        const data = await response.json();

        // Insert the shortened URL inside the result container with a clickable link
        document.getElementById('result').innerHTML = `
            <strong>Shortened URL:</strong> 
            <a href="${data.short_url}" target="_blank">${data.short_url}</a>
        `;

    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        document.getElementById('result').innerText = `Error: ${error.message}`;
    }
});

document.getElementById("viewUrls").addEventListener("click", function () {
    window.location.href = "urls.html";
});