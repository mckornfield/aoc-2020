import { readFileInputByLines } from '../../reader/src/main';
import { determineAnswerCount, determineAnswerCountForList, determineAnswerCountForListPt2 } from './day6';
test('said yes to 6 questions', () => {
    const answers = `abcx
    abcy
    abcz`

    expect(determineAnswerCount(answers)).toBe(6);
})

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

test('count unique answers from separate groups', () => {
    expect(determineAnswerCountForList(listsOfAnswers)).toBe(11);
})

test('part 1 let count all the answers', () => {
    const answerList = readFileInputByLines('day6');
    expect(determineAnswerCountForList(answerList)).toBe(6565);
})

test('pt 2 count same answers from separate groups', () => {
    expect(determineAnswerCountForListPt2(listsOfAnswers)).toBe(6);
})


test('part 2 let count all the common answers', () => {
    const answerList = readFileInputByLines('day6');
    expect(determineAnswerCountForListPt2(answerList)).toBe(3137);
})
