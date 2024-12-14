<script>
	import '../app.css';
	let { children } = $props();
	export const ssr = true
	export const prerender = true
</script>

{@render children()}
