<script lang="ts">
    import MainPage from "./lib/pages/MainPage.svelte";
    import { CurrentPage } from "./lib/pages/types";
    
    let currentPage: CurrentPage = $state(CurrentPage.Landing)
</script>

<!-- Header -->
<header>
    <div class="header-title">
      <span id="title">Ranking Estadual</span>
    </div>
</header>

<!-- main selector -->
{#if (currentPage == CurrentPage.Landing)}
    <MainPage />
{:else}
    <div></div>
{/if}

<footer>
    <div class="footer-description">
      <span id="footer">Daniel Guimarães @ 2025</span>
    </div>
</footer>


<style>
    header {
      min-width: 100vw;
      height: 6em;

      display: flex;
      flex-flow: column;
      justify-content: center;
      align-items: center;
    }

    #title {
      background-color: #0F828C;
      font-size: 2em;
      font-weight: 700;
      padding: .5em .9em;
      border-radius: 1em;
    }
</style>