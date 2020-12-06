import { readFileInputByLines } from '../../reader/src/main';
import { determineAnswerCount, determineAnswerCountForList } from './day6'
test('said yes to 6 questions', () => {
    const answers = `abcx
    abcy
    abcz`

    expect(determineAnswerCount(answers)).toBe(6);
})

test('count unique answers from separate groups', () => {
    const listsOfAnswers = ['abc',
        '',
        'a',
        'b',
        'c',
        '',
        'ab',
        'ac',
        '',
        'a',
        'a',
        'a',
        'a',
        '',
        'b']


    expect(determineAnswerCountForList(listsOfAnswers)).toBe(11);
})

test('part 1 let count all the answers',() => {
    const answerList = readFileInputByLines('day6');
    expect(determineAnswerCountForList(answerList)).toBe(6565);
})