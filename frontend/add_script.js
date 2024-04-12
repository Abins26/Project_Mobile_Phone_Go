document.getElementById('add-product-form').addEventListener('submit', function(event) {
    event.preventDefault();

    const formData = new FormData(this);

    fetch('http://localhost:8080/addproducts', {
        method: 'POST',
        body: formData
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.text();
    })
    .then(data => {
        console.log('Product added successfully:', data);
        window.alert("New Product added Successfully ")
        // Handle success,  show a success message to the user
    })
    .catch(error => {
        console.error('Error adding product:', error);
        // Handle error, show an error message to the user
    });
});


