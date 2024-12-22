<script>
    import Coordinates from "$lib/Coordinates.svelte";
    import Logo from "$lib/Logo.svelte";
    import Map from "$lib/Map.svelte";
    import Search from "$lib/Search.svelte";
    import View from "$lib/View.svelte";
    import Zoom from "$lib/Zoom.svelte";
</script>

<div class="map-container">
    <div class="map">
        <Map />
    </div>
    <div class="ui-overlay">
        <div class="left-bar">
            <Logo />
            <Zoom />
            <View />
        </div>
        <div class="top-bar">
            <Search />
            <Coordinates />
        </div>
    </div>
</div>

<style>
    .map {
        position: absolute;
        height: 100%;
        width: 100%;
        background-color: #008080;
    }

    .ui-overlay {
        display: flex;
        position: absolute;
        z-index: 10;
        height: 100%;
        width: 100%;
        pointer-events: none;
    }

    .top-bar {
        display: flex;
        margin-top: 10px;
        gap: 10px;
        height: 40px;
    }

    .left-bar {
        display: flex;
        flex-direction: column;
        margin: 10px;
        gap: 10px;
    }

    .map-container {
        margin: 0;
        height: 100vh;
        width: 100vw;
        background-color: black;
    }
</style>
