<script lang="ts">
	import { Styles } from "sveltestrap";
	import TitleBar from "./components/TitleBar.svelte";
</script>

<Styles />

<body class="grid-container">
	<div style="grid-area: titlebar;">titlebar <TitleBar /></div>
	<div style="grid-area: console;">console</div>
	<div style="grid-area: mesh;">mesh</div>
	<div style="grid-area: geometry;">geometry</div>
	<div style="grid-area: solver;">solver</div>
	<div style="grid-area: display;">display</div>
	<div style="grid-area: plot;">plot</div>
	<div style="grid-area: parameters;">parameters</div>
</body>

<style>
	.grid-container {
		display: grid;
		grid-template-areas:
			"titlebar titlebar parameters console "
			"geometry mesh solver console"
			"display display plot plot"
			"display display plot plot"
			"display display plot plot";
		grid-gap: 10px;
		background-color: #2196f3;
		padding: 10px;
	}

	.grid-container > div {
		background-color: rgba(255, 255, 255, 0.8);
		text-align: center;
		padding: 20px 0;
		font-size: 30px;
	}
</style>
