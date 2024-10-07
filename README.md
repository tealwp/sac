# Science As Code

This project aims to bridge the gap between science and code. It provides a platform for people with a scientific background to explain complex concepts using basic programming models, functions, and methods.

## Current Status

The project is in its early stages. Currently, it includes:

1. A home page
2. A few test content files (check `/static/content/...`)

## Project Structure

The `static/content` directory is organized to reflect the URL structure for accessing scientific topics:

```
static/content/
└── science/
    ├── [category1]/
    │   ├── [subject1].go
    │   └── [subject2].go
    ├── [category2]/
    │   ├── [subject3].go
    │   └── [subject4].go
    └── chemistry/
        └── chemical_reactions/
            └── reactions.go
```

Worth noting that these can be any depth of nested content.

This structure corresponds to URL paths like:
`http://localhost:4000/science/[category]/[subject]`

For example, the test subject page is accessed at:
`http://localhost:4000/science/testsubject/testsubject`

### Running the Server

To run the server, follow these steps:

1. Install Go
2. Open a terminal
3. Navigate to the project directory
4. Run the following command:

```bash
bash ./bin/run.sh 
```

The server will start and you can access it at `http://localhost:4000`.

### Navigating the Pages

Once the server is running, you can access the following pages:

1. Home Page: `http://localhost:4000/`
2. Test Subject Page: `http://localhost:4000/science/chemistry/chemical_reactions/reactions` (and any other url-like paths in the static/content directory)

## Future Plans

- Add more scientific topics
- Add more test subjects
- Add `Output()` function to every content file that demonstrates the concept
    = this is weird since we are not executing the code in the content files.....we may need to run and store the outputs in a static file on build? Maybe another sideloaded package to handle this?
- Feature to display nested content paths when navigating to a non-content directory eg. /science/ should render a list of accessible paths underneat, and a list of .go content files nested directly in this directory

