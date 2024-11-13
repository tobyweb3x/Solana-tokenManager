format:
	@templ fmt .

generate:
	@tailwindcss -i assets/styles/input.css  -o assets/styles/output.css --minify
	@pnpm build
	@templ generate && templ fmt .

dev:
	@templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

tailwind:
	@tailwindcss -i assets/styles/input.css  -o assets/styles/output.css --watch --minify

run:
	air

.PHONY: format generate dev run tailwind
