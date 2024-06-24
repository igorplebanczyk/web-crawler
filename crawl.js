export { normalizeURL };

function normalizeURL(url) {
    if (!new URL(url)) { throw new Error('Invalid URL') }; 
    return url.replace(/^https?:\/\//, '').replace(/\/$/, '');
}