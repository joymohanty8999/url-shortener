document.addEventListener("DOMContentLoaded", function () {
    fetch("https://snip-snip-go-2f69a42960b8.herokuapp.com/api/urls") // Makes a GET request to /api/urls to retrieve all shortened URLs
        .then(response => response.json())
        .then(data => {
            const urlList = document.getElementById("urlList");
            
            // Loops through the data array and displays each shortened URL as a clickable link

            data.forEach(url => {
                const shortUrl = `https://snip-snip-go-2f69a42960b8.herokuapp.com/api/${url.short_url}`;
                urlList.innerHTML += `<p><a href="${shortUrl}" target="_blank">${url.short_url}</a> → ${url.original_url}</p>`; // Displays the shortened URL and the original URL as mapping
            });
        })
        .catch(error => console.error("Error fetching URLs:", error));
});