# Javascript

## Memory vs. DOM Storage

> I am more at risk to slow down my code execution by holding these things in memory as objects instead of on the DOM? Also, since these objects will need to be accessed by multiple parts of the code, these would have to be global objects. Right?

> There's one rule that is a constant in web development: the DOM is slow. Selecting, accessing, and iterating DOM elements and reading and modifying their properties is one of the single slowest things you can do in browser JS. On the other hand, working with native JS objects is very fast. As a rule, I avoid working with the DOM as much as possible.

## NPM Access Tokens

[NPM Access Token](https://docs.npmjs.com/about-access-tokens)

> An access token is an alternative to using your username and password for authenticating to npm when using the API or the npm command-line interface (CLI). An access token is a hexadecimal string that you can use to authenticate, and which gives you the right to install and/or publish your modules.

> The npm CLI automatically generates an access token for you when you run npm login. You can also create an access token to give other tools (such as continuous integration testing environments) access to your npm packages. For example, GitHub Actions provides the ability to store secrets, like access tokens, that you can then use to authenticate. When your workflow runs, it will be able to complete npm tasks as you, including installing private packages you can access.
# Resources

* [Stackoverflow Memory vs Dom](https://stackoverflow.com/questions/14993681/javascript-script-memory-vs-dom-storage)