document.addEventListener('DOMContentLoaded', function() {
    const loginForm = document.getElementById('login-form');
    const usernameInput = document.getElementById('username');
    const passwordInput = document.getElementById('password');

    loginForm.addEventListener('submit', function(event) {
        event.preventDefault();
        const username = usernameInput.value;
        const password = passwordInput.value;

        // Example of sending login data to the server
        // Replace with your actual API endpoint
        fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        })
        .then(response => {
            if (response.ok) {
                return response.json();
            } else {
                throw new Error('Login failed');
            }
        })
        .then(data => {
            // Redirect to user or admin home based on role
            if (data.role === 'user') {
                window.location.href='user_list_product.html';
            } else if (data.role === 'admin') {
                window.location.assign( 'admin_home.html');
            } else {
                throw new Error('Invalid role');
            }
        })
        .catch(error => {
            console.error('Login error:', error);
            // Display error message to the user
        });
    });
});
