package handlers

// This is the file for constants used in the handlers

// Headers for conent type
const headerContentType = "Content-Type"
const contentTypeJSON = "application/json"
const contentTypeText = "text/plain"

// CORS
const accessControlAllowOrigin = "Access-Control-Allow-Origin"
const accessControlValue = "*"

// What methods we want to allow
const accessControlAllowMethods = "Access-Control-Allow-Methods"
const accessControlMethods = "GET,POST,DELETE,PATCH"

// Expose these headers
const exposeHeaders = "Access-Control-Expose-Headers"
const exposedHeaders = "Authorization"

// Allow these headers
const allowHeaders = "Access-Control-Allow-Headers"
const allowedHeaders = "Authorization"

// Allowed age for access control
const accessControlAllowAge = "Access-Control-Max-Age"
const accessControlAge = "600"
