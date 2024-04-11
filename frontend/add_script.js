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
        // Handle success, e.g., show a success message to the user
    })
    .catch(error => {
        console.error('Error adding product:', error);
        // Handle error, e.g., show an error message to the user
    });
});


// function addProduct(formData) {
//     fetch('http://localhost:8080/addproducts', {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'multipart/form-data' // Corrected content type
//         },
//         body: formData
//     })
//     .then(response => {
//         if (!response.ok) {
//             throw new Error('Network response was not ok');
//         }
//         return response.json();
//     })
//     .then(data => {
//         console.log('Product added successfully:', data);
//         // Handle success, e.g., show a success message to the user
//     })
//     .catch(error => {
//         console.error('Error adding product:', error);
//         // Handle error, e.g., show an error message to the user
//     });
// }

// // Example usage:
// const form = document.getElementById('add-product-form');
// form.addEventListener('submit', function(event) {
//     event.preventDefault(); // Prevent form submission
//     const formData = new FormData(form);
//     addProduct(formData); // Call addProduct function with form data
// });
