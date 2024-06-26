import { crawlPage } from './crawl.js';
import { printReport } from './report.js';

async function main() {
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
        const crawledPages = await crawlPage(baseURL);
        console.log('Crawl complete.');
        printReport(crawledPages);
    }
}

main();