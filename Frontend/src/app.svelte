<script>
  import { onMount } from "svelte";
  import ProductList from "./ProductList.svelte";

  let products = [];

  onMount(async () => {
    const response = await fetch("http://localhost:8080/api/products");
    products = await response.json();
  });
</script>

<h1>Products</h1>

<ProductList {products} />
