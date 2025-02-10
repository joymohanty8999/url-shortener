document.getElementById('shorten-form').addEventListener('submit', async function(e) {
    e.preventDefault();
    const url = document.getElementById('url-input').value; // Listens for the submit event and gets the URL value

    try {
        const response = await fetch('/api/shorten', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ url: url })
        }); // Sends a POST request /api/shorten with the user's URL as JSON

        if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
        }

        // Updates the page with the shortened URL as a clickable link

        const data = await response.json();

        document.getElementById('result').innerHTML = `
            <strong>Shortened URL:</strong> 
            <a href="${data.short_url}" target="_blank">${data.short_url}</a> 
        `;

    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        document.getElementById('result').innerText = `Error: ${error.message}`;
    }
});

// Redirects the user to the URLs page when the "View URLs" button is clicked

document.getElementById("viewUrls").addEventListener("click", function () {
    window.location.href = "urls.html";
});