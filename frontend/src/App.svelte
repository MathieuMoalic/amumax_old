<script lang="ts">
	import { Styles } from "sveltestrap";
	import TitleBar from "./components/TitleBar.svelte";
	import Plot from "./components/Plot.svelte";
	import Display from "./components/Display.svelte";
</script>

<Styles />

<body class="grid-container">
	<!-- <div style="grid-area: titlebar;">titlebar <TitleBar /></div> -->
	<div style="grid-area: console;">console</div>
	<div style="grid-area: mesh;">mesh</div>
	<div style="grid-area: geometry;">geometry</div>
	<div style="grid-area: solver;">solver</div>
	<div style="grid-area: display;"><Display /></div>
	<!-- <div style="grid-area: plot;"><Plot /></div> -->
	<div style="grid-area: parameters;">parameters</div>
</body>

<style>
	:root {
		--bg: hsl(200, 85%, 13%);
		--edge: hsl(207, 76%, 43%);
		--btn: hsl(194, 82%, 28%);
		--hover: hsla(194, 82%, 28%, 0.349);
		font-size: 24px;
	}
	:global(body) {
		margin: 0px;
		border: 0px;
		padding: 0px;
	}
	.grid-container {
		display: grid;
		grid-template-areas:
			"titlebar titlebar parameters console console"
			"geometry mesh solver console console"
			"display display display plot plot"
			"display display display plot plot"
			"display display display plot plot";
		grid-gap: 3px;
		background-color: var(--edge);
		padding: 12px;
	}

	.grid-container > div {
		text-align: center;
		color: white;
		padding: 20px 0;
		font-size: 30px;
		background-color: var(--bg);
	}
</style>
