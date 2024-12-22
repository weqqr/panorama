<script lang="ts">
    import Cube from "./icons/CubeIcon.svelte";
    import SearchIcon from "./icons/SearchIcon.svelte";

    function hasContent(): boolean {
        return true;
    }

    const text = "wo";

    const suggestions = [
        {
            text: "Hello, world",
            pos: [1241, 8114],
        },
        {
            text: "mega sword of woo",
            pos: [490, -8313],
        }
    ];
</script>

<div class="container">
    <div class="bar">
        <input class="text" placeholder="Search Places or Coordinates..." />
        <a class="search" href="/#"><SearchIcon /></a>
    </div>
    <div>
        {#each suggestions as suggestion}

        {/each}
    </div>
</div>

<style>
    .bar {
        display: flex;
        align-items: center;
        padding: 0 15px 0 15px;
        width: 300px;
        height: 40px;
        background-color: #f0f0f0;
        border-radius: 15px;
        box-shadow: 0px 0px 15px 0px rgba(0, 0, 0, 0.5);
        transition: background-color 0.15s ease;
    }

    .bar:has(input:focus) {
        background-color: white;
    }

    .text {
        font-size: 11pt;
        height: 100%;
        width: calc(100% - 30px);
        border-style: none;
        background-color: transparent;
        outline: none;
    }

    .search {
        display: flex;
        color: var(--fg-inactive);
    }

    .search:hover {
        color: var(--fg);
    }
</style>
