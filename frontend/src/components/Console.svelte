<script>
</script>

<main>
    <div class="yaxis">yaxis</div>
    <div class="c1">
        <div class="plot">plot</div>
        <div class="xaxis">xaxis</div>
    </div>
</main>

<style>
    main {
        width: 100%;
        height: 100%;
        border: 2px solid blue;
        display: flex;
    }
    .yaxis {
        width: 20%;
        height: 100%;
        border: 2px solid red;
    }
    .c1 {
        width: 80%;
        height: 100%;
        border: 2px solid purple;
        display: flex;
        flex-direction: column;
    }
    .plot {
        width: 100%;
        height: 80%;
        border: 2px solid green;
    }
    .xaxis {
        width: 100%;
        height: 20%;
        border: 2px solid yellow;
    }
</style>
