document.getElementById('edit-product-form').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent form submission

    const formData = new FormData(this);
    const productId = getProductIdFromURL(); // You need to implement this function

    fetch(`http://localhost:8080/update_product/${Id}`, {
        method: 'POST',
        body: formData
    })
    .then(response => {
        if (response.ok) {
            alert("Product updated successfully");
            // Redirect to product details page or any other page
            window.location.href = `product_details.html?id=${Id}`;
        } else {
            throw new Error('Failed to update product');
        }
    })
    .catch(error => console.error('Error updating product:', error));
});
