/*
Provide a low level interface to the Discord API but also maintain clean readable code.
TODO think about the license
TODO need any nulls?

RestAPI does not use One Methods e.g. Channels().One(cID) because too much stutter.
RestAPI does not have all methods defined on single Client because too little structure and hard code completion/doc finding
RestAPI Client methods follow Discord API naming as close as possible ignoring the HTTP method names. HTTP method names do not always get the point across on what is happening. E.g. POST on Messages().BulkDelete or TriggerTypingIndicator
RestAPI Client must be discoverable
RestAPI Client must be easy to use and easy to read the resulting code (more important to have good readability and 0 magic
*/
package discgo
