import { readFileSync } from 'fs';
import path from 'path';

export function readFileInputByLines(fileName: string): string[] {
    return readFileSync(path.resolve(__dirname, fileName), 'utf-8')
        .split('\n').filter(Boolean)

}
