<script>
	import '../app.css';
	let { children } = $props();
	export const csr = true
	export const ssr = true
	export const prerender = false
</script>

{@render children()}
