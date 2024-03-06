<script>
  export let products;
</script>

{#if products.length > 0}
  <ul>
    {#each products as product}
      <li>
        <h2>{product.name}</h2>
        <p>{product.description}</p>
        <p>Price: ${product.price}</p>
        <p>Quantity: {product.quantity}</p>
      </li>
    {/each}
  </ul>
{:else}
  <p>No products available</p>
{/if}
