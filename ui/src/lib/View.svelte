<script>
    import Cube from "./icons/CubeIcon.svelte";
    import Window from "./icons/WindowIcon.svelte";

</script>
<div class="container">
    <div class="button active">2D</div>
    <div class="button">3D</div>
</div>

<style>
    .container {
        display: flex;
        flex-direction: column;
        align-items: center;
        height: 80px;
        width: 40px;
        background-color: #f0f0f0;
        border-radius: 15px;
        box-shadow: 0px 0px 15px 0px rgba(0, 0, 0, 0.5);
    }

    .button {
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 12pt;
        flex-grow: 1;
        user-select: none;
        cursor: pointer;
        width: 40px;
        font-weight: bold;
        color: var(--fg-inactive);
    }

    .button:hover {
        color: var(--fg);
    }

    .active {
        color: var(--fg-active);
    }
</style>
