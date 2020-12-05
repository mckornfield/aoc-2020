import { readFileSync } from 'fs';
import path from 'path';

export function readFileInputByLines(fileName: string): string[] {
    return readFileInputToString(fileName)
        .split('\n').filter(Boolean)
}

export function readFileInputToString(fileName: string): string {
    const puzzleSuffix = "_puzzle.txt"
    if (!fileName.endsWith(puzzleSuffix)) {
        fileName += puzzleSuffix;
    }
    return readFileSync(path.resolve(__dirname, fileName), 'utf-8')
}
