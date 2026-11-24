# commonprefix.com

Vercel serves `public/` exactly as it is in this repository: there is no build step on
Vercel (`vercel.json`). Every page in `public/` is built elsewhere and committed here.

## Where the pages come from

The site is built from **[common-prefix-coder](https://github.com/giocherx/common-prefix-coder)**
(Vite + TypeScript): the home page, Coder (`/coder`), Team, Research, Blog and every blog post,
Podcast, Careers, the partner and grant pages and the 404.

To publish a change:

1. Make the change in `common-prefix-coder`, then build it:

   ```sh
   npm install
   npm run build        # writes dist/
   ```

2. Copy the build into this repository's `public/`. File names in `assets/` change with every
   build, so remove the old ones first:

   ```sh
   rm public/assets/*.js public/assets/*.css public/assets/*.woff2
   rsync -a ../common-prefix-coder/dist/ public/
   ```

3. Check what changed (`git status`), commit, and open a pull request. Vercel posts a preview
   link on the pull request: check it on a phone and an iPad before merging.

Build from the commit you mean to publish: a build carries every change in that branch.

### Blog posts

Posts are pages in `common-prefix-coder/blog/<slug>.html`. A new post can start as a copy of
an existing one. `scripts/gen-posts.py` there converted the previous site's posts once; given
slugs, it converts only those (`python3 -I scripts/gen-posts.py <old public/blog> . <slug>`).

## Kept from the previous site

- `public/static/`: files linked from outside (PDFs, images, fonts, Prism for code in posts).
  For team pictures, add a smaller version with the suffix `_w150` (`sips -Z 150 name_w150.jpg`).
- `public/clients/deviantart.html`, `enosys.html`, `pantos.html`: partner pages of the previous
  site, no longer linked from the new one.
- `public/blog/The Landscape….html` and `The Past….html`: old addresses of two posts, which
  now redirect to them.
- `public/.well-known/xrp-ledger.toml` and the redirects in `vercel.json`.

## The Go generator (previous site)

`main.go` and `templates/` built the previous site. **Do not run it to publish**: it writes into
`public/` and would overwrite the current pages. On its own it now does nothing and says why.

```sh
go run . --serve          # serve public/ on http://localhost:8100 (builds nothing)
go run . --legacy-pages   # rebuild the previous site's pages (overwrites public/)
```
