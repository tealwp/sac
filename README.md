# Science As Code

This project aims to bridge the gap between science and code. It provides a platform for people with a scientific background to explain complex concepts using basic programming models, types, functions, and methods.

For the same reasons, this project is a great place for an engineer who understands code, and how it operates, to learn high-level scientific processes or topics.

Generally speaking the code in the content directory should build off of eachother.

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
    │   └── [subject4]/
    |       └── [topic].go
    └── chemistry/
        └── chemical_reactions/
            └── reactions.go
```

**Note:** these can be any depth of nested content.

This structure corresponds to URL paths like:
`http://localhost:4000/science/[category]/[subject]`

For example, a test page for chemical reactions is accessed at:
`http://localhost:4000/science/chemistry/chemical_reactions/reactions`

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
2. Topic Page: `http://localhost:4000/science/chemistry/chemical_bonding/bonding` (and any other url-like paths in the static/content directory)

## Contributing

Yay, please contribute with content or server code!

For **server code:**
- The usual expectations apply; Submit an MR and you'll get a review.
- If you're updating the UI, thank you. It is uuugggllllyyyyyy.
- Try to follow best practices with file structure. I know that isn't entirely the case right now with existing code.
- Looking for something to contribute here? Scroll down to the [Future Plans section](#future-plans) for some ideas, these are generally ordered by priority.


For **content code:**
- These files can import one another, and the standard library. Nothing else.
- The code is not meant to be executable -- or even useful. But it is meant to not break the linter.
- Comments are great, the file parser uses these.
- Try to keep it high-level.
- Try to define constants as...well....constants.
- Remember, you are teaching somebody a science topic, and they understand code. **Code is the medium, don't lose sight of that.**
- The content shouldn't use any *magical syntax* that takes away from the underlying logic. e.g., don't use a std lib function/method that does multiple logical steps, taking away from the learning at this level.

## Future Plans

- Navigation bar feature to display nested content paths when navigating to a non-content directory eg. `/science/` should render a list of accessible paths underneath, and a list of .go content files nested directly in this directory. (This is really poorly started in the content handler.)
    - Consider the route `.../science/chemistry/atoms/subatomic/electrons`, we should show the content that *would* be shown for 
    `.../science/chemistry/atoms/subatomic/electrons/electrons.go` (same name file, in this directory) but also show the directory contents in a sidebar for navigation.
- Add more content files (currently *mostly* AI generated for testing purposes, and haven't been proofed at all)
- Make the internal imports `github.com/tealwp/sac/...` clickable, in the rendered UI, to go to the related nested content path.
- Add `Output()` function to every content file that demonstrates the concept
    - this is weird since we are not executing the code in the content files.....we may need to run and store the outputs in a static file on build? Maybe another sideloaded package to handle this?
