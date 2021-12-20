<script lang="ts">
    import { onMount } from "svelte";
    import { getDir } from "../api/crud";

    // import { Icon, Col, Container, Row } from "sveltestrap";

    let dir = "Loading ... ";
    onMount(async () => {
        const res = await getDir();
        console.log;
        dir = res.json.res;
    });
</script>

<div>
    {dir}
</div>
