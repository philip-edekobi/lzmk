package codegen

import (
	"fmt"
)

func generateInitalBoilerPlate(title, author, date string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en" class="scroll-smooth">
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1"  />

	<title>%s</title>
	<meta name="author" content="%s" />
	<meta name="date" content="%s" />

	<!-- Tailwind CDN -->
	<script src="https://cdn.tailwindcss.com"></script>

	<script>
		tailwind.config = {
			darkMode: 'class',
			theme: {
				extend: {
					fontFamily: {
						mono: ['JetBrains Mono', 'monospace'],
					}
				}
			}
		}
	</script>

	<!-- Theme bootstrap -->
	<script>
		const t = localstorage.getItem("theme");
		if (t === "dark" || (!t && matchMedia("(prefers-color-scheme: dark)").matches)) {
			document.documentElement.classList.add("dark");
		}
	</script>
</head>

<body class="bg-white dark:bg-zinc-950 text-zinc-900 dark:text-zinc-100 antialiased">

	<!-- Header -->
	<header class="border-b border-zinc-200 dark:border-zinc-800">
		<div class="max-w-3xl mx-auto px-6 py-10">
			<h1 class="text-3xl font-bold tracking-tight">
				%s
			</h1>

			<p>
				<span class="font-medium">%s</span> .
				<time datetime="%s">%s</time>
			</p>
		</div>
	</header>

	<!-- Article -->
	<main class="max-w-3xl mx-auto px-6 py-16 space-y-16">
	`, title, author, date, title, author, convertDateToIsoTime(date), date)
}

func generateClosingBoilerPlate(author, date string) string {
	year := getYear(date)

	return fmt.Sprintf(`
	</main>

	<!-- Footer -->
	<footer class="border-t border-zinc-200 dark:border-zinc-800">
	<div class="max-w-3xl mx-auto px-6 py-8 text-sm text-zinc-500 dark:text-zinc-400">
			&copy; %d %s
		</div>
	</footer>

	<!-- Theme toggle -->
	<script>
		document.addEventListener("keydown", e => {
			if (e.key === "t") {
				const root = document.documentElement;
				const dark = root.classList.toggle("dark");
				localStorage.setItem("theme", dark ? "dark" : "light");
			}
		})
	</script>

</body>
</html>
	`, year, author)
}
