format:
	@templ fmt .

generate:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css --minify
	@pnpm build
	@templ generate && templ fmt .

templ:
	@templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

tailwind:
	@tailwindcss -i public/assets/styles/input.css  -o public/assets/styles/output.css --watch

.PHONY: format generate dev run tailwind

