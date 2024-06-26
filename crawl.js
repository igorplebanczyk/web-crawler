import { JSDOM } from 'jsdom';
import fetch from 'node-fetch';

export { normalizeURL, getURLsFromHTML, crawlPage };

function normalizeURL(url) {
    url = new URL(url)

    let fullURL = `${url.host}${url.pathname}`
    if (fullURL.slice(-1) === '/') { fullURL = fullURL.slice(0, -1) }

    return fullURL
  }

function getURLsFromHTML(htmlBody, baseURL) {
    const dom = new JSDOM(htmlBody);
    const base = new URL(baseURL);
    return Array.from(dom.window.document.querySelectorAll('a')).map(a => new URL(a.href, base).href);
}

async function getHTMLFromURL(url) {
    try {
        const response = await fetch(url);
        if (response.status >= 400) { throw new Error(`Failed to fetch page: ${response.status}`); }
        if (response.headers.get('content-type').indexOf('text/html') === -1) { throw new Error('Page is not HTML'); }
        const html = await response.text();
        return html;
    } catch (error) {
        console.error(error.message);
    }
}

async function crawlPage(baseURL, currentURL = baseURL, pages = {}) {
    if (new URL(currentURL).hostname !== new URL(baseURL).hostname) { return pages; }
    
    const normalizedURL = normalizeURL(currentURL);

    if (pages[normalizedURL] > 0) { 
        pages[normalizedURL]++;
        return pages; 
    }

    pages[normalizedURL] = 1;

    console.log(`Crawling ${currentURL}...`);
    let html = '';
    try {
        html = await getHTMLFromURL(currentURL);
    } catch (error) {
        console.log(error.message);
        return pages;
    }

    const urls = getURLsFromHTML(html, baseURL);
    for (const url of urls) { await crawlPage(baseURL, url, pages); }

    return pages;
}