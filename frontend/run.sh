#!/bin/bash
make dev &
TEMPL_PID=$!

make tailwind &
TAILWIND_PID=$!

pnpm dev &
PNPM_PID=$!

trap "kill $TEMPL_PID $TAILWIND_PID $PNPM_PID 2>/dev/null" INT TERM
wait $TEMPL_PID $TAILWIND_PID $PNPM_PID
