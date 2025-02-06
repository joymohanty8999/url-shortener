document.addEventListener("DOMContentLoaded", function () {
    fetch("https://snip-snip-go-2f69a42960b8.herokuapp.com/api/urls")
        .then(response => response.json())
        .then(data => {
            const urlList = document.getElementById("urlList");
            urlList.innerHTML = "<h3>Your Shortened URLs:</h3>";

            data.forEach(url => {
                urlList.innerHTML += `<p><a href="${url.short_url}" target="_blank">${url.short_url}</a> → ${url.original_url}</p>`;
            });
        })
        .catch(error => console.error("Error fetching URLs:", error));
});