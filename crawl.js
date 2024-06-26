export { normalizeURL };
export { getURLsFromHTML };
import { JSDOM } from 'jsdom';

function normalizeURL(url) {
    if (!new URL(url)) { throw new Error('Invalid URL') }; 
    return url.replace(/^https?:\/\//, '').replace(/\/$/, '');
}

function getURLsFromHTML(htmlBody, baseURL) {
    const dom = new JSDOM(htmlBody);
    const base = new URL(baseURL);
    return Array.from(dom.window.document.querySelectorAll('a')).map(a => new URL(a.href, base).href);
}