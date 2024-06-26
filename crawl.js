import { JSDOM } from 'jsdom';

export { normalizeURL };
export { getURLsFromHTML };
export { crawlPage };

function normalizeURL(url) {
    if (!new URL(url)) { throw new Error('Invalid URL') }; 
    return url.replace(/^https?:\/\//, '').replace(/\/$/, '');
}

function getURLsFromHTML(htmlBody, baseURL) {
    const dom = new JSDOM(htmlBody);
    const base = new URL(baseURL);
    return Array.from(dom.window.document.querySelectorAll('a')).map(a => new URL(a.href, base).href);
}

async function crawlPage(baseURL) {
    try {
        const response = await fetch(baseURL);
        const html = await response.text();
        
        if (response.status >= 400) { throw new Error(`Failed to fetch page: ${response.status}`); }
        if (response.headers.get('content-type').indexOf('text/html') === -1) { throw new Error('Page is not HTML'); }

        console.log(html);
    } catch (error) {
        console.error(error);
    }
}