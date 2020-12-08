import { readFileSync } from 'fs';
import path from 'path';

export function readFileInputByLines(fileName: string): string[] {
    return readFileInputToString(fileName)
        .split('\n');
}
export function readFileInputByLinesSkipLastLine(fileName: string): string[] {
    const lines = readFileInputByLines(fileName);
    lines.pop();
    return lines;
}

export function readFileInputToString(fileName: string): string {
    const puzzleSuffix = "_puzzle.txt"
    if (!fileName.endsWith(puzzleSuffix)) {
        fileName += puzzleSuffix;
    }
    return readFileSync(path.resolve(__dirname, fileName), 'utf-8')
}
