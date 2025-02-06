document.addEventListener("DOMContentLoaded", function () {
    fetch("https://snip-snip-go-2f69a42960b8.herokuapp.com/api/urls")
        .then(response => response.json())
        .then(data => {
            const urlList = document.getElementById("urlList");

            data.forEach(url => {
                const shortUrl = `https://snip-snip-go-2f69a42960b8.herokuapp.com/api/${url.short_url}`;
                urlList.innerHTML += `<p><a href="${url.shortUrl}" target="_blank">${url.shortUrl}</a> → ${url.original_url}</p>`;
            });
        })
        .catch(error => console.error("Error fetching URLs:", error));
});