const fs = require('fs');
const path = require('path');

const DIRECTORY_PATH = './';

async function findGoFiles(directoryPath) {
    const files = await fs.promises.readdir(directoryPath);
    const goFiles = [];

    for (const file of files) {
        const filePath = path.join(directoryPath, file);
        const stat = await fs.promises.stat(filePath);
        if (stat.isDirectory()) {
            let subDirFiles = (await findGoFiles(filePath)).map(subDirFile => path.join(file, subDirFile));
            goFiles.push(...subDirFiles);
        } else if (path.extname(file) === '.go') {
            goFiles.push(file);
        }
    }

    return goFiles;
}

const isImportLine = (line) => {
    return line.startsWith('import')
        || line.startsWith('package')
        || line.startsWith('"')
        || line.startsWith('_');
}

const isCommentLine = (line) => {
    return line.startsWith('//') || line.startsWith('/*');
}

const shouldLineCount = (line) => {
    return line !== '' // Remove empty lines
        && !isCommentLine(line) // Remove comments
        && !isImportLine(line) // Remove import lines
        && line !== '}' && line !== ')'; // Remove closing brackets
}


const countLinesOfCode = (filePath, fileContent) => {
    const lines = fileContent.split('\n');
    const trimmedLines = lines.map(line => line.trim());

    const typeLines = trimmedLines.filter(line => line.startsWith('type'));
    if (typeLines.length > 0) {
        console.log(filePath, 0);
        return 0;
    }

    const countedLines = trimmedLines.filter(shouldLineCount);

    console.log(filePath, countedLines.length);

    return countedLines.length;
}

const countTotalLinesOfCode = (goFiles) => {
    let totalLines = 0;

    goFiles.forEach((file) => {
        const filePath = path.join(DIRECTORY_PATH, file);
        const fileContent = fs.readFileSync(filePath, 'utf-8');
        const linesOfCode = countLinesOfCode(filePath, fileContent);
        totalLines += linesOfCode;
    });

    console.log(`Total lines of code in ${goFiles.length} .go files: ${totalLines}`);
}

(async () => {
    const goFiles = await findGoFiles(DIRECTORY_PATH);

    countTotalLinesOfCode(goFiles);
})();
