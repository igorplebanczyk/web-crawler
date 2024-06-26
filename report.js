export { printReport };

function printReport(pages) {
    console.log("\n--------------------------------\nReport:\n");

    const sortedPages = Object.entries(pages)
    .sort(([,a],[,b]) => b-a)
    .reduce((r, [k, v]) => ({ ...r, [k]: v }), {});

    for (const page in sortedPages) {
        console.log(`Found ${sortedPages[page]} internal links to ${page}`);
    }

    console.log("\n--------------------------------\n");
}