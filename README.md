# Craig Stobo Portfolio Website

A modern portfolio website built with Go, HTMX, Alpine.js, and Tailwind CSS, designed with the University of Otago color scheme.

## Technologies Used

- **Backend**: Go with Gorilla Mux router
- **Frontend**: HTMX for dynamic interactions, Alpine.js for client-side state management
- **Styling**: Tailwind CSS with custom University of Otago color palette
- **Build**: Static site generation capability

## Color Scheme

### Primary Colors
- **St Andrew's Blue**: `#060d87` - Primary brand color
- **Gold Rush**: `#f4b142` - Secondary brand color
- **Black**: `#000000`
- **White**: `#ffffff`

### Secondary Colors
- **Deep Blue**: `#000044`
- **Deep Gold**: `#ee8a07`
- **True Blue**: `#317bff`
- **True Yellow**: `#ffde46`
- **Sky**: `#daebfc`
- **Ivory**: `#fbf9ed`

## Getting Started

### Prerequisites
- Go 1.21 or later
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd craig-portfolio-website
```

2. Install dependencies:
```bash
make deps
# or
go mod tidy
```

### Development

Start the development server:
```bash
make dev
# or
go run main.go
```

Visit `http://localhost:8080` to view the website.

### Building Static Site

Generate a static site for deployment:
```bash
make build-static
# or
curl http://localhost:8080/build
```

The static site will be generated in the `./dist/` directory.

## Project Structure

```
├── main.go              # Main application server
├── templates/           # HTML templates
│   ├── index.html      # Main page template
│   └── partials/       # HTMX partial templates
├── static/             # Static assets
│   ├── images/         # Images and logos
│   ├── docs/          # PDF documents
│   ├── css/           # Custom CSS (if needed)
│   └── js/            # Custom JavaScript (if needed)
├── dist/              # Generated static site
├── Makefile          # Build commands
└── README.md         # This file
```

## Features

- **Responsive Design**: Mobile-first approach with Tailwind CSS
- **Smooth Animations**: CSS transitions and Alpine.js animations
- **Dynamic Content**: HTMX for seamless content switching
- **Static Site Generation**: Build deployable static files
- **University Branding**: Official University of Otago color scheme

## Sections

1. **Landing Page**: Hero section with animated background
2. **About**: Personal summary and qualifications with tab switching
3. **Services**: Key areas of expertise with alternating layout
4. **My Work**: Current and previous positions with company logos
5. **Articles & Press**: Publications and media coverage
6. **Contact**: Social media and contact links

## Deployment

The built static site in `./dist/` can be deployed to any static hosting service:
- Netlify
- Vercel
- GitHub Pages
- AWS S3
- Any web server

## Development Commands

```bash
make deps         # Install dependencies
make dev          # Start development server
make build-static # Generate static site
make clean        # Clean build artifacts
make fmt          # Format Go code
make lint         # Run linter
make test         # Run tests
make build        # Build binary
make all          # Full build process
```

## License

© 2024 Craig Stobo. All rights reserved.