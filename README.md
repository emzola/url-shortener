# url-shortener

POST a URL to <code>/v1/shorturl</code> endpoint and get a JSON response with short_url properties. 

Example JSON Reponse: <code>{short_url : 'a'}</code>

Visit <code>/v1/shorturl/<short_url></code> to be redirected to the original URL.

If you pass an invalid URL that doesn't follow the valid http://www.example.com format, the JSON response will contain <code>{error: 'invalid url'}</code>

Make a DELETE request to <code>/v1/shorturl/<short_url></code>, the short_url will be deleted from the Database and return a JSON response.

Example: <code>{message : 'short url successfully deleted'}</code>


