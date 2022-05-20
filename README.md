# WC Builder

Is a simple tool written in Go for creating HTML custom elements.

It does only 2 things:
    - Generates a new customa element extenting HTMLElement
    - Bundle elements into a single js file

## Setting up


## Generating element

By using the command:
```
wc-builder new <element-name>
```

A folder named elements will be generating containing your element separated between `element.html` and `element.js`.


### Generated files
**elements/element/element.js**

This file contains a class with the default methods of a HTMLElement lifecycle.

**elements/element/element.html**

This file contains the template which will be attached to the shadow DOM.


## Bundling

By using the command:
```
wc-builder build
```

All elements will be bundled into a file in **dist/bundle.js**