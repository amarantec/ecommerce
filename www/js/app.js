document.addEventListener("DOMContentLoaded", function () {
    fetch('/api/products')
        .then(response => response.json())
        .then(products => {
            const productList = document.getElementById('product-list');
            products.forEach(product => {
                const li = document.createElement('li');
                li.className = 'product-item';
                li.textContent = `${product.title} - ${product.price}`;
                productList.appendChild(li);
            });
        })
        .catch(error => console.error('Error fetching products:', error));
});
