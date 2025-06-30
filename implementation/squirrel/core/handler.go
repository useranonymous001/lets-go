package core

/*
This defines, the handler function that the user can create for each route
*/

type HandlerFunc func(*Request, *Response) // handler function
