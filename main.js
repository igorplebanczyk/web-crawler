import { crawlPage } from './crawl.js';

function main() {
    if (process.argv.length < 3) {
        console.log('Provide a URL to start crawling from.');
        return;
    } 
    else if (process.argv.length > 3) {
        console.log('Too many arguments');
        return;
    }
    else {
        const baseURL = process.argv[2];
        console.log(`Starting crawl at ${baseURL}...`);
        console.log(JSON.stringify(crawlPage(baseURL)));
    }
}

main();