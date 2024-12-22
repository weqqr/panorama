<div class="container">
    <div class="button">&plus;</div>
    <div class="button">&minus;</div>
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
        font-size: 16pt;
        flex-grow: 1;
        user-select: none;
        cursor: pointer;
        width: 40px;
        font-weight: bold;
    }
</style>
