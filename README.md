# static-site-cms

This is a cms for static sites, allowing for easy editing of your site's content without the need to edit the source code.

### Development

It is built with [Svelte](https://svelte.dev/), [Tailwind](https://tailwindcss.com/) on the frontend and [Go](https://golang.org/) and [SQLite](https://www.sqlite.org/index.html) on the backend.

There are development Dockerfiles for both the frontend and backend, which can be used to run the application in a container.
These mount the source code into the container, so changes made to the source code will be reflected live in the running application.

To run the application in a container, run the following commands:

```bash
docker compose up dev --build (-d)
```

This will start the svelte dev server on port 5173, the go server on port 8080 and the static site will be served on port 3000.

The static site is served from the `_local_dev` directory, which is mounted as a volume into the container.

### Roadmap

- Allow uploads of static content which will be copied to the static site directory
- Allow this static content to be updated in the UI including text and images
- Add version control to the static site directory using git
- Add backup and restore functionality
- Build a CLI tool to generate the necessary files for the static site
- Add monitoring (using prometheus)
- Add user management with roles/permissions
- Add a way to add custom pages to the UI
- Allow for pre-defined footer and header content in one file and shown on all pages
- Add a way to add custom css using the UI (maybe using tailwind style classes for flexibility)
- Provide SEO control for each page in the UI

### Contributing

Contributions are welcome, please open an issue or PR.

#### License

GNU General Public License Version 3
