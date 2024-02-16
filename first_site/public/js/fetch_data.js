document.addEventListener('DOMContentLoaded', function() {
    fetch('http://localhost:8080/user')
        .then(response => response.json())
        .then(data => {
            const postsElement = document.getElementById('posts');
            console.log("DATA", data);
            data.forEach(post => {
                const postElement = document.createElement('div');
                postElement.innerHTML = `
                    <h2>${post.title}</h2>
                    <p>${post.body}</p>
                `;
                postsElement.appendChild(postElement);
            });
        })
        .catch(error => {
            console.error('Error fetching data:', error);
            document.getElementById('posts').innerHTML = '<p>Error loading posts.</p>';
        });
});
