<script>
    import ClipboardIcon from "./icons/ClipboardIcon.svelte";
    import LinkIcon from "./icons/LinkIcon.svelte";

    export let x = 0;
    export let y = 0;
    export let z = 0;
</script>

<div class="bar">
    <p><span class="box x">X</span><span class="text">{x}</span></p>
    <p><span class="box y">Y</span><span class="text">{y}</span></p>
    <p><span class="box z">Z</span><span class="text">{z}</span></p>

    <a href="/#" class="copy">
        <ClipboardIcon />
    </a>

    <a href="/#" class="copy">
        <LinkIcon />
    </a>
</div>

<style>
    .bar {
        display: flex;
        align-items: center;
        gap: 10px;
        color: black;
        padding: 0 15px 0 15px;
        height: 40px;
        background-color: #f0f0f0;
        border-radius: 15px;
        box-shadow: 0px 0px 15px 0px rgba(0, 0, 0, 0.5);
    }

    .text {
        display: inline-block;
        min-width: 40px;
        margin-left: 5px;
        font-size: 13pt;
        font-family: monospace;
    }

    .box {
        display: inline-block;
        font-weight: bold;
        height: 14pt;
        width: 14pt;
        border-radius: 5px;
        text-align: center;
        user-select: none;
    }

    .x {
        background-color: rgba(219, 138, 138, 0.5);
        color: rgb(120, 20, 10);
    }

    .y {
        background-color: rgba(131, 218, 153, 0.5);
        color: rgb(10, 100, 10);
    }

    .z {
        background-color: rgba(127, 159, 248, 0.5);
        color: rgb(10, 10, 120);
    }

    .copy {
        display: flex;
        color: black;
    }
</style>
